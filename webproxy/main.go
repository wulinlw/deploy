package main

import (
	//"encoding/json"
	"fmt"
	"net/http"

	//"strings"
	sc "../spacecraft"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	//"io/ioutil"
	"log"
)

const (
	//address = "localhost:50051"
	address = "192.168.9.97:50051"
)

//form:gsid
func index(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if req.Form == nil {
		w.Write([]byte("{\"error\":\"没有参数\"}"))
		return
	} else {
		funcName := req.FormValue("funcName")
		for i, n := range req.Form {
			fmt.Println(i, n)
		}
		var result string
		if funcName == "ComplexCommand" {
			param := &sc.SpecifiedCommandParams{
				Command: req.FormValue("command"),
				Dir:     req.FormValue("dir")}
			//fmt.Println(param)
			//			run(param)
			result = run(param, req.FormValue("ip"))
			fmt.Println(result)
		} else if funcName == "SendFile" {
			param := &sc.SendFileParams{
				FileAbsolutePath: req.FormValue("FileAbsolutePath"),
				FileContent:      []byte(req.FormValue("FileContent")),
				StoragePath:      req.FormValue("StoragePath")}
			result = sendfile(param, req.FormValue("ip"))
			fmt.Println(result)
		} else if funcName == "GetFileList" {
			param := &sc.SvnUpParam{
				Dir: req.FormValue("dir")}
			result = getFileList(param, req.FormValue("ip"))
			fmt.Println(result)
		}

		w.Write([]byte(result))
	}
}

func main() {
	http.HandleFunc("/index", index)
	http.ListenAndServe(":8000", nil)
}

func run(param interface{}, ip string) string {
	// Set up a connection to the server.
	conn, err := grpc.Dial(ip, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := sc.NewSpacecraftClient(conn)
	r, err := c.ComplexCommand(context.Background(), param.(*sc.SpecifiedCommandParams))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	//	log.Printf("%#v", r)
	//	log.Println(r)
	return string(r.String_)
}

func sendfile(param interface{}, ip string) string {
	// Set up a connection to the server.
	conn, err := grpc.Dial(ip, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := sc.NewSpacecraftClient(conn)
	r, err := c.SendFile(context.Background(), param.(*sc.SendFileParams))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	//	log.Printf("%#v", r)
	//	log.Println(r)
	return string(r.String_)
}

func getFileList(param interface{}, ip string) string {
	// Set up a connection to the server.
	conn, err := grpc.Dial(ip, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := sc.NewSpacecraftClient(conn)
	r, err := c.GetFileList(context.Background(), param.(*sc.SvnUpParam))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	//	log.Printf("%#v", r)
	//	log.Println(r)
	return string(r.String_)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
