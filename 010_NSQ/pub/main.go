package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-nsq"
)

var (
	producer *nsq.Producer
	//nsqd的地址，使用了tcp监听的port
	tcpNsqdAddr = "127.0.0.1:4150"
	topic       = "Insert"
)

//发送消息结构体
type ProduceMessage struct {
	Name  string `json:"name"`
	Index int    `json:"index"`
	//参数类型，不同参数对应不同的处理方法
	Type int `json:"type"` //0:偶数，1：奇数
}

//producer
func main() {
	//初始化配置
	config := initConfig()
	//创建Producer
	var createProducerErr error
	producer, createProducerErr = nsq.NewProducer(tcpNsqdAddr, config)
	if createProducerErr != nil {
		fmt.Println("New producer failed :" + createProducerErr.Error())
	}
	for i := 0; i < 100; i++ {
		//主题内容
		tContent := &ProduceMessage{
			Name:  "lucas_nsq_是一个实时消息队列",
			Index: i,
			Type:  i % 2, //取余数
		}
		//将数据编码
		body, err := json.Marshal(tContent)
		if err != nil {
			fmt.Println("json_marshal_err:" + err.Error())
		}
		//发布消息
		err = producer.Publish(topic, body)
		if err != nil {
			fmt.Println("Publish failed :" + err.Error())
		}
	}
	log.Println("All producer publish end !")
	select {}
}

// initConfig 初始化配置
func initConfig() *nsq.Config {
	config := nsq.NewConfig()
	return config
}
