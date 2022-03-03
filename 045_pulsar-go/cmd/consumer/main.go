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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/pkg/errors"
)

func main() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://172.17.0.1:6650",
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
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
		"my-topic",
		"my-topic-2",
	}
	hostname, _ := os.Hostname()
	msgCh := make(chan pulsar.ConsumerMessage, 10000)
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topics:           topics,
		Name:             "pulsar-consumer",
		SubscriptionName: "pulsar-consumer-subscribe-1",
		Properties: map[string]string{
			"host":    hostname,
			"ip":      "127.0.0.1",
			"srvName": "pulsar-producer",
			"time":    time.Now().Local().Format("2006-01-02 15:04:05"),
		},
		Type:           pulsar.Failover,
		MessageChannel: msgCh,
	})
	if err != nil {
		err = errors.Wrap(err, "pulsar client subscribe error")
		log.Fatal(err)
	}
	defer consumer.Close()

	for cm := range msgCh {
		msg := cm.Message
		fmt.Printf("Received message ID :%v ; Content :%s\n", msg.ID(), msg.Payload())
	}

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
