/*
@File    :   main.go
@Time    :   2022/03/03 13:46:51
@Author  :   lpp
@Version :   1.0.0
@Contact :   golpp@qq.com
@Desc    :   pulsar 生产者示例
*/
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	// go PulsarProducerMetrics()

	// producer
	hostname, _ := os.Hostname()
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "my-topic-1",
		// Topic: "my-topic-2",
		// Topic: "my-topic-3",
		Name: "pulsar-producer-1",
		// Name: "pulsar-producer-2",
		// Name: "pulsar-producer-3",
		Properties: map[string]string{
			"host":    hostname,
			"ip":      "127.0.0.1",
			"srvName": "pulsar-producer-1",
			// "srvName": "pulsar-producer-2",
			// "srvName": "pulsar-producer-3",
			"time": time.Now().Local().Format("2006-01-02 15:04:05"),
		},
		SendTimeout: time.Second * 5,
		// Schema: &pulsar.JSONSchema{},
		DisableBatching: true,
	})
	if err != nil {
		err = errors.Wrap(err, "new pulsar producer error")
		log.Fatal(err)
	}
	defer producer.Close()

	// write your business logic here

	TickServerProducer(producer)

}

// Start a separate goroutine for Prometheus metrics
// In this case, Prometheus metrics can be accessed via http://localhost:12112/metrics
func PulsarProducerMetrics() {
	prometheusPort := 12112
	log.Printf("Starting Prometheus metrics at http://localhost:%v/metrics\n", prometheusPort)
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":"+strconv.Itoa(prometheusPort), nil)
	if err != nil {
		log.Printf("pulsar-producer prometheus listen error :%s", err.Error())
	}
}

func HTTPServerProducer(producer pulsar.Producer) {
	ctx := context.Background()
	webPort := 8082
	http.HandleFunc("/produce", func(w http.ResponseWriter, r *http.Request) {
		msgId, err := producer.Send(ctx, &pulsar.ProducerMessage{
			Payload: []byte(fmt.Sprintf("hello world")),
		})
		if err != nil {
			err = errors.Wrap(err, "pulsar-producer send error :")
			log.Fatal(err)
		} else {
			log.Printf("Published message: %v", msgId)
			fmt.Fprintf(w, "Published message: %v", msgId)
		}

		//
		if err = producer.Flush(); err != nil {
			log.Printf("flush message [msgID = %v ] error: %s", msgId, err.Error())
		}

	})

	err := http.ListenAndServe(":"+strconv.Itoa(webPort), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func TickServerProducer(producer pulsar.Producer) {
	log.Println("tick server producer start ...")
	ctx := context.Background()
	ticker := time.NewTicker(time.Millisecond * 100)

	var sendCounter int64
	for {
		select {
		case t := <-ticker.C:
			tStr := t.Format("2006-01-02 15:04:05.000")
			topic := producer.Topic()
			data := []byte(tStr + " " + topic)
			sendData(ctx, producer, data)
			sendCounter++
			log.Printf("sendCounter : %d\n", sendCounter)
		default:
		}
	}
}

func sendData(ctx context.Context, producer pulsar.Producer, data []byte) {
	msgId, err := producer.Send(ctx, &pulsar.ProducerMessage{
		Payload: data,
		// Value: interface{},
	})
	if err != nil {
		err = errors.Wrap(err, "pulsar-producer send error :")
		log.Fatal(err)
	} else {
		log.Printf("Published message: %v,tipic:%s\n", msgId, producer.Topic())
	}

	//
	if err = producer.Flush(); err != nil {
		log.Printf("flush message [msgID = %v ] error: %s", msgId, err.Error())
	}
}
