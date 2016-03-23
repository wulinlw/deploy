package main

import (
	"fmt"

	sc "../spacecraft"
	//"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"reflect"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure(), grpc.WithTimeout(time.Second*5), grpc.WithBlock())
	if err != nil {
		log.Println(err)
	}
	c := sc.NewSpacecraftClient(conn)

	ct := reflect.TypeOf(c)
	cc := reflect.ValueOf(c)
	fmt.Println(cc, cc.NumMethod(), cc.Method(0))
	fmt.Println(ct, ct.Method(0)) //method
	fmt.Println(ct.Method(0).Name)
	fmt.Println(ct.Method(0).PkgPath)
	fmt.Println(ct.Method(0).Type) //type
	fmt.Println(ct.Method(0).Func)
	fmt.Println(ct.Method(0).Index)
	fmt.Println(ct.Method(0).Type.In(2)) //参数的类型
}
