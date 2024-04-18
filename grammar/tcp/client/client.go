package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	// 连接到服务器
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}
	defer conn.Close()

	// 读取数据
	message, err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading: %v\n", err)
		return
	}

	fmt.Println("Message from server:", string(message))
}
