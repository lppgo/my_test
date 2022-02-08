[toc]

# 1: httptest

- 在 Web 开发场景下的单元测试，如果涉及到 HTTP 请求推荐大家使用 Go 标准库 net/http/httptest 进行测试，能够显著提高测试效率。

在这一小节，我们以常见的 gin 框架为例，演示如何为 http server 编写单元测试。

假设我们的业务逻辑是搭建一个 http server 端，对外提供 HTTP 服务。我们编写了一个 helloHandler 函数，用来处理用户请求。

## 1.1 httptest_demo/gin.go

```go
// gin.go
package httptest_demo

import (
 "fmt"
 "net/http"

 "github.com/gin-gonic/gin"
)

// Param 请求参数
type Param struct {
 Name string `json:"name"`
}

// helloHandler /hello请求处理函数
func helloHandler(c *gin.Context) {
 var p Param
 if err := c.ShouldBindJSON(&p); err != nil {
  c.JSON(http.StatusOK, gin.H{
   "msg": "we need a name",
  })
  return
 }
 c.JSON(http.StatusOK, gin.H{
  "msg": fmt.Sprintf("hello %s", p.Name),
 })
}

// SetupRouter 路由
func SetupRouter() *gin.Engine {
 router := gin.Default()
 router.POST("/hello", helloHandler)
 return router
}
```

## 1.2 httptest_demo/gin_test.go

- 现在我们需要为 helloHandler 函数编写单元测试，这种情况下我们就可以使用 httptest 这个工具 mock 一个 HTTP 请求和响应记录器，让我们的 server 端接收并处理我们 mock 的 HTTP 请求，同时使用响应记录器来记录 server 端返回的响应内容。

- 单元测试的示例代码如下：

```go
// gin_test.go
package httptest_demo

import (
 "encoding/json"
 "net/http"
 "net/http/httptest"
 "strings"
 "testing"

 "github.com/stretchr/testify/assert"
)

func Test_helloHandler(t *testing.T) {
 // 定义两个测试用例
 tests := []struct {
  name   string
  param  string
  expect string
 }{
  {"base case", `{"name": "liwenzhou"}`, "hello liwenzhou"},
  {"bad case", "", "we need a name"},
 }

 r := SetupRouter()

 for _, tt := range tests {
  t.Run(tt.name, func(t *testing.T) {
   // mock一个HTTP请求
   req := httptest.NewRequest(
    "POST",                      // 请求方法
    "/hello",                    // 请求URL
    strings.NewReader(tt.param), // 请求参数
   )

   // mock一个响应记录器
   w := httptest.NewRecorder()

   // 让server端处理mock请求并记录返回的响应内容
   r.ServeHTTP(w, req)

   // 校验状态码是否符合预期
   assert.Equal(t, http.StatusOK, w.Code)

   // 解析并检验响应内容是否复合预期
   var resp map[string]string
   err := json.Unmarshal([]byte(w.Body.String()), &resp)
   assert.Nil(t, err)
   assert.Equal(t, tt.expect, resp["msg"])
  })
 }
}
```

## 1.3 执行单元测试查看结果

```bash
Running tool: D:\SoftBox\Go\bin\go.exe test -timeout 30s -coverprofile=C:\Users\ADMINI~1\AppData\Local\Temp\vscode-go9CwX48\go-code-cover -run ^Test_helloHandler$ test-1/httptest_demo -v

=== RUN   Test_helloHandler
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /hello                    --> test-1/httptest_demo.helloHandler (3 handlers)
=== RUN   Test_helloHandler/base_case
[GIN] 2021/12/31 - 10:00:36 | 200 |         550µs |       192.0.2.1 | POST     "/hello"
=== RUN   Test_helloHandler/bad_case
[GIN] 2021/12/31 - 10:00:36 | 200 |            0s |       192.0.2.1 | POST     "/hello"
--- PASS: Test_helloHandler (0.04s)
    --- PASS: Test_helloHandler/base_case (0.04s)
    --- PASS: Test_helloHandler/bad_case (0.00s)
PASS
coverage: 100.0% of statements
ok  	test-1/httptest_demo	0.096s	coverage: 100.0% of statements
```

# 2: gock

- 上面的示例介绍了如何在 HTTP Server 服务类场景下为请求处理函数编写单元测试，那么如果我们是在代码中请求外部 API 的场景（比如通过 API 调用其他服务获取返回值）又该怎么编写单元测试呢？

例如，我们有以下业务逻辑代码，依赖外部 API：http://your-api.com/post提供的数据。

```go
// api.go

// ReqParam API请求参数
type ReqParam struct {
 X int `json:"x"`
}

// Result API返回结果
type Result struct {
 Value int `json:"value"`
}

func GetResultByAPI(x, y int) int {
 p := &ReqParam{X: x}
 b, _ := json.Marshal(p)

 // 调用其他服务的API
 resp, err := http.Post(
  "http://your-api.com/post",
  "application/json",
  bytes.NewBuffer(b),
 )
 if err != nil {
  return -1
 }
 body, _ := ioutil.ReadAll(resp.Body)
 var ret Result
 if err := json.Unmarshal(body, &ret); err != nil {
  return -1
 }
 // 这里是对API返回的数据做一些逻辑处理
 return ret.Value + y
}
```

- 在对类似上述这类业务代码编写单元测试的时候，如果不想在测试过程中真正去发送请求或者依赖的外部接口还没有开发完成时，我们可以在单元测试中对依赖的 API 进行 mock。

## 2.1 安装

- 这里推荐使用 gock 这个库。 `go get -u gopkg.in/h2non/gock.v1`

## 2.2 使用示例

- 使用 gock 对外部 API 进行 mock，即 mock 指定参数返回约定好的响应内容。下面的代码中 mock 了两组数据，组成了两个测试用例。

```go
// api_test.go
package gock_demo

import (
 "testing"

 "github.com/stretchr/testify/assert"
 "gopkg.in/h2non/gock.v1"
)

func TestGetResultByAPI(t *testing.T) {
 defer gock.Off() // 测试执行后刷新挂起的mock

 // mock 请求外部api时传参x=1返回100
 gock.New("http://your-api.com").
  Post("/post").
  MatchType("json").
  JSON(map[string]int{"x": 1}).
  Reply(200).
  JSON(map[string]int{"value": 100})

 // 调用我们的业务函数
 res := GetResultByAPI(1, 1)
 // 校验返回结果是否符合预期
 assert.Equal(t, res, 101)

 // mock 请求外部api时传参x=2返回200
 gock.New("http://your-api.com").
  Post("/post").
  MatchType("json").
  JSON(map[string]int{"x": 2}).
  Reply(200).
  JSON(map[string]int{"value": 200})

 // 调用我们的业务函数
 res = GetResultByAPI(2, 2)
 // 校验返回结果是否符合预期
 assert.Equal(t, res, 202)

 assert.True(t, gock.IsDone()) // 断言mock被触发
}
```

## 2.3 执行 gock 测试用例

```go
$ go test -v
=== RUN   TestGetResultByAPI
--- PASS: TestGetResultByAPI (0.00s)
PASS
ok      golang-unit-test-demo/gock_demo 0.054s
```

# 3: summary 总结
