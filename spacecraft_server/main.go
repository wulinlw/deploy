package main

import (
	"log"
	"net"

	"../deploy/command"
	"../deploy/upload"
	sc "../spacecraft"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement spacecraft.BaseServer.
type server struct{}

func (s *server) SvnUp(context.Context, *sc.SvnUpParam) (*sc.VersionNum, error) {
	return &sc.VersionNum{Version: 123}, nil
}
func (s *server) SvnCheckout(ctx context.Context, in *sc.SvnCheckoutParams) (*sc.VersionNum, error) {
	version := command.SvnCheckout(in.SvnUrl, in.Dir)
	return &sc.VersionNum{Version: int32(version)}, nil
}
func (s *server) SvnUpToRevision(context.Context, *sc.SvnUpToRevisionParams) (*sc.VersionNum, error) {
	return &sc.VersionNum{Version: 123}, nil
}
func (s *server) SvnInfo(context.Context, *sc.SvnUpParam) (*sc.ResponseStr, error) {
	return &sc.ResponseStr{String_: "okkkkkkkkkkk"}, nil
}
func (s *server) SpecifiedCommand(context.Context, *sc.SpecifiedCommandParams) (*sc.ResponseStr, error) {
	return &sc.ResponseStr{String_: "okkkkkkkkkkk"}, nil
}
func (s *server) ComplexCommand(context.Context, *sc.SpecifiedCommandParams) (*sc.ResponseStr, error) {
	return &sc.ResponseStr{String_: "okkkkkkkkkkk"}, nil
}
func (s *server) SendFile(ctx context.Context, in *sc.SendFileParams) (*sc.ResponseStr, error) {
	upload.Upload(in.FileAbsolutePath, in.FileContent, in.StoragePath)
	return &sc.ResponseStr{String_: "okkkkkkkkkkk"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	sc.RegisterSpacecraftServer(s, &server{})
	s.Serve(lis)
}
