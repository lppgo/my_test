# Hello Service

This is the Hello service

Generated with

```
micro new --namespace=go.micro --type=service hello
```

```go
// protoc 编译
cd foobar
protoc --proto_path=.:$GOPATH/src --go_out=. --micro_out=. proto/foobar/foobar.proto
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.service.hello
- Type: service
- Alias: hello

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./hello-service
```

Build a docker image
```
make docker
```