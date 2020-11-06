[toc]

# go 单元测试与测试覆盖度检测

## 1：单元测试 Unit Test

```go
// 普通单元测试
```

## 2：web 开发 handler 测试器

### 2.1 ：net/http/httptest

#### 2.1.1：注意

**注意**
(1) 假设在 server 中 handler 已经写好，按下面方法进行测试
(2) 如果 Web Server 有操作数据库的行为，需要在 init 函数中进行数据库的连接。
(3) 其他前置参数初始化
(4) http.NewRequest 替换为 httptest.NewRequest
(5) httptest.NewRequest 的第三个参数可以用来传递 body 数据，必须实现 io.Reader 接口
(6) 解析响应时没直接使用 ResponseRecorder，而是调用了 Result 函数

#### 2.1.2: GET handler 测试示例

```go
import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHealthCheckHandler(t *testing.T) {
    //创建一个请求
    req, err := http.NewRequest("GET", "/health-check", nil)
    if err != nil {
        t.Fatal(err)
    }

    // 我们创建一个 ResponseRecorder (which satisfies http.ResponseWriter)来记录响应
    rr := httptest.NewRecorder()

    //直接使用HealthCheckHandler，传入参数rr,req
    HealthCheckHandler(rr, req)

    // 检测返回的状态码
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // 检测返回的数据
    expected := `{"alive": true}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

```

#### 2.1.3: POST handler 测试示例

```go
func TestHealthCheckHandler2(t *testing.T) {
    reqData := struct {
        Info string `json:"info"`
    }{Info: "P123451"}

    reqBody, _ := json.Marshal(reqData)
    // fmt.Println("input:", string(reqBody))
    req := httptest.NewRequest(http.MethodPost,"/health-check",bytes.NewReader(reqBody))

    req.Header.Set("userid", "wdt")
    req.Header.Set("commpay", "brk")

    rr := httptest.NewRecorder()
    HealthCheckHandler(rr, req)

    result := rr.Result()

    body, _ := ioutil.ReadAll(result.Body)
    fmt.Println(string(body))

    if result.StatusCode != http.StatusOK {
        t.Errorf("expected status 200,",result.StatusCode)
    }
}
```

#### 2.1.4: 结合 context 使用

```go
func TestGetProjectsHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/api/users", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    // e.g. func GetUsersHandler(ctx context.Context, w http.ResponseWriter, r *http.Request)
    handler := http.HandlerFunc(GetUsersHandler)

    // Populate the request's context with our test data.
    ctx := req.Context()
    ctx = context.WithValue(ctx, "app.auth.token", "abc123")

    // Add our context to the request: note that WithContext returns a copy of
    // the request, which we must assign.
    req = req.WithContext(ctx)
    handler.ServeHTTP(rr, req)

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }
}
```

### 2.2 ：https://github.com/gavv/httpexpect

## 3：测试覆盖度检测

```go
// 1: 编写单元测试

// 2：生成单元测试可执行文件xxx.test;生成测试覆盖率profile文件coverage.data
go test -c -covermode=count -coverpkg ./ -coverprofile=coverage.data ./

// 3: 执行单元测试可执行文件
./xxx.test
// 4: 将测试覆盖率profile文件输出为html文件，可在网页中直接查看各文件的测试覆盖率情况
go tool cover -html=coverage.data -o coverage.html
```
