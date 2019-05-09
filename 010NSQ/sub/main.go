package main

import (
	"encoding/json"
	"fmt"

	"github.com/nsqio/go-nsq"
)

var (
	//nsqd的地址，使用tcp监听端口
	nsqLookupds = "127.0.0.1:4150"
	topic       = "Insert" //主题
	channel     = "my_nsq_channel_1"
	gocount     = 10 //该参数指示为消息处理生成的goroutine的数量
)

//消费消息结构体
type ConsumeMessage struct {
	Name  string `json:"name"`
	Index int    `json:"index"`
	//参数类型，不同参数对应不同的处理方法
	Type int `json:"type"` //0:偶数，1：奇数
}

const (
	Odd  = 1 //类型是奇数
	Even = 0 //类型是偶数
)

//根据类型进行消息分发处理器
type Handler struct {
}

func main() {
	//初始化配置
	config := nsq.NewConfig()
	config.LookupdPollInterval = time.Second //设置consumer重连时间,默认是60s
	//创建消费者
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		fmt.Println("NewConsumer failed :" + err.Error())
	}
	//消息分发处理
	consumer.AddConcurrentHandlers(&Handler{}, gocount)
	// consumer.AddHandler(&Handler{})
	//连接对应的nsqd
	//ConnectToNSQLookupds将多个nsqlookupd地址添加到此comsumer实例的列表中。
	//如果添加第一个地址，它将启动一个HTTP请求来发现配置主题的nsqd生成器。
	err = consumer.ConnectToNSQD(nsqLookupds)
	if err != nil {
		fmt.Println("ConnectToNSQD failed : " + err.Error())
		panic(err)
	}
	fmt.Println("end ... ")
	select {}
}

// 实现HandleMessage方法,进行消息分发处理
func (s *Handler) HandleMessage(message *nsq.Message) error {
	msg := ConsumeMessage{}
	err := json.Unmarshal(message.Body, &msg)
	if err != nil {
		fmt.Println("解码message发生错误：" + err.Error())
		return err
	}
	if 0 == msg.Type {
		OddErr := HandleOddMessage(msg)
		if OddErr != nil {
			fmt.Println("OddErr :" + OddErr.Error())
			return OddErr
		}
	} else if 1 == msg.Type {
		EvenErr := HandleEvenMessage(msg)
		if EvenErr != nil {
			fmt.Println("EvenErr :" + EvenErr.Error())
			return EvenErr
		}
	}
	message.Finish()
	return nil
}

var oddCount, evenCount int

//奇数处理
func HandleOddMessage(msg ConsumeMessage) error {
	oddCount++
	fmt.Printf("message 是奇数，index是%d\n", oddCount)
	fmt.Printf("%+v\n", msg)
	return nil
}

//偶数处理
func HandleEvenMessage(msg ConsumeMessage) error {
	oddCount++
	fmt.Printf("message 是偶数，index是%d\n", oddCount)
	fmt.Printf("%+v\n", msg)
	return nil
}
