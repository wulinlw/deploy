package main

import (
	"fmt"
	"net/http"

	//"strings"
	sc "../spacecraft"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	//"io/ioutil"
	"log"
	"time"
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
		} else if funcName == "Live" {
			param := &sc.Empty{}
			result = live(param, req.FormValue("ip"))
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
	conn, c, err := createConnect(ip)
	if conn != nil { //连接失败时为nil，不能close
		defer conn.Close()
	}
	if err != nil {
		log.Printf("connect failed: %v", err)
		return "connect failed:" + ip
	}
	r, err := c.ComplexCommand(context.Background(), param.(*sc.SpecifiedCommandParams))
	if err != nil {
		log.Printf("call rpc failed: %v", err)
	}
	//	log.Printf("%#v", r)
	return string(r.String_)
}

func createConnect(ip string) (*grpc.ClientConn, sc.SpacecraftClient, error) {
	conn, err := grpc.Dial(ip, grpc.WithInsecure(), grpc.WithTimeout(time.Second*5), grpc.WithBlock())
	//defer conn.Close()
	c := sc.NewSpacecraftClient(conn)
	return conn, c, err
}

func sendfile(param interface{}, ip string) string {
	conn, c, err := createConnect(ip)
	if conn != nil { //连接失败时为nil，不能close
		defer conn.Close()
	}
	if err != nil {
		log.Printf("connect failed: %v", err)
		return "connect failed:" + ip
	}
	r, err := c.SendFile(context.Background(), param.(*sc.SendFileParams))
	if err != nil {
		log.Printf("could not greet: %v", err)
	}
	//	log.Printf("%#v", r)
	return string(r.String_)
}

func getFileList(param interface{}, ip string) string {
	conn, c, err := createConnect(ip)
	if conn != nil { //连接失败时为nil，不能close
		defer conn.Close()
	}
	if err != nil {
		log.Printf("connect failed: %v", err)
		return "connect failed:" + ip
	}
	r, err := c.GetFileList(context.Background(), param.(*sc.SvnUpParam))
	if err != nil {
		log.Printf("could not greet: %v", err)
	}
	//	log.Printf("%#v", r)
	return string(r.String_)
}

func live(param interface{}, ip string) string {
	conn, c, err := createConnect(ip)
	if conn != nil { //连接失败时为nil，不能close
		defer conn.Close()
	}
	if err != nil {
		log.Printf("connect failed: %v", err)
		return "connect failed:" + ip
	}
	r, err := c.Live(context.Background(), param.(*sc.Empty))
	if err != nil {
		log.Printf("could not greet: %v", err)
	}
	//	log.Printf("%#v", r)
	return string(r.String_)
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
