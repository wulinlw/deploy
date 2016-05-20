package main

import (
	sc "../spacecraft"
	"errors"
	"fmt"
	//"github.com/garyburd/redigo/redis"
	//	"github.com/gorilla/websocket"
	"encoding/json"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//var upgrader = websocket.Upgrader{
//	ReadBufferSize:  2048,
//	WriteBufferSize: 2048,
//	CheckOrigin:     func(r *http.Request) bool { return true },
//}
var responseChan chan string

type jsonresult struct {
	UniqueId string `json:"uniqueId"`
	Ip       string `json:"ip"`
	Command  string `json:"command"`
	Result   string `json:"result"`
}
type liveresult struct {
	Ip     string
	Mid    string
	Result int
}

func main() {
	responseChan = make(chan string, 100)
	go h.run()
	http.HandleFunc("/ws", serveWs)
	http.HandleFunc("/index", index)
	http.HandleFunc("/live", live)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	ip := req.FormValue("ip")
	conn, _, err := createConnect(ip)
	if conn != nil { //连接失败时为nil，不能close
		defer conn.Close()
	}
	if err != nil {
		log.Printf("connect failed: %v", err)
		log.Printf("connect failed:" + ip)
		return
	}
	//	redisConn, err := createConnectRedis()
	//	if redisConn != nil { //连接失败时为nil，不能close
	//		defer redisConn.Close()
	//	}
	//	if err != nil {
	//		log.Printf("connect redis failed: %v", err)
	//		log.Printf("connect redis failed:" + ip)
	//		return
	//	}

	clientObj := sc.NewSpacecraftClient(conn)
	cc := reflect.ValueOf(clientObj)

	apiList := getApis(clientObj)
	apiExist := checkApiExist(apiList, req.FormValue("apiName"))
	if !apiExist {
		fmt.Println("undefined apiName:" + req.FormValue("apiName"))
		return
	}
	tt := cc.MethodByName(req.FormValue("apiName"))
	paramsStruct := getStruct(req.FormValue("apiName"))
	paramsStruct = fillStruct(paramsStruct, req)
	fmt.Printf("\n======\n%s\n", req.FormValue("apiName"))
	fmt.Printf("%#v", paramsStruct)
	params := []interface{}{context.Background(), paramsStruct}
	if tt.IsValid() {
		args := make([]reflect.Value, len(params))
		for k, param := range params {
			args[k] = reflect.ValueOf(param)
		}
		// 调用
		ret := tt.Call(args)
		//fmt.Println(ret, ret[0].Kind(), ret[0].String(), ret[0].Elem().FieldByName("String_").String())
		//fmt.Printf("%#v", ret[0].Elem().FieldByName("String_").String())
		if ret[0].Kind() == reflect.String {
			fmt.Printf("%s ret[0].Elem().FieldByName(\"String_\") called result: %s\n", "方法名", ret[0].String())
		}

		response := ret[0].Elem().FieldByName("String_").String()
		if req.FormValue("apiName") == "ComplexCommand" {

			ipSlice := strings.Split(req.FormValue("ip"), ":")
			//			ip := ipSlice[0]
			//			command := req.FormValue("Command")
			//			fmt.Println(ip)
			//			fmt.Println(command)
			result := jsonresult{
				UniqueId: req.FormValue("uniqueId"),
				Ip:       ipSlice[0],
				Command:  req.FormValue("Command"),
				Result:   response,
			}
			fmt.Println(result)
			result_byte, err := json.Marshal(result)
			if err != nil {
				fmt.Println("json error:", err)
			}
			h.broadcast <- result_byte //websock 广播执行日志
		}
		w.Write([]byte(response))
	} else {
		fmt.Println("can't call ")
	}

}

func createConnect(ip string) (*grpc.ClientConn, sc.SpacecraftClient, error) {
	conn, err := grpc.Dial(ip, grpc.WithInsecure(), grpc.WithTimeout(time.Second*5), grpc.WithBlock())
	//defer conn.Close()
	c := sc.NewSpacecraftClient(conn)
	return conn, c, err
}

//func createConnectRedis() (redis.Conn, error) {
//	c, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialDatabase(0))
//	return c, err
//}

//func createConnectWebsocket() (websocket.Conn, error) {
//	conn, err := upgrader.Upgrade(w, r, nil)
//	return conn, err
//}

func getApis(object sc.SpacecraftClient) []string {
	var apiList = []string{}
	ct := reflect.TypeOf(object)
	for i := 0; i < ct.NumMethod(); i++ {
		apiList = append(apiList, ct.Method(i).Name)
	}
	return apiList
}

func checkApiExist(apiList []string, apiName string) bool {
	for _, v := range apiList {
		if v == apiName {
			return true
		}
	}
	return false
}

func fillStruct(s interface{}, req *http.Request) interface{} {
	//s := &sa{}
	//	s.Aaa = "aaa"
	for i := 0; i < reflect.ValueOf(s).Elem().NumField(); i++ {
		fieldName := reflect.TypeOf(s).Elem().Field(i).Name
		fieldValue := reflect.ValueOf(s).Elem().FieldByName(fieldName)
		valueType := fieldValue.Type().Kind().String()
		fmt.Println(fieldName, fieldValue, fieldValue.Type().Kind())
		//		fmt.Println(reflect.ValueOf(s).FieldByName(fieldName))
		if !fieldValue.CanSet() {
			log.Println("不可设置值，struct对象field不可设置")
		}
		if req.FormValue(fieldName) != "" {
			tmp := req.FormValue(fieldName)
			value, err := TypeConversion(tmp, valueType)
			if err != nil {
				log.Println(err)
			}
			fieldValue2 := value
			fieldValue.Set(fieldValue2)
		}
	}
	fmt.Sprintf("%#v", s)
	return s
}

func getStruct(apiName string) interface{} {
	switch apiName {
	case "ComplexCommand":
		return &sc.SpecifiedCommandParams{}
	case "GetFileList":
		return &sc.SvnUpParam{}
	case "Live":
		return &sc.Empty{}
	case "SendFile":
		return &sc.SendFileParams{}
	case "SpecifiedCommand":
		return &sc.SpecifiedCommandParams{}
	case "SvnCheckout":
		return &sc.SvnCheckoutParams{}
	case "SvnInfo":
		return &sc.SvnUpParam{}
	case "SvnUp":
		return &sc.SvnUpParam{}
	case "SvnUpToRevision":
		return &sc.SvnUpToRevisionParams{}
	default:
		return nil
	}
	return nil
}

//类型转换
func TypeConversion(value string, ntype string) (reflect.Value, error) {
	if ntype == "string" {
		return reflect.ValueOf(value), nil
	} else if ntype == "time.Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "int" {
		i, err := strconv.Atoi(value)
		return reflect.ValueOf(i), err
	} else if ntype == "int8" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int8(i)), err
	} else if ntype == "int32" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int64(i)), err
	} else if ntype == "int64" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(i), err
	} else if ntype == "float32" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(float32(i)), err
	} else if ntype == "float64" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(i), err
	} else if ntype == "slice" {
		i := []byte(value)
		return reflect.ValueOf(i), nil
	}
	//else if .......增加其他一些类型的转换
	return reflect.ValueOf(value), errors.New("未知的类型：" + ntype)
}

func live(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	liveChan := make(chan *liveresult, 1024)
	for k, _ := range req.Form {
		fmt.Println(req.FormValue(k))
		go liveCheck(req.FormValue(k), liveChan)
	}

	var results []*liveresult
	timeOut := time.NewTimer(time.Second * 1)
	for {
		select {
		case info := <-liveChan:
			results = append(results, info)
		case <-timeOut.C:
			result_byte, err := json.Marshal(results)
			if err != nil {
				fmt.Println("json error:", err)
			}
			w.WriteHeader(http.StatusOK)
			w.Write(result_byte)
			close(liveChan)
			return
		}
	}

}
func liveCheck(machineInfo string, liveChan chan *liveresult) bool {
	ipSlice := strings.Split(machineInfo, ":")
	ip := ipSlice[0]
	mid := ipSlice[1]
	var result int = 0
	status := &liveresult{
		Ip:     ip,
		Mid:    mid,
		Result: result,
	}

	conn, c, err := createConnect(ip + ":50051")
	if conn != nil { //连接失败时为nil，不能close
		defer conn.Close()
	}
	if err != nil {
		log.Printf("connect failed %s : %v", ip, err)
		return false
	}
	param := &sc.Empty{}
	r, err := c.Live(context.Background(), param)
	if err != nil {
		log.Printf("could not call live: %v", err)
		return false
	}
	if string(r.String_) == "ok" {
		status.Result = 1
	}

	liveChan <- status
	return true
}
