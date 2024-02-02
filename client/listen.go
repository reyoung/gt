package client

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gliderlabs/ssh"
	"github.com/reyoung/gt/proto"
	"io"
	"log"
	"sync"
)

type acceptResult struct {
	Accept   *proto.ListenResponse_Accept
	recvChan <-chan *proto.Data
}

type remoteListener struct {
	listenClient      proto.GT_ListenClient
	port              int
	acceptChan        chan *acceptResult
	recvChan          map[uint32]chan *proto.Data
	sendChan          chan *proto.DataWithSerial
	sendRecvChanMutex sync.RWMutex
}

func (r *remoteListener) start() {
	r.acceptChan = make(chan *acceptResult, 4096)
	r.recvChan = make(map[uint32]chan *proto.Data)
	r.sendChan = make(chan *proto.DataWithSerial, 4096)

	go func() {
		for {
			item := <-r.sendChan
			err := r.listenClient.Send(&proto.ListenRequest{Req: &proto.ListenRequest_Payload{Payload: item}})
			if err != nil {
				log.Printf("send payload failed: %v", err)
				return
			}
		}
	}()

	go func() {
		for {
			recv, err := r.listenClient.Recv()
			if err != nil {
				log.Printf("recv failed: %v", err)
				return
			}

			if accept := recv.GetAccept(); accept != nil {
				recvCh := make(chan *proto.Data, 4096)
				accessResult := &acceptResult{
					Accept:   accept,
					recvChan: recvCh,
				}
				r.recvChan[accept.SerialId] = recvCh
				r.acceptChan <- accessResult
				continue
			}

			if payload := recv.GetPayload(); payload != nil {
				ch, ok := r.recvChan[payload.SerialId]
				if !ok {
					log.Printf("recv chan not found: %d", payload.SerialId)
					continue
				}

				ch <- payload.Data
				if payload.Data.Close {
					delete(r.recvChan, payload.SerialId)
					close(ch)
				}
				continue
			}

			log.Printf("unknown frame: %v", recv)
		}
	}()
}

type remoteRWC struct {
	Accept   *proto.ListenResponse_Accept
	recvChan <-chan *proto.Data
	buffer   *bytes.Buffer
	sendChan chan<- *proto.DataWithSerial
}

func (r *remoteRWC) bufSize() int {
	if r.buffer == nil {
		return 0
	}
	return r.buffer.Len()
}

func (r *remoteRWC) Read(p []byte) (n int, err error) {
	if r.bufSize() == 0 {
		data, ok := <-r.recvChan
		if !ok {
			return 0, io.EOF
		}

		if data.Close {
			return 0, io.EOF
		}

		r.buffer = bytes.NewBuffer(data.Data)
		return r.Read(p)
	}
	return r.buffer.Read(p)
}

func (r *remoteRWC) Write(p []byte) (n int, err error) {
	n = len(p)
	for len(p) > 2048 {
		toSend := p[:2048]
		p = p[2048:]

		r.sendChan <- &proto.DataWithSerial{
			Data: &proto.Data{
				Data: toSend,
			},
			SerialId: r.Accept.SerialId,
		}
	}
	if len(p) > 0 {
		r.sendChan <- &proto.DataWithSerial{
			Data: &proto.Data{
				Data: p,
			},
		}
	}
	return 0, nil
}

func (r *remoteRWC) Close() error {
	r.sendChan <- &proto.DataWithSerial{
		Data: &proto.Data{
			Close: true,
		},
		SerialId: r.Accept.SerialId,
	}
	return nil
}

func (r *remoteListener) Accept() (*ssh.AcceptResult, error) {
	accept := <-r.acceptChan

	rwc := &remoteRWC{
		Accept:   accept.Accept,
		recvChan: accept.recvChan,
		sendChan: r.sendChan,
	}

	return &ssh.AcceptResult{
		RWC:           rwc,
		RemoteAddress: accept.Accept.Addr,
		RemotePort:    int(accept.Accept.Port),
	}, nil
}

func (r *remoteListener) Close() error {
	return r.listenClient.CloseSend()
}

func (r *remoteListener) Port() int {
	return r.port
}

func Listen(client proto.GTClient) func(address string) (ssh.RemoteListener, error) {
	return func(address string) (ssh.RemoteListener, error) {
		cli, err := client.Listen(context.Background())
		if err != nil {
			return nil, fmt.Errorf("listen failed: %w", err)
		}

		err = cli.Send(&proto.ListenRequest{Req: &proto.ListenRequest_Head_{Head: &proto.ListenRequest_Head{Address: address}}})
		if err != nil {
			return nil, fmt.Errorf("send head frame failed: %w", err)
		}

		rsp, err := cli.Recv()
		if err != nil {
			return nil, fmt.Errorf("recv head frame failed: %w", err)
		}

		head := rsp.GetHead()
		if head == nil {
			return nil, fmt.Errorf("head should be not nil")
		}
		if head.Error != "" {
			return nil, fmt.Errorf("listen failed: %s", head.Error)
		}

		listener := &remoteListener{listenClient: cli, port: int(head.Port)}
		listener.start()
		return listener, nil
	}
}
