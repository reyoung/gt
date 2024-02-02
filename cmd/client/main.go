package main

import (
	context "context"
	"errors"
	"flag"
	"github.com/gliderlabs/ssh"
	"github.com/reyoung/gt/client"
	"github.com/reyoung/gt/common"
	"github.com/reyoung/gt/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
)

func main() {
	grpcAddr := flag.String("grpc_address", ":8080", "grpc address")
	svrAddr := flag.String("svr_address", ":2222", "ssh server address")
	flag.Parse()

	conn := common.Panic2(grpc.Dial(*grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials())))
	defer conn.Close()

	cli := proto.NewGTClient(conn)

	handler := &ssh.ForwardedTCPHandler{}

	svr := &ssh.Server{
		Addr: *svrAddr,
		Handler: func(session ssh.Session) {
			pty, winCh, isPty := session.Pty()

			if !isPty {
				log.Printf("no pty allocated, %s", session.RawCommand())
				err := client.Exec(cli, session)
				log.Printf("no pty allocated, exec done, %s %v", session.RawCommand(), err)

				if err != nil {
					session.Write([]byte(err.Error()))
					var errExecExit *client.ErrExecExit
					if ok := errors.As(err, &errExecExit); ok {
						session.Exit(errExecExit.Code())
					} else {
						session.Exit(-1)
					}
				} else {
					session.Exit(0)
				}

				return
			}

			err := client.Pty(cli, pty, winCh, session)
			if err != nil {
				session.Write([]byte(err.Error()))
				session.Exit(1)
				return
			}
			session.Exit(0)
		},
		LocalPortForwardingCallback: func(ctx ssh.Context, destinationHost string, destinationPort uint32) bool {
			log.Printf("local port forwarding: %s:%d", destinationHost, destinationPort)
			return true
		},
		ReversePortForwardingCallback: func(ctx ssh.Context, bindHost string, bindPort uint32) bool {
			log.Printf("reverse port forwarding: %s:%d", bindHost, bindPort)
			return true
		},
		ChannelHandlers: map[string]ssh.ChannelHandler{
			"direct-tcpip": ssh.DirectTCPIPHandler,
			"session":      ssh.DefaultSessionHandler,
		},
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
		RequestHandlers: map[string]ssh.RequestHandler{
			"tcpip-forward":        handler.HandleSSHRequest,
			"cancel-tcpip-forward": handler.HandleSSHRequest,
		},
		LocalPortForwardingDialer: func(ctx context.Context, network, address string) (io.ReadWriteCloser, error) {
			if network != "tcp" {
				return nil, errors.New("only tcp network supported")
			}
			return client.DialTCP(cli)(ctx, address)
		},
		ReversePortListen: client.Listen(cli),
	}
	svr.ListenAndServe()
}
