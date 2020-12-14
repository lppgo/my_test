[toc]

# 1 micro dockerfile

```dockerfile
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

# 2 常用 dockerfile

```dockerfile
FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY service/hello/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/hello service/hello/hello.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/hello /app/hello
COPY --from=builder /app/etc /app/etc

CMD ["./hello", "-f", "etc/hello-api.yaml"]
```
