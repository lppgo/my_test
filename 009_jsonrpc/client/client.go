package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type ArithRequest struct {
	A int
	B int
}

type ArithResponse struct {
	Pro int //乘积
	Quo int //商
	Rem int //余数
}

func main() {
	client, err := jsonrpc.Dial("tcp", "127.0.0.1:1236")
	if err != nil {
		log.Fatalln(err.Error())
	}

	req := ArithRequest{9, 2}
	var res ArithResponse

	err = client.Call("Arith.Multiply", req, &res)
	if err != nil {
		log.Fatalln("arith.Multiply error:", err)
	}
	fmt.Printf("%d*%d=%d\n", req.A, req.B, res.Pro)

	divCall := client.Go("Arith.Division", req, &res, nil)
	divReply := <-divCall.Done
	if divReply.Error != nil {
		log.Fatalln("arith.Division error:", err)
	}
	fmt.Printf("%d / %d, quo is %d, rem is %d\n", req.A, req.B, res.Quo, res.Rem)
}
