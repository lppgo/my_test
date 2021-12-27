package main

import (
	"mylog/zaplogger"
	"time"
)

func init() {
	logpath := "./log/"
	logfileName := "log.log"
	srvName := "srvA"
	zaplogger.LogConf(srvName, logpath, logfileName)
}

func main() {
	tick := time.NewTicker(time.Second * 3)
	for {
		select {
		case <-tick.C:
			zaplogger.Logger.Debug("logger one debug")
			zaplogger.Logger.Info("logger two info")
			zaplogger.Logger.Error("logger three error")
		default:
			// Logger.Info("main_info", "logger default ")

		}
	}
}
