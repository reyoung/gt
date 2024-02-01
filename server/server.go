package server

import (
	"fmt"
	"github.com/reyoung/gt/proto"
)

type Server struct {
	proto.UnimplementedGTServer
}

func (s *Server) GT(gtServer proto.GT_GTServer) error {
	req, err := gtServer.Recv()
	if err != nil {
		return fmt.Errorf("recv head frame failed: %w", err)
	}
	head := req.GetHead()
	if head == nil {
		return fmt.Errorf("head is nil, first frame should be head")
	}
	switch v := head.Head.(type) {
	case *proto.Request_Head_Stfp_:
		return sftpServer(gtServer)

	case *proto.Request_Head_Pty_:
		return ptyServer(v.Pty, gtServer)
	}

	return fmt.Errorf("unknown head type: %T", head.Head)
}
