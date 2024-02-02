package server

import (
	"github.com/reyoung/gt/proto"
	"log"
)

type writeConnDoneChan struct {
	Serial uint32
	Done   chan<- struct{}
}

type writeConnLoop struct {
	doneChans map[uint32]chan<- struct{}
	writeDone chan writeConnDoneChan
}

func (w *writeConnLoop) Start(cMap *connMap, listenServer proto.GT_ListenServer) {
	w.doneChans = make(map[uint32]chan<- struct{})
	w.writeDone = make(chan writeConnDoneChan, 4096)

	go func() {
		defer func() {
			for _, done := range w.doneChans {
				done <- struct{}{}
			}
		}()
		for {
			req, err := listenServer.Recv()
			if err != nil {
				log.Printf("recv head frame failed: %v", err)
				return
			}

			payload := req.GetPayload()
			if payload == nil {
				log.Printf("payload should be not nil")
				continue
			}

			for {
				select {
				case done := <-w.writeDone:
					w.doneChans[done.Serial] = done.Done
					continue
				default:
				}
				break
			}

			serial := payload.GetSerialId()
			conn := cMap.Get(serial)
			if conn == nil {
				log.Printf("conn not found: %d", serial)
				continue
			}
			doneCh := w.doneChans[serial]
			if doneCh == nil {
				log.Printf("done chan not found: %d", serial)
				continue
			}

			data := payload.GetData()
			if data == nil {
				log.Printf("data should be not nil")
				continue
			}

			if len(data.GetData()) != 0 {
				_, err := conn.Write(payload.GetData().GetData())
				if err != nil {
					log.Printf("write data failed: %v", err)
					doneCh <- struct{}{}
					delete(w.doneChans, serial)
					_ = conn.Close()
					continue
				}
			}

			if data.GetClose() {
				doneCh <- struct{}{}
				delete(w.doneChans, serial)
				_ = conn.Close()
				continue
			}
		}
	}()
}
