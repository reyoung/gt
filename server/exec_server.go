package server

import (
	"errors"
	"fmt"
	"github.com/reyoung/gt/common"
	"github.com/reyoung/gt/proto"
	"io"
	"log"
	"os/exec"
	"strings"
)

func postHeaderError(server proto.GT_GTServer, err error) error {
	if err != nil {
		_ = server.Send(&proto.Response{
			Rsp: &proto.Response_Head_{Head: &proto.Response_Head{Head: &proto.Response_Head_Error_{Error: &proto.Response_Head_Error{
				Message: err.Error(),
			}}}},
		})
		return err
	} else {
		return server.Send(&proto.Response{Rsp: &proto.Response_Head_{Head: &proto.Response_Head{}}})
	}
}

func postExecResponse(server proto.GT_GTServer, err error) error {
	if err != nil {
		var execRsp proto.Response_ExecDone
		execRsp.Error = err.Error()
		var execErr *exec.ExitError

		if errors.As(err, &execErr) {
			execRsp.ExitCode = uint32(execErr.ExitCode())
		}

		return server.Send(&proto.Response{Rsp: &proto.Response_ExecDone_{ExecDone: &execRsp}})
	}

	return server.Send(&proto.Response{Rsp: &proto.Response_ExecDone_{ExecDone: &proto.Response_ExecDone{}}})
}

func execServer(execReq *proto.Request_Head_Exec, server proto.GT_GTServer) error {
	if len(execReq.Commands) == 0 {
		return postHeaderError(server, fmt.Errorf("no command to execute"))
	}
	if err := postHeaderError(server, nil); err != nil {
		return fmt.Errorf("post header error: %w", err)
	}

	cmds := append([]string{"bash", "-c"}, execReq.Commands...)

	log.Printf("exec command: %s\n", strings.Join(cmds, " "))

	cmd := exec.Command(cmds[0], cmds[1:]...)
	ds := newServerDataStream(server)
	rwc := common.DataStreamToReadWriteCloser(ds)
	defer func() {
		rwc.Close()
	}()
	cmd.Stdout = rwc
	cmd.Stderr = rwc
	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		return postExecResponse(server, fmt.Errorf("get stdin pipe failed: %w", err))
	}
	defer func() {
		stdinPipe.Close()
	}()
	go func() {
		_, _ = io.Copy(stdinPipe, rwc)
	}()
	//cmd.Stdin = rwc

	if err := cmd.Start(); err != nil {
		return postExecResponse(server, fmt.Errorf("start command error: %w", err))
	}
	if err := cmd.Wait(); err != nil {
		return postExecResponse(server, fmt.Errorf("wait command error: %w", err))
	}
	return postExecResponse(server, nil)
}
