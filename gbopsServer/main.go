package main

import (
	"log"
	"net"

	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"

	"../deploy/command"
	"../deploy/upload"
	sc "../spacecraft"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port        = ":50051"
	SysInfoTime = 30
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
func (s *server) ComplexCommand(ctx context.Context, in *sc.SpecifiedCommandParams) (*sc.ResponseStr, error) {
	result := command.ComplexCommand(in.Command, in.Dir)
	return &sc.ResponseStr{String_: result}, nil
}
func (s *server) SendFile(ctx context.Context, in *sc.SendFileParams) (*sc.ResponseStr, error) {
	upload.Upload(in.RelativePath, in.FileContent)
	return &sc.ResponseStr{String_: "ok"}, nil
}
func (s *server) GetFileList(ctx context.Context, in *sc.SvnUpParam) (*sc.ResponseStr, error) {
	result := upload.GetFileList(in.Dir)
	return &sc.ResponseStr{String_: result}, nil
}
func (s *server) Live(ctx context.Context, in *sc.Empty) (*sc.ResponseStr, error) {
	return &sc.ResponseStr{String_: "ok"}, nil
}

func main() {
	go KeepHeartBeat()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	sc.RegisterSpacecraftServer(s, &server{})
	s.Serve(lis)
}

func KeepHeartBeat() {
	//初始化定时器
	t := time.NewTicker(SysInfoTime * time.Second)
	for {
		select {
		case <-t.C:
			SysInfoUpload()
		}
	}
}

func SysInfoUpload() {
	defer func() { //必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println("SysInfoUpload error", err) //这里的err其实就是panic传入的内容，"bug"
		}
	}()
	memInfo, _ := mem.VirtualMemory()
	cpuInfo, _ := cpu.Percent(time.Second*SysInfoTime, false)
	// almost every return value is a struct
	//fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", memInfo.Total, memInfo.Free, memInfo.UsedPercent)
	// convert to JSON. String() is also implemented
	//fmt.Println(memInfo)
	//fmt.Println(cpuInfo)

	p := url.Values{}
	p.Set("memTotal", strconv.Itoa(int(memInfo.Total/(1024*1024))))
	p.Set("memFree", strconv.Itoa(int(memInfo.Free/(1024*1024))))
	p.Set("memPercent", strconv.Itoa(int(memInfo.UsedPercent)))
	p.Set("cpuPercent", strconv.Itoa(int(cpuInfo[0])))
	body := ioutil.NopCloser(strings.NewReader(p.Encode())) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://gbops.gamebar.com/collect/sysinfo", body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时
	//fmt.Println(p)
	resp, _ := client.Do(req) //发送
	defer resp.Body.Close()   //一定要关闭resp.Body
	//data, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(data))
}
