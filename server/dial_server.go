package server

import (
	"fmt"
	"github.com/reyoung/gt/common"
	"github.com/reyoung/gt/proto"
	"io"
	"net"
)

func dialServer(dial *proto.Request_Head_Dial, server proto.GT_GTServer) error {
	var dialer net.Dialer
	conn, err := dialer.DialContext(server.Context(), "tcp", dial.Addr)
	if err != nil {
		return postHeaderError(server, err)
	}
	if err := postHeaderError(server, nil); err != nil {
		return fmt.Errorf("post header error failed: %w", err)
	}
	dataStream := newServerDataStream(server)
	rwc := common.DataStreamToReadWriteCloser(dataStream)

	done := make(chan struct{})

	defer conn.Close()
	go func() {
		defer close(done)
		_, _ = io.Copy(conn, rwc)
	}()
	_, _ = io.Copy(rwc, conn)
	<-done
	return nil
}
