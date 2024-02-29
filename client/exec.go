package client

import (
	"fmt"
	"github.com/gliderlabs/ssh"
	"github.com/reyoung/gt/common"
	"github.com/reyoung/gt/proto"
	"io"
	"log"
)

type ErrExecExit struct {
	execDone *proto.Response_ExecDone
}

func (e ErrExecExit) Error() string {
	return e.execDone.String()
}

func (e ErrExecExit) Code() int {
	if e.execDone.Error != "" {
		if e.execDone.ExitCode == 0 {
			return -1
		} else {
			return int(e.execDone.ExitCode)
		}
	}
	return int(e.execDone.ExitCode)
}

func Exec(cli proto.GTClient, session ssh.Session) error {
	gtCli, err := cli.GT(session.Context())
	if err != nil {
		return fmt.Errorf("call GT failed: %w", err)
	}
	err = gtCli.Send(&proto.Request{Req: &proto.Request_Head_{Head: &proto.Request_Head{Head: &proto.Request_Head_Exec_{
		Exec: &proto.Request_Head_Exec{
			Command: session.RawCommand(),
			Envs:    session.Environ(),
		}}}}})

	if err != nil {
		return fmt.Errorf("send head frame failed: %w", err)
	}
	//session

	rsp, err := gtCli.Recv()
	if err != nil {
		return fmt.Errorf("recv head frame failed: %w", err)
	}
	if head := rsp.GetHead(); head == nil {
		return fmt.Errorf("head is nil, first frame should be head")
	} else if err := head.GetError(); err != nil {
		return fmt.Errorf("head error: %s", err.Message)
	}

	ds := newDataStream(gtCli)
	rwc := common.DataStreamToReadWriteCloser(ds)

	go func() {
		_, _ = io.Copy(rwc, session)
		_ = rwc.Close()
	}()

	written, _ := io.Copy(session, rwc)
	log.Printf("exec done, written %d", written)

	execDone, ok := <-ds.execDoneChan
	if !ok {
		return fmt.Errorf("exec done channel closed")
	}

	if execDone.Error == "" && execDone.ExitCode == 0 {
		return nil
	} else {
		return ErrExecExit{execDone: execDone}
	}

}
