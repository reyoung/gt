package common

import (
	"bytes"
	"github.com/reyoung/gt/proto"
	"io"
)

type dataStreamRWC struct {
	stream     DataStream
	readBuffer *bytes.Buffer
	count      int
}

func (d *dataStreamRWC) bufSize() int {
	if d.readBuffer == nil {
		return 0
	}
	return d.readBuffer.Len()
}

func (d *dataStreamRWC) Read(p []byte) (n int, err error) {
	if d.bufSize() == 0 {
		data, err := d.stream.Recv()
		if err != nil {
			return 0, err
		}

		if data.Close {
			return 0, io.EOF
		}

		d.count += len(data.Data)

		d.readBuffer = bytes.NewBuffer(data.Data)
	}

	return d.readBuffer.Read(p)
}

func (d *dataStreamRWC) Write(p []byte) (n int, err error) {
	const maxChunkSize = 2 * 1024 * 1024
	nSend := 0
	for len(p) > maxChunkSize {
		err = d.stream.Send(&proto.Data{
			Data: p[:maxChunkSize],
		})
		if err != nil {
			return
		}
		nSend += maxChunkSize
		p = p[maxChunkSize:]
	}
	if len(p) != 0 {
		err = d.stream.Send(&proto.Data{
			Data: p,
		})
		if err != nil {
			return
		}
		nSend += len(p)
	}
	return nSend, nil
}

func (d *dataStreamRWC) Close() error {
	return d.stream.Send(&proto.Data{Close: true})
}

func DataStreamToReadWriteCloser(stream DataStream) io.ReadWriteCloser {
	return &dataStreamRWC{
		stream: stream,
	}
}
