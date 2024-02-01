package main

import (
	"flag"
	"github.com/gliderlabs/ssh"
	"github.com/reyoung/gt/client"
	"github.com/reyoung/gt/common"
	"github.com/reyoung/gt/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	grpcAddr := flag.String("grpc_address", ":8080", "grpc address")
	svrAddr := flag.String("svr_address", ":2222", "ssh server address")
	flag.Parse()

	conn := common.Panic2(grpc.Dial(*grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials())))
	defer conn.Close()

	cli := proto.NewGTClient(conn)

	svr := &ssh.Server{
		Addr: *svrAddr,
		SubsystemHandlers: map[string]ssh.SubsystemHandler{
			"sftp": func(s ssh.Session) {
				err := client.SFtp(cli, s)
				if err != nil {
					s.Write([]byte(err.Error()))
					s.Exit(1)
					return
				}
				s.Exit(0)
			},
		},
	}
	svr.ListenAndServe()
}
