[TOC]

```go
1: Slice
2: copy()函数
3: int和[]byte互转
4: go Generate UUID
5: IsEmpty()判断给的值是否为空
6: 二分法对Slice进行插入排序
7: RSA加密解密
8：切片反转
9：对Get,Post请求的封装，防止http句柄泄露
```

# 二： 一些整理的题目

### 1: Slice

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

### 2: copy()函数

```go
func myCopy() {
	dst := []int{1, 2, 3}
	src := []int{4, 5, 6, 7, 8}
	n := copy(dst, src)
	fmt.Printf("dst: %v, n: %d", dst, n)
}
```

### 3: interface{} ，鸭子类型，简单工厂模式

### 5: 工厂模式

### 6: 值接收者和指针接收者

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

### 9：对 Get,Post 请求的封装，防止 http 句柄泄露

```go
// curl 发起 get请求
func CurlGet(uri string, timeout time.Duration) (result []byte, err error) {
	cli := &http.Client{}
	// 写入 uri 请求信息
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	// 设置超时
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req = req.WithContext(ctx)
	// 发起请求
    resp, err := cli.Do(req)
    // 关闭连接
	if err != nil {
        err = errors.WithStack(err)
		return
	}
	defer resp.Body.Close()
	// 读取 body
	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}
// curl 支持POST form表单形式
func CurlFormPOST(uri, token string, params map[string]interface{}, timeout time.Duration) (result []byte, err error) {
	cli := &http.Client{}
	values := url.Values{}
	for k, v := range params {
		if v != nil {
			values.Set(k, cast.ToString(v))
		}
	}
	// 写入 post 请求数据
	req, err := http.NewRequest(http.MethodPost, uri, strings.NewReader(values.Encode()))
	if err != nil {
		return
	}
	// 设置超时
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req = req.WithContext(ctx)
	// 设置 header
	req.Header.Set("ACCESS-TOKEN", token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    resp, err := cli.Do(req)
	if err != nil {
		return
    }
    // 必须关闭
	defer resp.Body.Close()
	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}
// curl 支持POST json
func CurlJsonPOST(uri, token string, params map[string]interface{}, timeout time.Duration) (result []byte, err error) {
	cli := &http.Client{}
	// 数据打包
	data, err := json.Marshal(params)
	if err != nil {
		return
	}
	// 写入 post 请求数据
	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	// 设置超时
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req = req.WithContext(ctx)
	// 设置 header
	req.Header.Set("ACCESS-TOKEN", token)
	req.Header.Set("Content-Type", "application/json")
	// 发起 http 请求
    resp, err := cli.Do(req)
	if err != nil {
		return
    }
    defer resp.Body.Close()
	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}
```
