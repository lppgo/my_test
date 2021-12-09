# 1: Summary

An example of using opentracing.

This example shows how to trace over the process boundaries and RPC calls.

# 2: jaeger

Jaeger start

    docker run \
    -p 5775:5775/udp \
    -p 16686:16686 \
    -p 6831:6831/udp \
    -p 6832:6832/udp \
    -p 5778:5778 \
    -p 14268:14268 \
    jaegertracing/all-in-one:latest

# 3: jaeger client 代码示例

Client side

<1>import

    import (
        "github.com/opentracing/opentracing-go/ext"
    )

<2>inject

    ext.SpanKindRPCClient.Set(reqSpan)
    ext.HTTPUrl.Set(reqSpan, reqURL)
    ext.HTTPMethod.Set(reqSpan, "GET")
    span.Tracer().Inject(
        span.Context(),
        opentracing.HTTPHeaders,
        opentracing.HTTPHeadersCarrier(req.Header),
    )

Server side

<1>import

    import (
        opentracing "github.com/opentracing/opentracing-go"
        "github.com/opentracing/opentracing-go/ext"
        otlog "github.com/opentracing/opentracing-go/log"
        "github.com/yurishkuro/opentracing-tutorial/go/lib/tracing"
    )

<2>Extrace span context from incoming http request

    spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))

<3>Creates a ChildOf reference to the passed spanCtx as well as sets a span.kind=server tag on the new span by using a special option RPCServerOption

    span := tracer.StartSpan("format", ext.RPCServerOption(spanCtx))
    defer span.Finish()

# 4： 压测

## 4.1 web server

- `hey -n 1000000 -c 100 http://localhost:8080/getList`

### 4.1.1 不开启 jaeger 采样

```json
$ hey -n 1000000 -c 100  http://localhost:8080/getList

Summary:
  Total:        41.4929 secs
  Slowest:      0.0641 secs
  Fastest:      0.0001 secs
  Average:      0.0041 secs
  Requests/sec: 24100.4969

  Total data:   16000000 bytes
  Size/request: 16 bytes

Response time histogram:
  0.000 [1]     |
  0.006 [794137]        |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.013 [169465]        |■■■■■■■■■
  0.019 [30348] |■■
  0.026 [4875]  |
  0.032 [928]   |
  0.038 [205]   |
  0.045 [30]    |
  0.051 [9]     |
  0.058 [1]     |
  0.064 [1]     |


Latency distribution:
  10% in 0.0006 secs
  25% in 0.0013 secs
  50% in 0.0030 secs
  75% in 0.0057 secs
  90% in 0.0092 secs
  95% in 0.0117 secs
  99% in 0.0174 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0000 secs, 0.0001 secs, 0.0641 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0069 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0292 secs
  resp wait:    0.0037 secs, 0.0000 secs, 0.0629 secs
  resp read:    0.0003 secs, 0.0000 secs, 0.0338 secs

Status code distribution:
  [200] 1000000 responses

```

## 4.1.2 开启全量采集

- 采集设置

```go
tracer, closer = traceconfig.TraceInit("web-tracer-server", "const", 1)
```

```json
$ hey -n 1000000 -c 100  http://localhost:8080/getList

Summary:
  Total:        81.3497 secs
  Slowest:      0.1411 secs
  Fastest:      0.0001 secs
  Average:      0.0081 secs
  Requests/sec: 12292.6010

  Total data:   16000000 bytes
  Size/request: 16 bytes

Response time histogram:
  0.000 [1]     |
  0.014 [844719]        |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.028 [137985]        |■■■■■■■
  0.042 [14664] |■
  0.056 [2127]  |
  0.071 [381]   |
  0.085 [108]   |
  0.099 [10]    |
  0.113 [3]     |
  0.127 [0]     |
  0.141 [2]     |


Latency distribution:
  10% in 0.0011 secs
  25% in 0.0030 secs
  50% in 0.0065 secs
  75% in 0.0112 secs
  90% in 0.0169 secs
  95% in 0.0211 secs
  99% in 0.0321 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0000 secs, 0.0001 secs, 0.1411 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0030 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0355 secs
  resp wait:    0.0078 secs, 0.0000 secs, 0.1215 secs
  resp read:    0.0002 secs, 0.0000 secs, 0.0384 secs

Status code distribution:
  [200] 1000000 responses
```

### 4.1.3 probabilistic 概率采集

- 采集设置

```go
tracer, closer = traceconfig.TraceInit("web-tracer-server", "probabilistic", 0.5)
```

- 压测结果

```json
$ hey -n 1000000 -c 100  http://localhost:8080/getList

Summary:
  Total:        81.9106 secs
  Slowest:      0.1619 secs
  Fastest:      0.0001 secs
  Average:      0.0081 secs
  Requests/sec: 12208.4377

  Total data:   16000000 bytes
  Size/request: 16 bytes

Response time histogram:
  0.000 [1]     |
  0.016 [865201]        |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.032 [116499]        |■■■■■
  0.049 [14838] |■
  0.065 [2658]  |
  0.081 [506]   |
  0.097 [149]   |
  0.113 [109]   |
  0.130 [30]    |
  0.146 [8]     |
  0.162 [1]     |


Latency distribution:
  10% in 0.0009 secs
  25% in 0.0022 secs
  50% in 0.0056 secs
  75% in 0.0115 secs
  90% in 0.0185 secs
  95% in 0.0240 secs
  99% in 0.0379 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0000 secs, 0.0001 secs, 0.1619 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0017 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0425 secs
  resp wait:    0.0076 secs, 0.0000 secs, 0.1618 secs
  resp read:    0.0004 secs, 0.0000 secs, 0.0411 secs

Status code distribution:
  [200] 1000000 responses
```
