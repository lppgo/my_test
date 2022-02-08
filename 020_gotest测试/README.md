[toc]

# 1: go 测试原则：

## 1.2 单元测试：目的在于发现程序设计或实现的逻辑错误！

1：文件名必须以"\*\_test.go"结尾

2：import testing 包

3：测试函数 func TestXxx(t \*testing.T){}

4：函数中通过 test.T 的 Error,Errorf,FailNow,Fatalf 方法，说明测试不通过，调用 Log 方法用来记录测试信息

5: go test -v -run=Test_Division

## 1.3 压力测试：目的在于程序性能测试，让线上程序能在高并发的情况下保持稳定！

1：测试函数 func Benchmark_xxx(t \*testing.B){}

2：go test -test.bench=".\*"

3：go test -test.bench=Benchmark_Division

4：go test -v -test.bench=Benchmark_Division

5：-count 可以指定执行多少次 go test -test.bench=Benchmark_division -count 10000

1:go test -run=文件名字 -bench=bench 名字 -cpuprofile=生产的 cprofile 文件名称 文件夹

2:
-test.cpuprofile cpu.out : 是否输出 cpu 性能分析文件

-test.memprofile mem.out : 是否输出内存性能分析文件

-test.blockprofile block.out : 是否输出内部 goroutine 阻塞的性能分析文件

go test -test.bench=Benchmark_Division -count 100 -cpuprofile=cpu.profile ./
生成 cpu.profile 和 xxx.test 文件

3：使用 go tool pprof

web ui 可视化：go tool pprof --web 24gotest 测试.test cpu.profile
