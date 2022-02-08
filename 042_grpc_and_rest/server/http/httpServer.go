/*
@File    :   httpServer.go
@Time    :   2022/02/08 10:13:42
@Author  :   lpp
@Version :   1.0.0
@Contact :   golpp@qq.com
@Desc    :   RESTFUL 服务
*/

package httpServer

import (
	"grpc-rest/protos"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
)

// RestServer 为订单服务实现了一个 REST 服务。
type RestServer struct {
	server       *http.Server
	orderService protos.OrderServiceServer // 与我们注入 gRPC 服务端的订单服务相同
	errCh        chan error
}

// NewRestServer 是一个创建 RestServer 的便捷函数
func NewRestServer(orderService protos.OrderServiceServer, port string) RestServer {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	rs := RestServer{
		server: &http.Server{
			Addr:    ":" + port,
			Handler: router,
		},
		orderService: orderService,
		errCh:        make(chan error),
	}
	// 注册路由
	router.POST("/order", rs.create)
	router.GET("/order/:id", rs.retrieve)
	router.PUT("/order", rs.update)
	router.DELETE("/order", rs.delete)
	router.GET("/order", rs.list)
	return rs
}

// Start 在后台启动 REST 服务，将错误推入错误通道
func (r RestServer) Start() {
	go func() {
		r.errCh <- r.server.ListenAndServe()
	}()
}

// Stop 停止服务
func (r RestServer) Stop() error {
	return r.server.Close()
}

// Error 返回服务端的错误通道
func (r RestServer) Error() chan error {
	return r.errCh
}

// create 是一个处理函数，它根据订单请求创建订单 (JSON 主体)
func (r RestServer) create(c *gin.Context) {
	var req protos.CreateOrderRequest
	// unmarshal 订单请求
	err := jsonpb.Unmarshal(c.Request.Body, &req)
	if err != nil {
		c.String(http.StatusInternalServerError, "error creating order request")
	}
	// 根据请求，使用订单服务创建订单
	resp, err := r.orderService.Create(c.Request.Context(), &req)
	if err != nil {
		c.String(http.StatusInternalServerError, "error creating order")
	}
	m := &jsonpb.Marshaler{}
	if err := m.Marshal(c.Writer, resp); err != nil {
		c.String(http.StatusInternalServerError, "error sending order response")
	}
}
func (r RestServer) retrieve(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}
func (r RestServer) update(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}
func (r RestServer) delete(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}
func (r RestServer) list(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}
