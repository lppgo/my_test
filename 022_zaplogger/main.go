package main

import (
	"mylog/zaplogger"

	"github.com/gin-gonic/gin"
)

func init() {
	logpath := "./log/"
	logfileName := "log.log"
	zaplogger.LogConf(logpath, logfileName)
}

func test(c *gin.Context) {
	zaplogger.Logger.Info("aaaaaaaa") //用于测试
}

func main() {
	router := gin.Default()
	router.POST("/test", test)
	router.Run(":8888")
}
