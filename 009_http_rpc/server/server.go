package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

//传入参数
type Args struct {
	A, B int
}

//返回参数
type Result struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Result) error {
	if args.B == 0 {
		return errors.New("divide bu zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
func init() {
	fmt.Println("基于HTTP协议实现的RPC，服务端的代码如下")
}
func main() {
	arith := new(Arith)
	//给arith对象注册了一个rpc服务
	rpc.Register(arith)
	//把rpc挂载到http服务上面,开启服务
	//当http服务打开的时候我们就可以通过rpc客户端来调用arith中符合rpc标准的的方法了。
	rpc.HandleHTTP()
	//
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

/*
golang的rpc支持三个级别的RPC：TCP、HTTP、JSONRPC

运行时一次客户机对服务器的RPC调用步骤有：
	调用客户端句柄，执行传送参数
	调用客户端系统内核发送网络信息
	消息传送到远程主机
	服务器句柄得到消息并取得参数
	执行远程过程
	执行的过程将结果返回服务器句柄
	服务器句柄返回结果，调用远程系统内核
	消息传回本地主机
	客户句柄由内核接收消息
	客户接收句柄返回的数据
（2）调用RPC服务
有两种方式：同步或异步
  同步方式client.Call("rpc上的公开类名:公开方法", 第一个传入的变量, 第二个传入的变量)
  异步方式divCall := client.Go("rpc上的公开类名:公开方法", 第一次传入的变量, 第二个传入的变量, nil)
  replyCall := <- divCall.Done        阻塞，等待异步完成

*/
