package client

import (
	"fmt"
	"github.com/reyoung/gt/proto"
)

type clientDataStream struct {
	client proto.GT_GTClient
}

func (c *clientDataStream) Send(data *proto.Data) error {
	return c.client.Send(&proto.Request{Req: &proto.Request_Data{Data: data}})
}

func (c *clientDataStream) Recv() (*proto.Data, error) {
	rsp, err := c.client.Recv()
	if err != nil {
		return nil, fmt.Errorf("recv data frame failed: %w", err)
	}
	data := rsp.GetData()
	if data == nil {
		return nil, fmt.Errorf("data is nil, frame should be data")
	}
	return data, nil
}
