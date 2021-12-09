package main

import (
	"context"
	"io"
	"time"

	pb "go-micro-v2-jaeger/protos"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"

	wrapperTrace "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

var (
	serverName        = "go.micro.srv.greeter"
	tracerServiceName = "tracer-srv"
	tracerAddress     = "localhost:6831"
)

type Greeter struct{}

func (s *Greeter) Hello(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	log.Log("Received Say.Hello request")
	rsp.Greeting = "Hello " + req.Name

	// return errors.New("process error !")
	return nil
}

func main() {
	// tracer
	tracer, io, err := NewTracer(tracerServiceName, tracerAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(tracer)

	// service
	service := micro.NewService(
		micro.Name(serverName),
		micro.WrapHandler(wrapperTrace.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	pb.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
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
