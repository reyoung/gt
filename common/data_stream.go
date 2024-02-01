package common

import "github.com/reyoung/gt/proto"

type DataStream interface {
	Send(data *proto.Data) error
	Recv() (*proto.Data, error)
}
