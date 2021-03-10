package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
)

// 定义全局限流器对象
var rateLimiter ratelimit.Limiter

// gin 限流中间件---漏桶法
func leakyBucket() gin.HandlerFunc {
	prev := time.Now()
	return func(c *gin.Context) {
		takeNow := rateLimiter.Take()
		fmt.Println("===>", takeNow.Sub(prev)) //打印每次调用的时间间隔
		prev = takeNow
	}
}

func main() {
	rateLimiter = ratelimit.New(10)
	router := gin.Default()
	router.GET("/ping", leakyBucket(), func(c *gin.Context) {
		c.JSON(200, "ok")
	})
	router.Run("0.0.0.0:6868")
}

//  使用ab 压测
//  ab -n 200 -c 2 http://172.29.213.233:6868/ping