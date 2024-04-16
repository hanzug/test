package rpc

import (
	"log"
	"net"
	"net/rpc"
)

// 1. 定义RPC服务端方法
type Args struct {
	A, B int
	Rf   Raft
}

type Raft struct {
	A, B int
}

type Reply struct {
	Result int
}

type Arith int

func (t *Raft) Multiply(args *Args, reply *Reply) error {
	reply.Result = args.A * args.B
	return nil
}

func (t *Arith) Multiply(args *Args, reply *Reply) error {
	reply.Result = args.A * args.B
	return nil
}

func server() {
	arith := new(Arith)
	rpc.Register(arith)

	// 3. 启动RPC服务端
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
