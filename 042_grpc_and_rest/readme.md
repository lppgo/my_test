[toc]

# 1: summary

本项目是一个 grpc+rest 的服务 demo
源自:https://mp.weixin.qq.com/s/6FZVr_gInuQLUJREDNzrKQ

# 2: 项目实现

## 2.1 写 proto,生成 pb,grpc.pb

```go
protoc \
--go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
./order.proto


```

```go

protoc-gen-go 和 protoc-gen-go-grpc 这两个插件有什么不同？

当使用参数 --go_out=plugins=grpc:xxx 生成时，生成的文件 *.pb.go 包含消息序列化代码和 gRPC 代码。

当使用参数 --go_out=xxx --go-grpc_out=xxx 生成时，会生成两个文件 *.pb.go 和 *._grpc.pb.go ，它们分别是消息序列化代码和 gRPC 代码;

```

## 2.2 初始化 grpc 服务

## 2.3 初始化 restful 服务

## 2.4 结合 grpc 和 REST 服务

## 2.5 添加错误处理 errCh

## 2.6 refactor 代码，初始化 appserver
