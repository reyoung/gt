package server

import (
	"fmt"
	"github.com/reyoung/gt/proto"
)

type serverDataStream struct {
	server           proto.GT_GTServer
	resizeWindowChan chan *proto.Request_ResizeWindow
}

func newServerDataStream(server proto.GT_GTServer) *serverDataStream {
	return &serverDataStream{
		server:           server,
		resizeWindowChan: make(chan *proto.Request_ResizeWindow, 4096),
	}
}

func (s *serverDataStream) Send(data *proto.Data) error {
	return s.server.Send(&proto.Response{Rsp: &proto.Response_Data{Data: data}})
}

func (s *serverDataStream) Recv() (*proto.Data, error) {
	rsp, err := s.server.Recv()
	if err != nil {
		return nil, fmt.Errorf("recv data frame failed: %w", err)
	}
	if resizeWindow := rsp.GetResizeWindow(); resizeWindow != nil {
		ctx := s.server.Context()
		select {
		case s.resizeWindowChan <- resizeWindow:
		case <-ctx.Done():
			return nil, ctx.Err()
		}
		return s.Recv()
	}

	data := rsp.GetData()
	if data == nil {
		return nil, fmt.Errorf("data is nil, frame should be data")
	}
	return data, nil
}
