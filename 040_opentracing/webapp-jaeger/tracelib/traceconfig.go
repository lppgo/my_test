package traceconfig

import (
	"fmt"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func TraceInit(serviceName string, samplerType string, samplerParam float64) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  samplerType,
			Param: samplerParam, //
		},
		Reporter: &config.ReporterConfig{
			LocalAgentHostPort: "localhost:6831",
			LogSpans:           true,
		},
	}

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("Init failed: %v\n", err))
	}

	return tracer, closer
}

/*

其中关于SamplerConfig的Type可以选择
	const，全量采集。param采样率设置1,0 分别对应打开和关闭
	probabilistic ，概率采集。param默认万份之一，0~1之间取值，
	rateLimiting ，限速采集。param每秒采样的个数
	remote 动态采集策略。param值于probabilistic的参数一样。
	在收到实际值之前的初始采样率。改值可以通过环境变量的JAEGER_SAMPLER_PARAM设定
*/
