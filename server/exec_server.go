package server

import (
	"errors"
	"fmt"
	"github.com/reyoung/gt/common"
	"github.com/reyoung/gt/proto"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
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

func execServer(execReq *proto.Request_Head_Exec, server proto.GT_GTServer) (err error) {

	if err := postHeaderError(server, nil); err != nil {
		return fmt.Errorf("post header error: %w", err)
	}

	// see https://github.com/openssh/openssh-portable/blob/master/session.c#L1709
	cmds := []string{"bash", "-c", execReq.Command}
	if execReq.Command == "" {
		cmds = []string{"bash", "-"}
	}

	log.Printf("exec command: %s\n", strings.Join(cmds, " "))
	defer func() {
		log.Printf("exec command: %s done %v\n", strings.Join(cmds, " "), err)
	}()

	cmd := exec.Command(cmds[0], cmds[1:]...)
	cmd.Env = execReq.Envs
	for _, envKey := range []string{"PATH", "HOME", "SHELL", "USER", "TZ", "DISPLAY", "TMPDIR"} {
		envVal, ok := os.LookupEnv(envKey)
		if ok {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", envKey, envVal))
		}
	}
	cmd.Env = append(cmd.Env, "SSH_ORIGINAL_COMMAND="+execReq.Command)

	ds := newServerDataStream(server)
	rwc := common.DataStreamToReadWriteCloser(ds)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return postExecResponse(server, fmt.Errorf("get stdout pipe failed: %w", err))
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return postExecResponse(server, fmt.Errorf("get stderr pipe failed: %w", err))
	}

	var outClosed sync.WaitGroup
	outClosed.Add(2)

	go func() {
		defer outClosed.Done()
		defer stdoutPipe.Close()
		_, _ = io.Copy(rwc, stdoutPipe)
	}()

	go func() {
		defer outClosed.Done()
		defer stderrPipe.Close()
		_, _ = io.Copy(rwc, stderrPipe)
	}()

	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		return postExecResponse(server, fmt.Errorf("get stdin pipe failed: %w", err))
	}
	go func() {
		_, _ = io.Copy(stdinPipe, rwc)
		stdinPipe.Close()
	}()

	if err := cmd.Start(); err != nil {
		return postExecResponse(server, fmt.Errorf("start command error: %w", err))
	}
	if err := cmd.Wait(); err != nil {
		return postExecResponse(server, fmt.Errorf("wait command error: %w", err))
	}
	outClosed.Wait()
	rwc.Close()
	return postExecResponse(server, nil)
}
