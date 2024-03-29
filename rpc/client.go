package rpc

import (
	"fmt"
	"log"
	"net/rpc"
)

// 4. 定义RPC客户端
func client() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer client.Close()

	// 5. 调用RPC方法
	args := &Args{7, 8, Raft{}}
	var reply Reply
	err = client.Call("Arith.Raft.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply.Result)
}
