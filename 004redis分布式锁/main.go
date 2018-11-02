package main

import (
	"fmt"
	"log"
	"time"

	"git.yeeuu.com/yeeuu/my_test/004redis分布式锁/db"
	"git.yeeuu.com/yeeuu/my_test/004redis分布式锁/redislock"
	"github.com/robfig/cron"
)

var i int

func main() {
	log.Println("Is running ...")
	c := cron.New()
	spec := "0-59/3 * * * * *"
	c.AddFunc(spec, cronWork)
	c.Start()
	for {

	}
}

func cronWork() {
	log.Println("cron job running ...")
	// go func() {
	// 	for {

	// fmt.Println(fmt.Sprintf("%v"), conn)

	rdlock := redislock.RedisLock{
		LockKey: "sunstrider_hotel_redislock",
		Value:   "exists",
	}
	conn := db.GetConn("127.0.0.1:6379", "")
	defer conn.Close()

	// err := rdlock.Lock(&conn, 24*59*59)
	err := rdlock.Lock(&conn, 5)
	i++
	fmt.Println("err:", i, err)
	if err == nil {
		log.Println("Do Something ... ")
		// fmt.Println(str)
	}
	// rdlock.Unlock(&conn) //key过期自动解锁
	time.Sleep(3 * time.Second)
	// 	}
	// }()
}
