package main

import (
	"context"
	"log"

	hello "github.com/asim/go-micro/examples/v4/greeter/srv/proto/hello"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
)

func main() {
	CallGrpcByHttp()
}

func CallGrpcByHttp() {
	// create a new service
	service := micro.NewService(
		micro.Registry(registry.NewRegistry()),
	)
	// parse command line flags
	service.Init()
	c := service.Client()

	req := &hello.Request{Name: "call grpc server by http client"}
	// create request/response
	request := client.NewRequest("go.micro.srv.greeter", "Say.Hello", req)

	response := new(hello.Response)
	// call service
	err := c.Call(context.Background(), request, response)
	log.Printf("err:%v response:%#v\n", err, response)
}
