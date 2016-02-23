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
			result = run(param)
			fmt.Println(result)
		}

		w.Write([]byte(result))
	}
}

func main() {
	http.HandleFunc("/index", index)
	http.ListenAndServe(":8000", nil)
}

func run(param interface{}) string {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
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

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
