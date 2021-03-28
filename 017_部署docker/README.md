<!-- [toc] -->
- [通过二进制部署docker](#通过二进制部署docker)
- [1. 安装docker 环境。配置docker desktop](#1-安装docker-环境配置docker-desktop)
- [2. 拉取镜像](#2-拉取镜像)
- [3. 进入go项目进行交叉编译](#3-进入go项目进行交叉编译)
- [4. 编写Dockerfile](#4-编写dockerfile)
  - [4.1 直接使用可执行文件](#41-直接使用可执行文件)
  - [4.2  使用源码，在docker进行build](#42--使用源码在docker进行build)
  - [4.3 使用分层构建(1)](#43-使用分层构建1)
  - [4.4 使用分层构建](#44-使用分层构建)
- [5. docker build 构建image](#5-docker-build-构建image)
- [6. docker run 运行容器，进行测试](#6-docker-run-运行容器进行测试)
- [7. docker push  提交image镜像](#7-docker-push--提交image镜像)

# 通过二进制部署docker

# 1. 安装docker 环境。配置docker desktop

> Docker daemon 配置文件

```json5
{
  //加速地址:Preferred Docker registry mirror
  "registry-mirrors": [
    "https://reg-mirror.qiniu.com",
    "https://registry.docker-cn.com",
    "http://hub-mirror.c.163.com",
    "https://3laho3y3.mirror.aliyuncs.com",
    "https://mirror.ccs.tencentyun.com"
  ],
  //Enable insecure registry communication
  "insecure-registries": ["http://harbor.yvjoy.com"],
  "debug": true,
  "experimental": false
}
```

# 2. 拉取镜像

# 3. 进入go项目进行交叉编译

```go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o xxx
```

# 4. 编写Dockerfile

## 4.1 直接使用可执行文件

 ```Dockerfile
 FROM busybox

MAINTAINER  Lucas "lucas@test.com"

# 设置固定的项目路径
WORKDIR /var/mnt/uwd/
ADD ./uwd ./uwd
ADD ./config/config.toml  ./config/config.toml
ADD ./public/ ./public/
RUN chmod 755 ./uwd
EXPOSE 55372
ENTRYPOINT ["./uwd"]
 ```

## 4.2  使用源码，在docker进行build

 ```Dockerfile
 FROM golang:latest

MAINTAINER lucas "golpp@qq.com"

WORKDIR /mnt/docker
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

COPY . .
RUN go mod download 
RUN CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -a -v -o mmsd-server

EXPOSE 8900
ENTRYPOINT ["./mmsd-server"]
 ```

## 4.3 使用分层构建(1)

>> 使用临时镜像

```Dockerfile
FROM golang:1.15.2 as builder_base
ARG A_GOPROXY
ENV GOPROXY=$A_GOPROXY
ENV CGO_ENABLED="0"
ENV GO111MODULE="on"
# RUN apk --no-cache add make git gcc libtool musl-dev
WORKDIR /tmp/sse
COPY go.mod .
COPY go.sum .
RUN go mod download
# RUN go mod tidy
COPY . .

RUN make

FROM alpine:3.12.0
# 设置环境变量
ENV MICRO_NAMESPACE="micro"
ENV MICRO_REGISTRY="etcd"
ENV MICRO_REGISTRY_ADDRESS="127.0.0.1:12379,127.0.0.1:22379,127.0.0.1:32379"
ENV MICRO_ETCD_AUTH_USERNAME="root"
ENV MICRO_ETCD_AUTH_PASSWORD="pwd@BA2020"
ENV MICRO_API_TYPE="web"
ENV MICRO_API_HANDLER="http"
ENV MICRO_API_ENABLE_CORS="true"
ENV MICRO_API_ADDRESS=":9060"
ENV GATEWAY_METRICS_ADDRESS=":9061"
ENV GATEWAY_REGISTER_WEB_ADDRESS=":9062"
# 默认是false，面对聚合网关是true，面对互联时是false
ENV GATEWAY_IS_INTERNAL_MODE="false"
# redis cluster, used to sync rate limit count
ENV GATEWAY_REDIS_ADDRESS="127.0.0.1:6379"
# ENV GATEWAY_REDIS_PASSWORD=""
# ENV GATEWAY_REDIS_DB="0"
# ENV GATEWAY_MAX_MESSAGE_BYTES="4096"
ENV GATEWAY_MIDDLE_PLATFORM_ADDRESS="http://10.112.2.102:10102"
ENV GATEWAY_MIDDLE_PLATFORM_LOGIN_ENDPOINT="/uums/user/login"
ENV GATEWAY_MIDDLE_PLATFORM_TOKEN_CHECK_ENDPOINT="/uums/token/check"
ENV GATEWAY_MIDDLE_PLATFORM_USER_GET_ENDPOINT="/user/get"
ENV GATEWAY_MIDDLE_PLATFORM_CHANGEPW_ENDPOINT="/user/changePassword"

COPY --from=builder_base /tmp/sse/micro /
EXPOSE 9060
# for metrics
EXPOSE 9061
# for web dashboard
EXPOSE 9062

ENTRYPOINT [ "/micro" ] 
CMD [ "service", "api" ]
```

## 4.4 使用分层构建

```Dockerfile
FROM golang:1.15-alpine AS builder
WORKDIR /workspace
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

# src code
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o main ./cmd/server

FROM alpine:3.12.0
COPY --from=builder /workspace/main /main
RUN chmod +x /main
ENV TZ=Asia/Shanghai
ENTRYPOINT ["/main"]
```

```go
// .dockerignore 忽略无用的文件，只 COPY go build 需要的文件。
# Ignore all files which are not go type
!**/*.go
!**/*.mod
!**/*.sum
```

# 5. docker build 构建image

```json
docker build -t xxx:v1.0.0 .
```

# 6. docker run 运行容器，进行测试

```json
docker run --rm --name uwd-c -d -p 55372:55372 xxx:v1.0.0
```

# 7. docker push  提交image镜像

```json5
  docker login xxx.com
  输入用户名/密码
  docker tag  2e25d8496557 xxxxx.com/home/uwd-i:v1.0.0
  docker push xxxxx.com/home/uwd-i:v1.0.0
  
  // docker run -v挂载本地卷，挂载本地配置文件
  docker run --rm -d -p 55372:55372 --name uwdcs -v /e/zoo/socket/uwd/config/config.toml:/var/mnt/uwd/config/config.toml uwd-i:v2.0.0
```
