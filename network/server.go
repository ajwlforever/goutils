package network

import (
	"fmt"
	"net"
)

func StartServer() {
	fmt.Println("Starting the server......")
	// 创建Listener

	listener, err := net.Listen("tcp", "localhost:8098")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return // end
	}

	//监听
	for {
		fmt.Println("trying to connect")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting")
			return // end
		}
		go doServerStuff(conn)
	}

}
func doServerStuff(conn net.Conn) {
	for {
		// 一直尝试读取连接中的数据
		buff := make([]byte, 512)
		len, err := conn.Read(buff)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}
		fmt.Println("Receive data:%v", string(buff[:len]))
	}
}
