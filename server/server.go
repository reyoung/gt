package server

import (
	"fmt"
	"github.com/reyoung/gt/proto"
	"log"
	"net"
	"strings"
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

	case *proto.Request_Head_Exec_:
		return execServer(v.Exec, gtServer)

	case *proto.Request_Head_Dial_:
		return dialServer(v.Dial, gtServer)
	}

	return fmt.Errorf("unknown head type: %T", head.Head)
}

func (s *Server) Listen(listenServer proto.GT_ListenServer) error {
	req, err := listenServer.Recv()
	if err != nil {
		return fmt.Errorf("recv head frame failed: %w", err)
	}
	head := req.GetHead()
	if head == nil {
		return fmt.Errorf("head should be not nil")
	}
	address := head.GetAddress()
	if address == "" {
		return fmt.Errorf("address should be not empty")
	}
	address = strings.Replace(address, "localhost", "0.0.0.0", 1)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		_ = listenServer.Send(&proto.ListenResponse{Rsp: &proto.ListenResponse_Head_{Head: &proto.ListenResponse_Head{
			Error: err.Error(),
		}}})
		return fmt.Errorf("listen failed: %w", err)
	}
	defer listener.Close()

	log.Printf("listen at %s", listener.Addr())

	err = listenServer.Send(&proto.ListenResponse{Rsp: &proto.ListenResponse_Head_{Head: &proto.ListenResponse_Head{
		Port: uint32(listener.Addr().(*net.TCPAddr).Port),
	}}})
	if err != nil {
		return fmt.Errorf("send listen response failed: %w", err)
	}

	cMap := newConnMap()

	var writeLoop writeConnLoop
	writeLoop.Start(cMap, listenServer)

	for {
		conn, err := listener.Accept()
		if err != nil {
			_ = listenServer.Send(&proto.ListenResponse{Rsp: &proto.ListenResponse_Head_{Head: &proto.ListenResponse_Head{
				Error: err.Error(),
			}}})
			return fmt.Errorf("accept failed: %w", err)
		}

		log.Printf("listener accepted %s", address)

		serialID := cMap.Put(conn)
		addr := conn.RemoteAddr().(*net.TCPAddr)

		doneChan := make(chan struct{}, 2) // 2 means read/write

		go func() {
			<-doneChan
			<-doneChan
			cMap.Delete(serialID)
		}()

		writeLoop.writeDone <- writeConnDoneChan{Serial: serialID, Done: doneChan}

		err = listenServer.Send(&proto.ListenResponse{Rsp: &proto.ListenResponse_Accept_{Accept: &proto.ListenResponse_Accept{
			Addr:     addr.IP.String(),
			Port:     uint32(addr.Port),
			SerialId: serialID,
		}}})
		if err != nil {
			return fmt.Errorf("send accept response failed: %w", err)
		}

		go readConnLoop(serialID, conn, listenServer, doneChan)

	}

	return nil
}
