# go-micro/v2  最新最全最详细入门教程
[toc]

## 一 : 环境配置
```go
// 需要的二进制插件
1: <protoc> :protobuf编译器
            https://github.com/protocolbuffers/protobuf/releases
2: <protoc-gen-go> :protobuf运行时 
            https://github.com/protocolbuffers/protobuf-go
            // cd protobuf-go/cmd/protoc-gen-go
            // 在当前目录go install

3: <protoc-gen-grpc-gateway> :
            github.com/grpc-ecosystem/grpc-gateway
            // git clone 
            // go install
4: <protoc-gen-swagger> :
            github.com/grpc-ecosystem/grpc-gateway
5: <micro> :
            go get -u -v github.com/micro/micro/v2
6: <protoc-gen-micro> :
            go get -u -v  github.com/micro/micro/v3/cmd/protoc-gen-micro@v2.9.3

=======
replace github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.14.1
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
```
### 1: 安装grpc-go (https://github.com/grpc/grpc-go)
```go
go get -u -v google.golang.org/grpc
```
### 2: 安装protobuf(https://github.com/protocolbuffers/protobuf)

#### 2.1:安装Protocol Compiler (protoc)
```go
(1):https://github.com/protocolbuffers/protobuf
(2):然后解压到本地.本地需要已经安装好 (apt-get install autoconf automake libtool)
```

#### 2.2:安装Protobuf Runtime (protoc-gen-go)
```go
https://github.com/protocolbuffers/protobuf
```
### 3: 安装micro
### 4: 安装服务注册中心
```go
// consul
// etcd 使用docker-compose方式
// etcd 客户端工具:
https://www.electronjs.org/apps/etcd-manager
```


## 二 : 安装go-micro
```go
// go-micro (安装go-micro框架)
go get -u -v github.com/micro/go-micro/v2
// micro（工具）;go 命令安装不了，用https://github.com/micro/micro/releases安装
go get -u -v github.com/micro/micro/v2
```
## 三 : 编写第一个服务
### 1：编写接口定义文件proto


### 2：编译接口文件
```go
// 生成pb.go文件
// grpc-gateway
// micro.go文件
 protoc -I /usr/local/include \
 -I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
 --plugin=protoc-gen-micro=$GOPATH/bin/protoc-gen-micro \
 --proto_path=. \
 --go_out=.\
 --micro_out=.\
 proto/*.proto
```
### 3：实现service-server
### 4：实现service-client
### 5：测试
## 四 ：使用其他插件，比如web
## 五 ：安装go-micro 
## 六 ：其他
```go
// 1: micro和go-micro区别
go-micro:  微服务开发库

micro:   基于go-micro开发的运行时工具集

// 2：micro的工具集组件
    A、API：将http请求转向内部应用
　　　　1、API：将http请求映射到API接口
　　　　2、RPC：将http请求映射到RPC服务
　　　　3、event：将http请求广播到订阅者
　　　　4、proxy：反向代理
　　　　5、web：支持websocket反向代理
　　B、Web：web反向代理与管理控制
　　C、Proxy：代理风格的请求，支持异构系统只需要瘦客户端便可调用Micro服务
　　　　1. 注意：与Micro API不同，Proxy只处理micro风格的RPC请求，而非http请求
　　D、Cli：以命令行操控Micro服务
　　E、Bot：与常见的通信软件对接，负责传递消息，远程指令操作
```
