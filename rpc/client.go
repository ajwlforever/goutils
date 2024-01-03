package rpc1

import (
	"fmt"
	"log"
	"net/rpc"
)

func InitClient() {
	// 得由sevice的基本接口
	var args = Args{A: 32, B: 22}
	var result = Result{C: 0}

	// client 连接到服务器端口
	client, err := rpc.DialHTTP("tcp", "localhost:1234") //HTTP拨打 loacalhost的1234端口
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result:", result)
	//rpc Call
	err = client.Call("MathService.AddService", args, &result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result:", result)
}
