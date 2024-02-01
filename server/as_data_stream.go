package server

import (
	"fmt"
	"github.com/reyoung/gt/proto"
)

type serverDataStream struct {
	server proto.GT_GTServer
}

func (s *serverDataStream) Send(data *proto.Data) error {
	return s.server.Send(&proto.Response{Rsp: &proto.Response_Data{Data: data}})
}

func (s *serverDataStream) Recv() (*proto.Data, error) {
	rsp, err := s.server.Recv()
	if err != nil {
		return nil, fmt.Errorf("recv data frame failed: %w", err)
	}
	data := rsp.GetData()
	if data == nil {
		return nil, fmt.Errorf("data is nil, frame should be data")
	}
	return data, nil
}
