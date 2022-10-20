[toc]

# 1: 介绍

<李平平 proto 项目模板>，一键安装依赖，生成各种 proto 文件。

- 该项目是公共的`xxx.proto`接口文件
- 以及生成的`*.pb.go`文件
  - `xxx.pb.go`
  - `xxx_grpc.pb.go`
  - `xxx_micro.pb.go`

# 2: 脚本

- `go_run.sh`是生成 pb.go 的脚本
- `Makefile`脚本

# 3: 包管理设置

## 3.1 gitsubmodule

```go
gitsubmodule
```

## 3.2 gomod

```go
(1) 修改hosts
(2)
go env -w GOINSECURE="gitlab.xxx.cn"
go env -w GONOSUMDB="gitlab.xxx.cn"
go env -w GONOPROXY="gitlab.xxx.cn"
go env -w GOPRIVATE="gitlab.xxx.cn"

// (3) go get
go get com.xxx.xxx.proto


// (4) git tag --> go get
go get com.xxx.xxx.proto@v1.0.0
```

# 4: other

```go
--go-micro_out=debug=true,,components="micro|http",paths=source_relative:./go_proto
--openapi_out==paths=source_relative:./go_proto
--swagger_out=logtostderr=true:./go_proto
```
