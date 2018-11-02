package main

import (
	"log"
	"time"

	"github.com/robfig/cron"
)

func main() {
	i := 0
	c := cron.New()
	spec := "0-59/3 * * * * *"
	// 秒，分，时，每个月的天，月，星期
	c.AddFunc(spec, func() {
		i++
		log.Println("cron runing : ", i)
	})
	c.Start()
	time.Sleep(3 * time.Minute)
}
