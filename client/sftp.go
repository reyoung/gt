package client

import (
	"fmt"
	"github.com/gliderlabs/ssh"
	"github.com/reyoung/gt/common"
	"github.com/reyoung/gt/proto"
	"io"
)

func SFtp(client proto.GTClient, s ssh.Session) error {
	ctx := s.Context()
	stream, err := client.GT(ctx)
	if err != nil {
		return fmt.Errorf("call GT failed: %w", err)
	}
	err = stream.Send(&proto.Request{Req: &proto.Request_Head_{
		Head: &proto.Request_Head{Head: &proto.Request_Head_Stfp_{Stfp: &proto.Request_Head_Stfp{}}}}})
	if err != nil {
		return fmt.Errorf("send head frame failed: %w", err)
	}
	rsp, err := stream.Recv()
	if err != nil {
		return fmt.Errorf("recv head frame failed: %w", err)
	}
	head := rsp.GetHead()
	if head == nil {
		return fmt.Errorf("head is nil, first frame should be head")
	}
	if err := head.GetError(); err != nil {
		return fmt.Errorf("head error: %s", err.Message)
	}

	rwc := common.DataStreamToReadWriteCloser(newDataStream(stream))

	go func() {
		_, _ = io.Copy(rwc, s)
	}()

	_, _ = io.Copy(s, rwc)
	return nil
}
