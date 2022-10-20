/*
@File    :   main.go
@Time    :   2022/03/03 14:59:20
@Author  :   lpp
@Version :   1.0.0
@Contact :   golpp@qq.com
@Desc    :   pulsar-go 消费者示例
*/
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/pkg/errors"
)

func main() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:                     "pulsar://172.17.0.1:6650",
		OperationTimeout:        30 * time.Second,
		ConnectionTimeout:       30 * time.Second,
		MaxConnectionsPerBroker: 10,
	})

	if err != nil {
		err = errors.Wrap(err, "new pulsar client error")
		log.Fatal(err)
	}
	defer client.Close()

	// metrics
	// go PulsarConsumerMetrics()

	// consumer
	topics := []string{
		"my-topic-1",
		"my-topic-2",
		"my-topic-3",
	}
	fmt.Println("consumer1 start ....")
	log.Println("consumer1 start ....")
	consumer1(client, topics)
	// consumer2(client, topics)

	log.Println("consumer1 end ....")
	fmt.Println("consumer1 end ....")

}

// Start a separate goroutine for Prometheus metrics
// In this case, Prometheus metrics can be accessed via http://localhost:12112/metrics
// func PulsarConsumerMetrics() {
// 	prometheusPort := 32113
// 	log.Printf("Starting Prometheus metrics at http://localhost:%v/metrics\n", prometheusPort)
// 	http.Handle("/metrics", promhttp.Handler())
// 	err := http.ListenAndServe(":"+strconv.Itoa(prometheusPort), nil)
// 	if err != nil {
// 		log.Printf("pulsar-consumer prometheus listen error :%s", err.Error())
// 	}
// }

func consumer1(client pulsar.Client, topics []string) {
	log.Println("consumer1 ...")
	hostname, _ := os.Hostname()
	msgCh := make(chan pulsar.ConsumerMessage, 5000)
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topics:           topics, // consumer订阅单topic可以保证消息的顺序，但订阅多个topic不能保证消息的顺序
		Name:             "pulsar-consumer",
		SubscriptionName: "pulsar-consumer-subscribe-1",
		Properties: map[string]string{
			"host":    hostname,
			"ip":      "127.0.0.1",
			"srvName": "pulsar-producer",
			"time":    time.Now().Local().Format("2006-01-02 15:04:05"),
		},
		Type: pulsar.Exclusive,
		DLQ: &pulsar.DLQPolicy{
			MaxDeliveries:    3,
			DeadLetterTopic:  "deadletter-topic",
			RetryLetterTopic: "retryletter-topic",
		},
		MessageChannel:    msgCh,
		ReceiverQueueSize: 5000,
	})
	if err != nil {
		err = errors.Wrap(err, "pulsar client subscribe error")
		log.Fatal(err)
	}
	defer consumer.Close()

	// consumer receive timeout
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	// defer cancel()

	var receivedCounter int64
	for {
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			err = errors.Wrap(err, "pulsar consumer2 receive error")
			log.Fatal(err)
		}
		receivedCounter++
		consumer.Ack(msg)
		fmt.Printf("Received[1] message ID :%v ; Content :%s\n", msg.ID(), msg.Payload())
		log.Printf("receivedCounter :%d \n", receivedCounter)
	}

}

// func consumer2(client pulsar.Client, topics []string) {
// 	hostname, _ := os.Hostname()
// 	msgCh := make(chan pulsar.ConsumerMessage, 5000)
// 	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
// 		Topics:           topics, // consumer订阅单topic可以保证消息的顺序，但订阅多个topic不能保证消息的顺序
// 		Name:             "pulsar-consumer-2",
// 		SubscriptionName: "pulsar-consumer-subscribe-2",
// 		Properties: map[string]string{
// 			"host":    hostname,
// 			"ip":      "127.0.0.1",
// 			"srvName": "pulsar-producer-2",
// 			"time":    time.Now().Local().Format("2006-01-02 15:04:05"),
// 		},
// 		Type: pulsar.Exclusive,
// 		DLQ: &pulsar.DLQPolicy{
// 			MaxDeliveries:    3,
// 			DeadLetterTopic:  "deadletter-topic",
// 			RetryLetterTopic: "retryletter-topic",
// 		},
// 		MessageChannel:    msgCh,
// 		ReceiverQueueSize: 5000,
// 	})
// 	if err != nil {
// 		err = errors.Wrap(err, "pulsar client subscribe error")
// 		log.Fatal(err)
// 	}
// 	defer consumer.Close()

// 	for cm := range msgCh {
// 		msg := cm.Message
// 		consumer.Ack(cm.Message)
// 		fmt.Printf("Received[2] message ID :%v ; Content :%s\n", msg.ID(), msg.Payload())
// 	}
// }

/*
(Reader)
Pulsar Reader 是消息处理程序，与 Pulsar 消费者非常相似，但有两个重要区别:

 1: you can specify where on a topic readers begin processing messages (consumers always begin with the latest available unacked message);
 2: readers 不会保留数据或确认消息。
*/

func consumer_reader() {
	// https://github.com/apache/pulsar-client-go
}
