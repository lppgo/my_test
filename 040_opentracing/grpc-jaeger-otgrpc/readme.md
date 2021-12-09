[toc]

# 1: 说明

# 2：jeager to client

# 3: jaeger to server

# 4: jaeger 采集设置

# 5：trace 对性能影响压测

## 5.1 unary grpc

### 5.1.1 对 origin unary grpc 进行压测(不带业务搜索逻辑)

- (1) 不采集

```json
$ ghz -insecure \
--proto ./protos/routeguide.proto \
--call routeguide.RouteGuide.GetFeature \
--skipTLS \
--insecure \
-d '{"latitude":409146138,"longitude":-746188906}' \
-c 100 \
-n 100000 \
localhost:50051

Summary:
  Count:        100000
  Total:        15.02 s
  Slowest:      119.34 ms
  Fastest:      0.15 ms
  Average:      6.13 ms
  Requests/sec: 6658.41

Response time histogram:
  0.155   [1]     |
  12.073  [88263] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  23.992  [10417] |∎∎∎∎∎
  35.910  [1016]  |
  47.829  [222]   |
  59.748  [61]    |
  71.666  [15]    |
  83.585  [3]     |
  95.504  [1]     |
  107.422 [0]     |
  119.341 [1]     |

Latency distribution:
  10 % in 1.18 ms
  25 % in 2.29 ms
  50 % in 4.71 ms
  75 % in 8.42 ms
  90 % in 12.77 ms
  95 % in 15.37 ms
  99 % in 26.97 ms

Status code distribution:
  [OK]   100000 responses

```

- 全量采集

```json
$ ghz -insecure \
--proto ./protos/routeguide.proto \
--call routeguide.RouteGuide.GetFeature \
--skipTLS \
--insecure \
-d '{"latitude":409146138,"longitude":-746188906}' \
-c 100 \
-n 100000 \
localhost:50051

Summary:
  Count:        100000
  Total:        27.56 s
  Slowest:      177.73 ms
  Fastest:      0.21 ms
  Average:      12.46 ms
  Requests/sec: 3627.79

Response time histogram:
  0.212   [1]     |
  17.963  [78096] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  35.715  [18026] |∎∎∎∎∎∎∎∎∎
  53.467  [3020]  |∎∎
  71.218  [612]   |
  88.970  [147]   |
  106.722 [75]    |
  124.473 [19]    |
  142.225 [1]     |
  159.976 [0]     |
  177.728 [3]     |

Latency distribution:
  10 % in 2.53 ms
  25 % in 4.94 ms
  50 % in 9.61 ms
  75 % in 16.75 ms
  90 % in 25.59 ms
  95 % in 32.78 ms
  99 % in 51.65 ms

Status code distribution:
  [OK]   100000 responses
```

### 5.1.2 对 origin unary grpc 进行压测(带业务搜索逻辑)

- (1) 不采集

```json
$ ghz -insecure \
--proto ./protos/routeguide.proto \
--call routeguide.RouteGuide.GetFeature \
--skipTLS \
--insecure \
-d '{"latitude":409146138,"longitude":-746188906}' \
-c 100 \
-n 100000 \
localhost:50051

Summary:
  Count:        100000
  Total:        12.86 s
  Slowest:      155.17 ms
  Fastest:      0.17 ms
  Average:      5.81 ms
  Requests/sec: 7774.98

Response time histogram:
  0.167   [1]     |
  15.668  [95154] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  31.168  [3027]  |∎
  46.669  [1125]  |
  62.169  [539]   |
  77.670  [69]    |
  93.170  [61]    |
  108.671 [5]     |
  124.172 [15]    |
  139.672 [1]     |
  155.173 [3]     |

Latency distribution:
  10 % in 1.08 ms
  25 % in 2.04 ms
  50 % in 3.96 ms
  75 % in 7.03 ms
  90 % in 10.80 ms
  95 % in 15.32 ms
  99 % in 40.99 ms

Status code distribution:
  [OK]   100000 responses

```

- (2) 全量采集

```json
$ ghz -insecure \
--proto ./protos/routeguide.proto \
--call routeguide.RouteGuide.GetFeature \
--skipTLS \
--insecure \
-d '{"latitude":409146138,"longitude":-746188906}' \
-c 100 \
-n 100000 \
localhost:50051

Summary:
  Count:        100000
  Total:        30.68 s
  Slowest:      199.20 ms
  Fastest:      0.24 ms
  Average:      14.78 ms
  Requests/sec: 3259.86

Response time histogram:
  0.243   [1]     |
  20.139  [74892] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  40.035  [21074] |∎∎∎∎∎∎∎∎∎∎∎
  59.931  [3158]  |∎∎
  79.827  [639]   |
  99.723  [150]   |
  119.619 [61]    |
  139.514 [20]    |
  159.410 [2]     |
  179.306 [2]     |
  199.202 [1]     |

Latency distribution:
  10 % in 3.02 ms
  25 % in 5.83 ms
  50 % in 11.54 ms
  75 % in 20.19 ms
  90 % in 30.15 ms
  95 % in 37.62 ms
  99 % in 58.47 ms

Status code distribution:
  [OK]   100000 responses
```

### 5.1.2 对 unary grpc 进行压测，全量采集

### 5.1.2 对 unary grpc 进行压测，概率采集(0.5)
