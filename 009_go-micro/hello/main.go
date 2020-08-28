package main

import (
	"hello/handler"
	"hello/subscriber"
	"time"

	"github.com/micro/go-micro/v2/registry"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry/etcd"

	hello "hello/proto/hello"
)

var etcdReg registry.Registry

func init() {
	etcdReg = etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
		registry.Timeout(10*time.Second),
	)
}
func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.hello"),
		micro.Version("latest"),
		micro.Registry(etcdReg),
	)

	// Initialise service
	service.Init()

	// Register Handler
	hello.RegisterHelloHandler(service.Server(), new(handler.Hello))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.hello", service.Server(), new(subscriber.Hello))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
