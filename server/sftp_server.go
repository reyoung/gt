package server

import (
	"fmt"
	"github.com/pkg/sftp"
	"github.com/reyoung/gt/common"
	"github.com/reyoung/gt/proto"
)

func sftpServer(server proto.GT_GTServer) error {
	err := server.Send(&proto.Response{Rsp: &proto.Response_Head_{Head: &proto.Response_Head{}}})
	if err != nil {
		return fmt.Errorf("send head frame failed: %w", err)
	}
	rwc := common.DataStreamToReadWriteCloser(newServerDataStream(server))

	svr, err := sftp.NewServer(rwc)
	if err != nil {
		return fmt.Errorf("new sftp Server failed: %w", err)
	}
	defer func() {
		svr.Close()
	}()

	_ = svr.Serve()
	return nil
}
