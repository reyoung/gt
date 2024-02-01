package client

import (
	"context"
	"fmt"
	"github.com/reyoung/gt/common"
	"github.com/reyoung/gt/proto"
	"io"
)

func DialTCP(cli proto.GTClient) func(ctx context.Context, address string) (io.ReadWriteCloser, error) {
	return func(ctx context.Context, address string) (io.ReadWriteCloser, error) {
		client, err := cli.GT(ctx)
		if err != nil {
			return nil, fmt.Errorf("gt failed: %w", err)
		}
		err = client.Send(&proto.Request{Req: &proto.Request_Head_{Head: &proto.Request_Head{Head: &proto.Request_Head_Dial_{
			Dial: &proto.Request_Head_Dial{Addr: address},
		}}}})
		if err != nil {
			return nil, fmt.Errorf("send dial head frame failed: %w", err)
		}
		rsp, err := client.Recv()
		if err != nil {
			return nil, fmt.Errorf("recv dial head frame failed: %w", err)
		}
		if head := rsp.GetHead(); head == nil {
			return nil, fmt.Errorf("head is nil, first frame should be head")
		} else if err := head.GetError(); err != nil {
			return nil, fmt.Errorf("head error: %s", err.Message)
		}

		ds := newDataStream(client)
		return common.DataStreamToReadWriteCloser(ds), nil
	}
}
