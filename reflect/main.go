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
	pt := ct.Method(0).Type.In(2)
	fmt.Println(ct.Method(0).Type.In(2).String()) //参数的类型
	fmt.Println(reflect.ValueOf(reflect.Indirect(cc).Type().Name()).Type().Name())
	fmt.Println(reflect.Indirect(cc).Type().Name())

	//apiList := getApis(c)
	//fmt.Println(apiList)
	fmt.Println("=============")
	tt := ct.Method(0).Type.In(2)
	fmt.Println(reflect.TypeOf(tt).Elem().Kind())
	fmt.Println(reflect.ValueOf(tt).Type().Kind())
	fmt.Println(reflect.ValueOf(tt).Elem().Kind())
	fmt.Println(reflect.ValueOf(tt).Elem().NumField())
	v := reflect.Indirect(reflect.ValueOf(tt))
	fmt.Printf("%#v", v)
	news := reflect.New(reflect.TypeOf(tt))
	fmt.Println(reflect.ValueOf(news).Type().NumField())
	fmt.Println(reflect.TypeOf(news).Field(0).Tag.Get("json"))

	//pa:=&interface{
	//	Dir:"/path",
	//	Command:"ps"}
	//var p map[string]string{}
	params := make(map[string]string)
	params["Dir"] = "/a/b/c"
	params["Command"] = "ps"
	fmt.Println(params)
	var sss interface{}
	fmt.Println(sss.(pt))
	for i, n := range params {
		fmt.Println(i, n)
		&sss{i: n}
	}
	//map[string]string{
	//	"Dir":"/a/b/c",
	//	"Command":"ps"
	//	}

	//强制提取字符串，在new出参数结构体，在赋值
}

func getApis(object sc.SpacecraftClient) []string {
	var re = []string{}
	ct := reflect.TypeOf(object)
	for i := 0; i < ct.NumMethod(); i++ {
		//fmt.Println(ct.Method(i).Name)
		re = append(re, ct.Method(i).Name)
	}

	return re
}
