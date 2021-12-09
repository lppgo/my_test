package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "go-micro-v2-jaeger/protos"

	"github.com/micro/go-micro/v2"

	wrapperTrace "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

var (
	serverName        = "go.micro.srv.greeter"
	tracerServiceName = "tracer-cli"
	tracerAddress     = "localhost:6831"
)

func main() {
	// tracer
	tracer, io, err := NewTracer(tracerServiceName, tracerAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(tracer)
	// ctx, span, err := wrapperTrace.StartSpanFromContext(context.Background(), opentracing.GlobalTracer(), "root")

	// create a new service
	service := micro.NewService(
		micro.WrapClient(wrapperTrace.NewClientWrapper(opentracing.GlobalTracer())),
	)

	// parse command line flags
	service.Init()

	// Use the generated client stub
	cl := pb.NewGreeterService(serverName, service.Client())

	// Make request
	rsp, err := cl.Hello(context.Background(), &pb.Request{
		Name: "Lucas",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Greeting)
}

// NewTracer 创建一个jaeger Tracer
func NewTracer(servicename string, addr string) (opentracing.Tracer, io.Closer, error) {
	cfg := jaegercfg.Configuration{
		ServiceName: servicename,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}

	sender, err := jaeger.NewUDPTransport(addr, 0)
	if err != nil {
		return nil, nil, err
	}

	reporter := jaeger.NewRemoteReporter(sender)
	// Initialize tracer with a logger and a metrics factory
	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Reporter(reporter),
	)

	return tracer, closer, err
}
