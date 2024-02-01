package server

import (
	"fmt"
	"github.com/creack/pty"
	"github.com/reyoung/gt/common"
	"github.com/reyoung/gt/proto"
	"io"
	"log"
	"os/exec"
)

func ptyServer(ptyReq *proto.Request_Head_Pty, server proto.GT_GTServer) error {
	c := exec.Command("/bin/bash")
	f, err := pty.Start(c)
	if err != nil {
		server.Send(&proto.Response{Rsp: &proto.Response_Head_{Head: &proto.Response_Head{
			Head: &proto.Response_Head_Error_{Error: &proto.Response_Head_Error{Message: err.Error()}},
		}}})
		return fmt.Errorf("start pty failed: %w", err)
	}
	err = pty.Setsize(f, &pty.Winsize{
		Rows: uint16(ptyReq.GetResizeWindow().Rows),
		Cols: uint16(ptyReq.GetResizeWindow().Cols),
	})
	if err != nil {
		server.Send(&proto.Response{Rsp: &proto.Response_Head_{Head: &proto.Response_Head{
			Head: &proto.Response_Head_Error_{Error: &proto.Response_Head_Error{Message: err.Error()}},
		}}})
		return fmt.Errorf("set pty size failed: %w", err)
	}

	err = server.Send(&proto.Response{Rsp: &proto.Response_Head_{Head: &proto.Response_Head{}}})
	if err != nil {
		return fmt.Errorf("send head frame failed: %w", err)
	}
	log.Printf("start pty success\n")

	done := make(chan struct{})
	defer func() {
		close(done)
	}()

	dataStream := newServerDataStream(server)
	go func() {
		for {
			select {
			case size := <-dataStream.resizeWindowChan:
				err := pty.Setsize(f, &pty.Winsize{
					Rows: uint16(size.Rows),
					Cols: uint16(size.Cols),
				})
				if err != nil {
					log.Printf("set pty size failed: %s", err.Error())
				}
			case <-done:
				return
			}
		}
	}()

	rwc := common.DataStreamToReadWriteCloser(dataStream)

	go func() {
		io.Copy(f, rwc)
	}()
	io.Copy(rwc, f)
	return nil
}
