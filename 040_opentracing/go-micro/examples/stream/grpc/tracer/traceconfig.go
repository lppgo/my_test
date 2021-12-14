package tracer

import (
	"context"
	"fmt"
	"io"
	"runtime"
	"strings"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc/metadata"
)

const (
	// _disable                        = false
	// _rpc_metrics                    = false
	_agent_addr    = "localhost:6831"
	_sampler_param = float64(1)
	// _reporter_queue_size            = 100
	// _reporter_buffer_flush_interval = time.Second
)

type TracerCfg struct {
	ServiceName                 string
	AgentAddr                   string
	Disable                     bool
	SamplerType                 string
	SamplerParam                float64
	ReporterQueueSize           int
	ReporterBufferFlushInterval time.Duration
}

func TraceInit(tracerCfg TracerCfg) (opentracing.Tracer, io.Closer) {
	if tracerCfg.SamplerType == "" {
		tracerCfg.SamplerType = jaeger.SamplerTypeConst
	}
	if tracerCfg.AgentAddr == "" {
		tracerCfg.AgentAddr = _agent_addr
	}

	cfg := &config.Configuration{
		ServiceName: tracerCfg.ServiceName,
		Sampler: &config.SamplerConfig{
			Type:  tracerCfg.SamplerType,
			Param: tracerCfg.SamplerParam,
		},
		Reporter: &config.ReporterConfig{
			LocalAgentHostPort:  tracerCfg.AgentAddr,
			LogSpans:            true,
			BufferFlushInterval: time.Second * 1,
			// QueueSize: 100,
		},
	}

	sender, err := jaeger.NewUDPTransport(tracerCfg.AgentAddr, 0)
	if err != nil {
		panic(fmt.Sprintf("NewUDPTransport failed: %v\n", err))
	}

	tracer, closer, err := cfg.NewTracer(
		config.Reporter(jaeger.NewRemoteReporter(sender)),
		config.Logger(jaeger.StdLogger),
		config.PoolSpans(true))
	if err != nil {
		panic(fmt.Sprintf("Init failed: %v\n", err))
	}

	return tracer, closer
}

/*
	ctx := context.WithValue(pctx, "originID", "originID-001")
	ctx = context.WithValue(ctx, "broker", "broker-001")
	ctx = context.WithValue(ctx, "account", "account-001")

*/
func NewSpan(ctx context.Context) opentracing.Span {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	span := opentracing.StartSpan(
		runtime.FuncForPC(pc[0]).Name(),
		opentracing.ChildOf(opentracing.SpanFromContext(ctx).Context()),
	)
	//
	if sc, ok := span.Context().(jaeger.SpanContext); ok {
		span.SetTag("traceID", sc.TraceID())
	}

	//
	if ctx.Value("originID") != nil {
		span.SetTag("originID", ctx.Value("originID"))
	}
	if ctx.Value("broker") != nil {
		span.SetTag("broker", ctx.Value("broker"))
	}
	if ctx.Value("account") != nil {
		span.SetTag("account", ctx.Value("account"))
	}

	return span
}

// metadataReaderWriter satisfies both the opentracing.TextMapReader and
// opentracing.TextMapWriter interfaces.
type metadataReaderWriter struct {
	metadata.MD
}

func (w metadataReaderWriter) Set(key, val string) {
	// The GRPC HPACK implementation rejects any uppercase keys here.
	//
	// As such, since the HTTP_HEADERS format is case-insensitive anyway, we
	// blindly lowercase the key (which is guaranteed to work in the
	// Inject/Extract sense per the OpenTracing spec).
	key = strings.ToLower(key)
	w.MD[key] = append(w.MD[key], val)
}

func (w metadataReaderWriter) ForeachKey(handler func(key, val string) error) error {
	for k, vals := range w.MD {
		for _, v := range vals {
			if err := handler(k, v); err != nil {
				return err
			}
		}
	}

	return nil
}

func ExtractSpanContext(ctx context.Context, tracer opentracing.Tracer) (opentracing.SpanContext, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}
	return tracer.Extract(opentracing.HTTPHeaders, metadataReaderWriter{md})
}
