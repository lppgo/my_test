# 1: 说明

# 2：jeager to client

# 3: jaeger to server

# 4: jaeger 采集设置

# 5：trace 对性能影响压测

## 5.1 unary grpc

### 5.1.1 对 origin unary grpc 进行压测

ghz -insecure \
--proto ./protos/greeter.proto \
--call protos.Greeter.Hello \
-d '{"name":"Lucas"}' \
-c 10 \
-n 100000 \
localhost:50051

### 5.1.2 对 unary grpc 进行压测，全量采集

### 5.1.2 对 unary grpc 进行压测，概率采集(0.5)
