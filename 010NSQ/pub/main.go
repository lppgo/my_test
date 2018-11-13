package main

import (
	"fmt"
	"strconv"

	"github.com/nsqio/go-nsq"
)

var (
	//nsqd的地址，使用了tcp监听的port
	tcpNsqdAddr = "127.0.0.1:4150"
)

//producer
func main() {
	//initConfig
	config := nsq.NewConfig()
	for i := 0; i < 10; i++ {
		//创建100个生产者Producer
		tPro, err := nsq.NewProducer(tcpNsqdAddr, config)
		if err != nil {
			fmt.Println("New producer failed :" + err.Error())
		}
		//主题
		topic := "Insert"
		//主题内容
		tContent := "New data " + strconv.Itoa(i)
		//发布消息
		err = tPro.Publish(topic, []byte(tContent))
		if err != nil {
			fmt.Println("Publish failed :" + err.Error())
		}
	}
}
