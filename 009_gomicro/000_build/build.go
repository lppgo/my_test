package build

// 用来安装一些二进制包，go install

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
	// protoc
	// protoc-gen-go
)

func readme() {
	// 1: 将需要安装二进制的包路径引入
	// 2：go mod tidy
	// 3: 执行go install 具体引入的某个包路径
}
