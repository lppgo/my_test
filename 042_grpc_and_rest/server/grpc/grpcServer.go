/*
@File    :   grpcServer.go
@Time    :   2022/02/08 10:13:28
@Author  :   lpp
@Version :   1.0.0
@Contact :   golpp@qq.com
@Desc    :   grpc服务
*/

package grpcServer

import (
	"context"
	"net"

	"grpc-rest/protos"

	"google.golang.org/grpc"
)

//
type OrderServiceImpl struct {
}

func (OrderServiceImpl) Create(context.Context, *protos.CreateOrderRequest) (*protos.CreateOrderResponse, error) {
	return nil, nil
}
func (OrderServiceImpl) Retrieve(context.Context, *protos.RetrieveOrderRequest) (*protos.RetrieveOrderResponse, error) {
	return nil, nil
}
func (OrderServiceImpl) Update(context.Context, *protos.UpdateOrderRequest) (*protos.UpdateOrderResponse, error) {
	return nil, nil
}
func (OrderServiceImpl) Delete(context.Context, *protos.DeleteOrderRequest) (*protos.DeleteOrderResponse, error) {
	return nil, nil
}
func (OrderServiceImpl) List(context.Context, *protos.ListOrderRequest) (*protos.ListOrderResponse, error) {
	return nil, nil
}

// GrpcServer 为订单服务实现 gRPC 服务
type GrpcServer struct {
	server   *grpc.Server
	listener net.Listener
	errCh    chan error
}

//NewGrpcServer 是一个创建 GrpcServer 的便捷函数
func NewGrpcServer(service protos.OrderServiceServer, port string) (GrpcServer, error) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return GrpcServer{}, err
	}
	server := grpc.NewServer()
	orderService := OrderServiceImpl{}
	protos.RegisterOrderServiceServer(server, &orderService)
	return GrpcServer{
		server:   server,
		listener: lis,
		errCh:    make(chan error),
	}, nil
}

// Start 在后台启动服务，将任何错误传入错误通道
func (g GrpcServer) Start() {
	go func() {
		g.errCh <- g.server.Serve(g.listener)
	}()
}

// Stop 停止 gRPC 服务
func (g GrpcServer) Stop() {
	g.server.GracefulStop()
}

//Error 返回服务的错误通道
func (g GrpcServer) Error() chan error {
	return g.errCh
}
