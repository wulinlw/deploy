package main

import (
	"log"
	"net"

	sc "../spacecraft"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement spacecraft.BaseServer.
type server struct{}

func (s *server) Testfunc1(ctx context.Context, in *sc.Testfunc1Params) (*sc.Testfunc1Result, error) {
	return &sc.Testfunc1Result{RId: "2", RName: "wulinlw", RPass: "123456"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	sc.RegisterBaseServer(s, &server{})
	s.Serve(lis)
}
