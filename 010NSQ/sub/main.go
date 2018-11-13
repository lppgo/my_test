package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/nsqio/go-nsq"
)

var (
	//nsqd的地址，使用tcp监听端口
	tcpNsqAddr = "127.0.0.1:4150"
)

//结构体实现HandleMessage接口方法
type NsqHandler struct {
	msgCount     int64  //消息数目
	nsqHandlerID string //标识ID
}

//实现HandleMessage方法
func (s *NsqHandler) HandleMessage(message *nsq.Message) error {
	//每次收到一条消息+1
	s.msgCount++
	//打印输出信息和ID
	fmt.Println(s.msgCount, s.nsqHandlerID)
	//打印消息的一些基本信息
	fmt.Printf("msg.TimesTamp=%v,msg.nsqaddr=%s,msg.body=%v\n", time.Unix(0, message.Timestamp).Format("2006-01-02 15:04:05"), message.NSQDAddress, string(message.Body))
	return nil
}

func main() {
	//初始化配置
	config := nsq.NewConfig()
	//创建消费者
	consumer, err := nsq.NewConsumer("Insert", "channel1", config)
	if err != nil {
		fmt.Println("NewConsumer failed :" + err.Error())
	}
	//添加处理回调
	handler := NsqHandler{
		nsqHandlerID: "1",
	}
	consumer.AddHandler(&handler)
	//连接对应的nsqd
	err = consumer.ConnectToNSQD(tcpNsqAddr)
	if err != nil {
		fmt.Println("ConnectToNSQD failed : " + err.Error())
	}

	//只是为了不结束此进程，这里没有意义
	wg := &sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
