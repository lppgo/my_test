/*
@File    :   main.go
@Time    :   2022/02/08 09:46:35
@Author  :   lpp
@Version :   1.0.0
@Contact :   golpp@qq.com
@Desc    :
*/

package main

import (
	grpcServer "grpc-rest/server/grpc"
	httpServer "grpc-rest/server/http"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	grpcPort = "50051"
	restPort = "8080"
)

//app 是一个便捷的封装，用于启动和关闭订单微服务所需的所有东西
type app struct {
	restServer httpServer.RestServer
	grpcServer grpcServer.GrpcServer
	/* Listens for an application termination signal
	   Ex. (Ctrl X, Docker container shutdown, etc) */
	shutdownCh chan os.Signal
}

// start 在后台启动 REST 和 gRPC 服务
func (a app) start() {
	a.restServer.Start() // non blocking now
	a.grpcServer.Start() // also non blocking :-)
}

// stop 关闭服务
func (a app) shutdown() error {
	a.grpcServer.Stop()
	return a.restServer.Stop()
}

// newApp 使用 REST 和 gRPC 服务创建一个新的应用程序
// 这个函数执行所有与应用程序相关的初始化
func newApp() (app, error) {
	orderService := grpcServer.OrderServiceImpl{} //
	gs, err := grpcServer.NewGrpcServer(orderService, grpcPort)
	if err != nil {
		return app{}, err
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	return app{
		restServer: httpServer.NewRestServer(orderService, restPort),
		grpcServer: gs,
		shutdownCh: quit,
	}, nil
}

// 运行启动应用程序，处理任何 REST 或 gRPC 服务的错误以及任何关机的信号
func run() error {
	app, err := newApp()
	if err != nil {
		return err
	}
	app.start()
	defer app.shutdown()
	select {
	case restErr := <-app.restServer.Error():
		return restErr
	case grpcErr := <-app.grpcServer.Error():
		return grpcErr
	case <-app.shutdownCh:
		return nil
	}
}
func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
