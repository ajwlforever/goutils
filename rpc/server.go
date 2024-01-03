package rpc1

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

// 服务端开启服务
func InintServer() {
	//  rpc 服务器端开启
	math := new(MathService)
	// rpc注册服务
	rpc.Register(math)
	// rpc
	rpc.HandleHTTP()
	fmt.Printf("server start ....") // rpc  处理http
	// listener, err := net.Listen("tcp", "localhost:1234") // 监听1234端口的tcp请求
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Fatal("Error listening")
	}
	// go http.Serve(listener, nil)

	fmt.Printf("server stop ....")
}
