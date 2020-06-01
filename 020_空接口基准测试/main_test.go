package main

import "testing"

//----------------------------------------测试interface{}---解析性能--------------------------------------------------
type InterfaceA interface{ AA() }
type InterfaceB interface{ BB() }

type A struct{ v int }
type B struct{ v int }

func (a *A) AA() {
	a.v += 1
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

func NormalSwitch(a *A) { a.AA() }

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

// ------------------------------------------------------------基准测试参数说明----------------------------------------------------------
//go test -bench=. -benchtime=5s -benchmem -run=none
/*
=============请求参数:
-bench=. ：表示的是运行所有的基准测试，. 表示全部。

-benchtime=5s:表示的是运行时间为5s，默认的时间是1s。

-benchmem:表示显示memory的指标。

-run=none:表示过滤掉单元测试，不去跑UT的cases。

=============响应参数:
BenchmarkJoinStrUseNor-8 执行的函数名称以及对应的GOMAXPROCS值。
79888155     b.N的值
15.5 ns/op   执行一次函数所花费的时间
0 B/op       执行一次函数分配的内存
0 allocs/op  执行一次函数所分配的内存次数
*/
// ----------------------------------------------------------------------------------------------------------------------
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
//----------------------------------------测试传[]User还是[]*User---解析性能--------------------------------------------------
type User struct {
	ID     int
	Name   string
	Age    int
	Weight float32
	Length float32
}

//直接返回[]User
func TransSlice() []User { return make([]User, 10000) }

//返回 []*User
func TransSliceAddr() []*User { return make([]*User, 10000) }

func RangeSlice() {
	for _, _ = range TransSlice() {
	}
}

func RangeSliceAddr() {
	for _, _ = range TransSliceAddr() {
	}
}

func Benchmark_RangeSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RangeSlice()
	}
}

func Benchmark_RangeSliceAddr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RangeSliceAddr()
	}
}

/*
$ go test -bench=. -benchtime=5s -benchmem -run=none
goos: windows
goarch: amd64
pkg: main_test
Benchmark_RangeSlice-8             97002             62700 ns/op          401408 B/op          1 allocs/op
Benchmark_RangeSliceAddr-8        348837             16696 ns/op           81920 B/op          1 allocs/op
PASS
ok      main_test       12.731s


//
*/
