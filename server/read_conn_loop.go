package server

import (
	"errors"
	"github.com/reyoung/gt/proto"
	"io"
	"log"
	"net"
)

func readConnLoop(serial uint32, conn net.Conn, listenServer proto.GT_ListenServer,
	done chan<- struct{}) {
	defer func() {
		done <- struct{}{}
	}()

	var buf [2048]byte
	var err error
	for {
		var n int
		n, err = conn.Read(buf[:])
		if err != nil {
			break
		}

		err = listenServer.Send(&proto.ListenResponse{Rsp: &proto.ListenResponse_Payload{
			Payload: &proto.DataWithSerial{
				Data: &proto.Data{
					Data: buf[:n],
				},
				SerialId: serial,
			},
		}})
		if err != nil {
			log.Printf("send data failed: %v\n", err)
			return
		}
	}

	if errors.Is(err, io.EOF) {
		err = listenServer.Send(&proto.ListenResponse{Rsp: &proto.ListenResponse_Payload{Payload: &proto.DataWithSerial{
			Data:     &proto.Data{Close: true},
			SerialId: serial,
		}}})
	}

	if err != nil {
		log.Printf("send close failed: %v\n", err)
	}
	return
}
