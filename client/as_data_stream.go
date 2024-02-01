package client

import (
	"fmt"
	"github.com/reyoung/gt/proto"
)

type clientDataStream struct {
	client       proto.GT_GTClient
	execDoneChan chan *proto.Response_ExecDone
	dataChan     chan *proto.Data
	errChan      chan error
}

func (c *clientDataStream) Send(data *proto.Data) error {
	//log.Printf("send data frame: %s", data.String())
	return c.client.Send(&proto.Request{Req: &proto.Request_Data{Data: data}})
}

func (c *clientDataStream) Recv() (*proto.Data, error) {
	select {
	case data, ok := <-c.dataChan:
		if !ok {
			return nil, fmt.Errorf("data channel closed")
		}
		return data, nil
	case err := <-c.errChan:
		return nil, err
	}
}

func (c *clientDataStream) start() {
	go func() {
		defer func() {
			close(c.dataChan)
			close(c.errChan)
			close(c.execDoneChan)
		}()

		for {
			rsp, err := c.client.Recv()
			if err != nil {
				c.errChan <- fmt.Errorf("recv data frame failed: %w", err)
				return
			}
			//log.Printf("recv frame: %s", rsp.String())
			if execDone := rsp.GetExecDone(); execDone != nil {
				c.execDoneChan <- execDone
				continue
			}

			data := rsp.GetData()
			if data == nil {
				c.errChan <- fmt.Errorf("data is nil, frame should be data")
				return
			}
			c.dataChan <- data
		}
	}()
}

func newDataStream(client proto.GT_GTClient) *clientDataStream {
	ds := &clientDataStream{
		client:       client,
		execDoneChan: make(chan *proto.Response_ExecDone, 1),
		dataChan:     make(chan *proto.Data, 4096),
		errChan:      make(chan error, 1),
	}
	ds.start()
	return ds
}
