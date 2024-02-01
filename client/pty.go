package client

import (
	"fmt"
	"github.com/gliderlabs/ssh"
	"github.com/reyoung/gt/common"
	"github.com/reyoung/gt/proto"
	"io"
	"log"
)

func Pty(client proto.GTClient, ptyReq ssh.Pty, winCh <-chan ssh.Window, s ssh.Session) error {
	log.Printf("pty allocated, %s", s.RawCommand())

	cli, err := client.GT(s.Context())
	if err != nil {
		return fmt.Errorf("call GT failed: %w", err)
	}

	err = cli.Send(&proto.Request{Req: &proto.Request_Head_{
		Head: &proto.Request_Head{Head: &proto.Request_Head_Pty_{Pty: &proto.Request_Head_Pty{
			Term: ptyReq.Term,
			ResizeWindow: &proto.Request_ResizeWindow{
				Rows: uint32(ptyReq.Window.Height),
				Cols: uint32(ptyReq.Window.Width),
			},
		}}},
	}})
	if err != nil {
		return fmt.Errorf("send head frame failed: %w", err)
	}

	rsp, err := cli.Recv()
	if err != nil {
		return fmt.Errorf("recv head frame failed: %w", err)
	}
	if head := rsp.GetHead(); head == nil {
		return fmt.Errorf("head is nil, first frame should be head")
	} else if err := head.GetError(); err != nil {
		return fmt.Errorf("head error: %s", err.Message)
	}

	go func() {
		for win := range winCh {
			err := cli.Send(&proto.Request{Req: &proto.Request_ResizeWindow_{
				ResizeWindow: &proto.Request_ResizeWindow{
					Rows: uint32(win.Height),
					Cols: uint32(win.Width),
				},
			}})
			if err != nil {
				log.Printf("send resize window frame failed: %s", err.Error())
				return
			}
		}
	}()

	dataStream := newDataStream(cli)
	rwc := common.DataStreamToReadWriteCloser(dataStream)

	go func() {
		_, _ = io.Copy(rwc, s)
	}()
	_, _ = io.Copy(s, rwc)
	return nil
}
