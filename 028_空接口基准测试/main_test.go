package main

import "testing"

type InterfaceA interface {
	AA()
}

type InterfaceB interface {
	BB()
}

type A struct {
	v int
}

func (a *A) AA() {
	a.v += 1
}

type B struct {
	v int
}

func (b *B) BB() {
	b.v += 1
}

func TypeSwitch(v interface{}) {
	switch v.(type) {
	case InterfaceA:
		v.(InterfaceA).AA()
	case InterfaceB:
		v.(InterfaceB).BB()
	}
}

func NormalSwitch(a *A) {
	a.AA()
}

func InterfaceSwitch(v interface{}) {
	v.(InterfaceA).AA()
}

func Benchmark_TypeSwitch(b *testing.B) {
	var a = new(A)

	for i := 0; i < b.N; i++ {
		TypeSwitch(a)
	}
}

func Benchmark_NormalSwitch(b *testing.B) {
	var a = new(A)

	for i := 0; i < b.N; i++ {
		NormalSwitch(a)
	}
}

func Benchmark_InterfaceSwitch(b *testing.B) {
	var a = new(A)

	for i := 0; i < b.N; i++ {
		InterfaceSwitch(a)
	}
}

//go test -bench=. -benchtime=5s -benchmem -run=none
/*
    -bench=. ：表示的是运行所有的基准测试，. 表示全部。

    -benchtime=5s:表示的是运行时间为5s，默认的时间是1s。

    -benchmem:表示显示memory的指标。

    -run=none:表示过滤掉单元测试，不去跑UT的cases。
*/

// 测试结果
/*
$ go test -bench=. -benchtime=10s -benchmem -run=none
goos: windows
goarch: amd64
pkg: main_test
Benchmark_TypeSwitch-8          773618374               16.0 ns/op             0 B/op          0 allocs/op
Benchmark_NormalSwitch-8        1000000000               1.27 ns/op            0 B/op          0 allocs/op
Benchmark_InterfaceSwitch-8     1000000000               7.88 ns/op            0 B/op          0 allocs/op
PASS
ok      main_test       24.068s

*/
