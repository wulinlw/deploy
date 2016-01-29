package main

import (
	"log"

	sc "../spacecraft"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
)

const (
	address = "192.168.9.97:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := sc.NewSpacecraftClient(conn)

	//r, err := c.SvnCheckout(context.Background(), &sc.SvnCheckoutParams{SvnUrl: "https://svn.td.gamebar.com/svn/private/liwai/test", Dir: "/root/goproject/svntest"})
	file := "E:/soft_package/17monipdb.exe"
	fileContent, err := ioutil.ReadFile(file)
	checkErr(err)
	r, err := c.SendFile(context.Background(), &sc.SendFileParams{
		FileAbsolutePath: file,
		FileContent:      fileContent,
		StoragePath:      "/root/goproject/svntest/xxx"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%#v", r)
	log.Println(r)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
