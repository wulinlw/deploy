package main

import (
	"log"

	sc "../spacecraft"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := sc.NewBaseClient(conn)

	r, err := c.Testfunc1(context.Background(), &sc.Testfunc1Params{Id: "1", Name: "my name?", Pass: "my pass!"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%#v", r)
	log.Println(r)
}
