package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

type Arith struct {
}

type ArithRequest struct {
	A int
	B int
}

type ArithResponse struct {
	Pro int //乘积
	Quo int //商，除法
	Rem int //余数

}

func (c *Arith) Multiply(req ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}
func (c *Arith) Division(req ArithRequest, res *ArithResponse) error {
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}
func init() {
	fmt.Println("json rpc 采用了json编码，而不是gob编码！服务端代码如下")
}
func main() {
	//注册rpc服务
	arith := new(Arith)
	rpc.Register(arith)

	listen, err := net.Listen("tcp", ":1236")
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Fprintf(os.Stdout, "%s", "start connection ... ")

	for {
		//接收客户端连接请求
		conn, err := listen.Accept()
		if err != nil {
			continue
		}

		//并发处理客户端请求
		go func(conn net.Conn) {
			fmt.Fprintf(os.Stdout, "%s", "new client in comming\n")
			jsonrpc.ServeConn(conn)
		}(conn)
	}

}
