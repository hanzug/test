package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 在本地3000端口监听
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}
	defer ln.Close()
	fmt.Println("Server is listening on port 3000...")

	// 接受连接
	conn, err := ln.Accept()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}
	defer conn.Close()

	// 发送消息到客户端
	_, err = conn.Write([]byte("Hello from server!"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error sending message: %v\n", err)
		return
	}
}
