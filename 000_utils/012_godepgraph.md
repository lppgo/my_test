[TOC]

# 1: godepgraph go 依赖可视化

## 1.1: install godepgraph

```go
go get github.com/kisielk/godepgraph
```

## 1.2: install graphviz

[graphviz 安装](http://graphviz.org/download/)

## 1.3: usage

```go
godepgraph github.com/kisielk/godepgraph | dot -Tpng -o godepgraph.png

// 忽略标准库软件包，请使用"-s"标志
godepgraph -s github.com/kisielk/godepgraph
// 忽略 vendor "-novendor"
// 根据包名忽略 "-i"
godepgraph -i github.com/foo/bar,github.com/baz/blah github.com/something/else
// 根据prefix忽略 "-p"
godepgraph -p github.com,launchpad.net bitbucket.org/foo/bar

```

## 1.4: color 说明

**green**: a package that is part of the Go standard library, installed in $GOROOT.

**blue**: a regular Go package found in $GOPATH.

**yellow**: a vendored Go package found in $GOPATH.

**orange**: a package found in $GOPATH that uses cgo by importing the special package "C".

# 2: godepgraph 来 deal with 包的循环引用
