- [三：一些整理的好用的库](#三一些整理的好用的库)
    - [1: anti ,一个高性能，低损耗的goroutine池](#1-anti-一个高性能低损耗的goroutine池)
    - [2: gpool对象复用池](#2-gpool对象复用池)
    - [3: gcache是一个高速的单进程缓存模块,gcache提供了并发安全的缓存控制接口](#3-gcache是一个高速的单进程缓存模块gcache提供了并发安全的缓存控制接口)
    - [4: gcache LRU缓存淘汰策略  Least Recently Used](#4-gcache-lru缓存淘汰策略--least-recently-used)
    - [5: goroutine池 grpool](#5-goroutine池-grpool)
    - [6: errgroup](#6-errgroup)
    - [7:](#7)
    - [8:](#8)
    - [9:](#9)
缓存模块gcache提供了并发安全的缓存控制接口)
    * [4: gcache LRU缓存淘汰策略  Least Recently Used](#4-gcache-lru缓存淘汰策略--least-recently-used)     
    * [5: goroutine池 grpool](#5-goroutine池-grpool)
    * [6: errgroup](#6-errgroup)
    * [7:](#7)
    * [8:](#8)
    * [9:](#9)

# 三：一些整理的好用的库

### 1: anti ,一个高性能，低损耗的goroutine池

```go
https://github.com/panjf2000/ants
```

### 2: gpool对象复用池

> ```go
> // github.com/gogf/gf/container/gpool
> func f1gpool() {
> 	//创建一个对象复用池
> 	//对象过期时间为3000毫秒
> 	//给定创建和销毁方法
> 	pool := gpool.New(3000,
> 		func() (interface{}, error) {
> 			return gtcp.NewConn("www.baidu.com:80")
> 		},
> 		func(i interface{}) {
> 			glog.Println("expire")
> 			i.(*gtcp.Conn).Close()
> 		})
> 
> 	//从对象复用池获取一个对象
> 	// Pick obj  from pool
> 	conn, err := pool.Get()
> 	if err != nil {
> 		glog.Error("pool.Get Err:", err)
> 	}
> 	result, err := conn.(*gtcp.Conn).SendRecv([]byte("HEAD / HTTP/1.1\n\n"), -1)
> 	if err != nil {
> 		glog.Error("conn SendRecv Err:", err)
> 	}
> 
> 	glog.Println(string(result))
> 
> 	//丢回池中，以重复利用
> 	pool.Put(conn)
> 
> 	//
> 	time.Sleep(time.Second * 5)
> }
> ```

### 3: gcache是一个高速的单进程缓存模块,gcache提供了并发安全的缓存控制接口

```go
// github.com/gogf/gf/os/gcache
func f2gcache() {
	//创建一个缓存对象
	cache := gcache.New()
	//不过期
	cache.Set("k1", "v1", 0)
	cache.Set("k2", "v2", 0)
	glog.Println(cache.Values())
	glog.Printf("根据key获取对应val:%v\n", cache.Get("k1"))

	glog.Printf("获取缓存大小:%d\n", cache.Size())
	glog.Printf("判断缓存中是否有指定的key:%v\n", cache.Contains("k1"))
	glog.Printf("删除指定key:%v的缓存\n", cache.Remove("k1"))
	glog.Println(cache.Get("k1") == nil)

	// 关闭缓存对象，让GC回收
	cache.Close()
}
```

### 4: gcache LRU缓存淘汰策略  Least Recently Used

```go
URL 缓存淘汰策略
1.新添加的数据放在头部 
2.被访问到的数据放在头部
3.超过最大缓存量的数据将被移除。
func f3LRU() {
	// 设置LRU淘汰数量
	c := gcache.New(3)

	// 添加10个元素，不过期
	for i := 0; i < 10; i++ {
		c.Set(i, i, 0)
	}
	glog.Println(c.Size())
	glog.Println(c.Keys())

	// 读取键名1，保证该键名是优先保留
	glog.Println(c.Get(1))

	// 等待一定时间后(默认1秒检查一次)，
	// 元素会被按照从旧到新的顺序进行淘汰
	time.Sleep(2 * time.Second)
	glog.Println(c.Size())
	glog.Println(c.Keys())
}
```

### 5: goroutine池 grpool

```go
// github.com/gogf/gf/os/grpool
func f4grpool() {
	pool := grpool.New(100)
	for i := 0; i < 1000; i++ {
		pool.Add(job)
	}
	glog.Println("worker:", pool.Size())
	glog.Println("  jobs:", pool.Jobs())
	//设定时间间隔
	gtimer.SetInterval(time.Second, func() {
		glog.Println("worker:", pool.Size())
		glog.Println("  jobs:", pool.Jobs())
		glog.Println()
	})
	select {}
}
```

### 6: errgroup

### 7:

### 8:

### 9:
