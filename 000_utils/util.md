

[TOC]

# 一： 一些整理的函数

### 1: string和[]byte高效互转
```go
// string高效转换为[]byte
func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// []byte高效转换为string
func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
```
### 2: 结构体和[]byte互转
```go
type MyStruct struct {
	A int
	B int
}

var sizeOfMyStruct = int(unsafe.Sizeof(MyStruct{}))

func MyStructToBytes(s *MyStruct) []byte {
	var x reflect.SliceHeader
	x.Len = sizeOfMyStruct
	x.Cap = sizeOfMyStruct
	x.Data = uintptr(unsafe.Pointer(s))
	return *(*[]byte)(unsafe.Pointer(&x))
}

func BytesToMyStruct(b []byte) *MyStruct {
	return (*MyStruct)(unsafe.Pointer(
		(*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
	))
}
```
### 3: int和[]byte互转
```go
func Int2Bytes(n int) []byte {
	data := int64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}
func Bytes2Int(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data)
}
```
### 4: go Generate UUID
```go
//"github.com/google/uuid"
func GenerateUUID() {
	id := uuid.New()
	fmt.Printf("%s %s\n", id, id.Version().String())
	id2, _ := uuid.NewRandom()
	fmt.Printf("%s %s\n", id2, id2.Version().String())
}
```
### 5: IsEmpty()判断给的值是否为空
```go
func IsEmpty(value interface{}) bool {
	if value == nil {
		return true
	}
	switch value := value.(type) {
	case int:
		return value == 0
	case int8:
		return value == 0
	case int16:
		return value == 0
	case int32:
		return value == 0
	case int64:
		return value == 0
	case uint:
		return value == 0
	case uint8:
		return value == 0
	case uint16:
		return value == 0
	case uint32:
		return value == 0
	case uint64:
		return value == 0
	case float32:
		return value == 0
	case float64:
		return value == 0
	case bool:
		return value == false
	case string:
		return value == ""
	case []byte:
		return len(value) == 0
	default:
		// Finally using reflect.
		rv := reflect.ValueOf(value)
		switch rv.Kind() {
		case reflect.Chan,
			reflect.Map,
			reflect.Slice,
			reflect.Array:
			return rv.Len() == 0

		case reflect.Func,
			reflect.Ptr,
			reflect.Interface,
			reflect.UnsafePointer:
			if rv.IsNil() {
				return true
			}
		}
	}
	return false
}
```
### 6: 二分法对Slice进行插入排序
```go
type muxEntry struct {
	h       http.Handler
	pattern string
}

func appendSorted(es []muxEntry, e muxEntry) []muxEntry {
	n := len(es)
	// 排序之后，进行二分搜索，若找见符合条件=true,返回index，如果不符合条件,i==n
	i := sort.Search(n, func(i int) bool {
		return len(es[i].pattern) < len(e.pattern) //sorted from longtest to shortest
		//return len(es[i].pattern) > len(e.pattern)
	})
	if i == n {
		return append(es, e)
	}
	// we now know that i points at where we want to insert
	es = append(es, muxEntry{}) // try to grow the slice in place, any entry works.
	copy(es[i+1:], es[i:])      // Move shorter entries down
	es[i] = e
	return es
}
```

### 7:RSA加密解密

```go
//加密
func RAS_Encrypt() {
	//生成密钥
	var err error
	privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey = privateKey.PublicKey
	//加密
	encryptedBytes, err = rsa.EncryptOAEP(sha256.New(), rand.Reader, &publicKey, msg, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("encrypted bytes: ", encryptedBytes)
}
// 解密
func RAS_Decrypt() {
	var err error
	decryptedBytes, err = privateKey.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		panic(err)
	}
	fmt.Println("decrypted message:", string(decryptedBytes))
}
```



# 二： 一些整理的题目

### 1:  Slice

```go
// 数组切片的知识点
// 1: 基本结构
// 2: Slice扩容
func mySliceArray() {
	nums := [3]int{}
	nums[0] = 1

	fmt.Printf("nums: %v , len: %d, cap: %d\n", nums, len(nums), cap(nums)) //
	dnums := nums[0:2]
	dnums[0] = 5
	fmt.Printf("nums: %v ,len: %d, cap: %d\n", nums, len(nums), cap(nums))
	fmt.Printf("dnums: %v, len: %d, cap: %d\n", dnums, len(dnums), cap(dnums))
	//fmt.Println(drums[2])
}
```

### 2:  copy()函数

```go
func myCopy() {
	dst := []int{1, 2, 3}
	src := []int{4, 5, 6, 7, 8}
	n := copy(dst, src)
	fmt.Printf("dst: %v, n: %d", dst, n)
}
```

### 3:  interface{} ，鸭子类型，简单工厂模式

### 5:  工厂模式

### 6:  值接收者和指针接收者

### 7：数组是值类型,切片是引用类型

### 8：切片反转

```go
func reverse() {
	s := []int{0, 1, 2, 3, 4, 5}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(s)
}
```

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

# 四：其他

