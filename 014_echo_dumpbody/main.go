package main

import (
    "context"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "net/http"
    "os"
    "os/signal"
    "strings"
    "time"
)

func main() {
    e := echo.New()
    e.Use(middleware.Recover())
    e.Use(middleware.RequestID())
    e.Use(middleware.Secure())
    e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
        Skipper: GzipDefaultSkipper,
        Level:   5,
    }))

    e.Use(middleware.BodyDumpWithConfig(DefaultBodyDumpConfig))

    api := e.Group("/api", Filter)
    {
        api.GET("/test", Test)
        api.POST("/files/upload/multi", UploadFile)
        api.POST("/img/code", ImageCode)
    }
    port := ":8080"
    go func() {
        if err := e.Start(port); err != nil {
            e.Logger.Info("shutting down the server")
        }
    }()
    // Wait for interrupt signal to gracefully shutdown the server with
    // a timeout of 10 seconds.
    quit := make(chan os.Signal)
    signal.Notify(quit, os.Interrupt)
    <-quit
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    if err := e.Shutdown(ctx); err != nil {
        e.Logger.Fatal(err)
    }
}

func Filter(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        // 路由拦截 - 登录身份、资源权限判断等
        println("Api路由拦截：", c.Path())
        return next(c)
    }
}

func Test(c echo.Context) error {
    return c.JSON(http.StatusOK, map[string]interface{}{
        "code": "200",
        "msg":  "请求成功",
        "data": "test",
    })
}

func ImageCode(c echo.Context) error {
    return c.JSON(http.StatusOK, map[string]interface{}{
        "code": "200",
        "msg":  "加载成功",
        "data": "123456",
    })
}

func UploadFile(c echo.Context) error {
    return c.JSON(http.StatusOK, map[string]interface{}{
        "code": "200",
        "msg":  "上传成功",
        "data": nil,
    })
}

// gzip 排除动态资源，如：图形验证码
func GzipDefaultSkipper(c echo.Context) bool {
    if c.Path() == "/api/img/code" {
        return true
    }
    return false
}

var DefaultBodyDumpConfig = middleware.BodyDumpConfig{
    Skipper: BodyDumpDefaultSkipper,
    Handler: func(c echo.Context, reqBody []byte, resBody []byte) {
        println("API请求结果拦截：", string(resBody))
        // 1、解析返回的json数据，判断接口执行成功或失败。如： {"code":"200","data":"test","msg":"请求成功"}
        // 2、保存操作日志
    },
}

// 排除文件，如果您的请求/响应有效负载非常大，例如文件上载/下载，需要进行排查。否则将影响响应时间
func BodyDumpDefaultSkipper(c echo.Context) bool {
    if strings.Contains(c.Path(), "/api/files/") {
        return true
    }
    return false
}
