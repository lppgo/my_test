[toc]

[Go Mock (gomock)简明教程](https://geektutu.com/post/quick-gomock.html)

https://geektutu.com/post/quick-gomock.html

## 1: gomock 测试简介

上一篇文章 Go Test 单元测试简明教程(https://geektutu.com/post/quick-go-test.html) 介绍了 Go 语言中
单元测试的常用方法，包括子测试(subtests)、表格驱动测试(table-driven tests)、帮助函数(helpers)、网络测试
和基准测试(Benchmark)等。这篇文章介绍一种新的测试方法，mock/stub 测试，当待测试的函数/对象的依赖关系很复杂，
并且有些依赖不能直接创建，例如数据库连接、文件 I/O 等。这种场景就非常适合使用 mock/stub 测试。简单来说，就是用
mock 对象模拟依赖项的行为。

gomock 是官方提供的 mock 框架，同时还提供了 mockgen 工具用来辅助生成测试代码。

```go
    go get -u github.com/golang/mock/gomock
    go get -u github.com/golang/mock/mockgen
```

## 2: 写一个demo
第一步：使用 mockgen 生成 db_mock.go。
般传递三个参数。包含需要被mock的接口得到源文件source，生成的目标文件destination，包名package。
```go
mockgen -source=db.go -destination=db_mock.go -package=main
``` 
第二步：新建 db_test.go，写测试用例
```go
func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用

	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))

	if v := GetFromDB(m, "Tom"); v != -1 {
		t.Fatal("expected -1, but got", v)
	}
}
```
第三步：执行测试
```go
    go test . -cover -v 
```
## 3: stub打桩
3.1 参数(Eq, Any, Not, Nil)
```go
m.EXPECT().Get(gomock.Eq("Tom")).Return(0, errors.New("not exist"))
m.EXPECT().Get(gomock.Any()).Return(630, nil)
m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil) 
m.EXPECT().Get(gomock.Nil()).Return(0, errors.New("nil")) 
Eq(value) 表示与 value 等价的值。

Any() 可以用来表示任意的入参。
Not(value) 用来表示非 value 以外的值。
Nil() 表示 None 值
```
3.2 返回值(Return, DoAndReturn)
```go
m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil)
m.EXPECT().Get(gomock.Any()).Do(func(key string) {
    t.Log(key)
})
m.EXPECT().Get(gomock.Any()).DoAndReturn(func(key string) (int, error) {
    if key == "Sam" {
        return 630, nil
    }
    return 0, errors.New("not exist")
})

Return 返回确定的值
Do Mock 方法被调用时，要执行的操作吗，忽略返回值。
DoAndReturn 可以动态地控制返回值。
```

3.3 调用次数(Times)
```go
func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil).Times(2)
	GetFromDB(m, "ABC")
	GetFromDB(m, "DEF")
}

Times() 断言 Mock 方法被调用的次数。
MaxTimes() 最大次数。
MinTimes() 最小次数。
AnyTimes() 任意次数（包括 0 次）。
```

3.4 调用顺序(InOrder)
```go
func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用

	m := NewMockDB(ctrl)
	o1 := m.EXPECT().Get(gomock.Eq("Tom")).Return(0, errors.New("not exist"))
	o2 := m.EXPECT().Get(gomock.Eq("Sam")).Return(630, nil)
	gomock.InOrder(o1, o2)
	GetFromDB(m, "Tom")
	GetFromDB(m, "Sam")
}

```
