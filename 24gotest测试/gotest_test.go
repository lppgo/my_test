package gotest

import (
	"testing"
)

//单元测试，debug 的时候注意打好断点，F11
func Test_Division(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("Division除法测试不通过！")
	} else {
		// 记录一些日志
		t.Log("Dvicesion 测试通过！")
		t.Log("i:", i)
	}

}

//性能测试
func Benchmark_Division(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}

func Benchmark_TimeConsuming(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	//...做一些初始化的工作,例如读取文件数据,数据库连接之类的,这样这些时间不影响我们测试函数本身的性能

	b.StartTimer() //重新开始时间计数
	// 或者 重置定时器
	// b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}
