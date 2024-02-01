package main

import (
	"flag"
	"github.com/reyoung/gt/proto"
	"github.com/reyoung/gt/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	addr := flag.String("address", ":8080", "grpc address")
	flag.Parse()
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterGTServer(s, &server.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
