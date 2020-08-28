package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Result struct {
	Quo, Rem int
}

func main() {
	// dialhttp 连接指定地址上的http rpc 服务器，监听默认的http rpc地址
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing failed :", err)
	}
	args := Args{
		A: 17,
		B: 8,
	}
	var reply int
	//client.Call()  调用远程方法(同步方法)
	//client.Go()  调用远程方法(异步方法)
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith:%d*%d=%d\n", args.A, args.B, reply)

	// Asyn call
	quot := new(Result)
	divCall := client.Go("Arith.Divide", args, quot, nil)
	replyCall := <-divCall.Done // will be equal to divCall
	if replyCall.Error != nil {
		log.Fatal("arith error:", replyCall.Error)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}
