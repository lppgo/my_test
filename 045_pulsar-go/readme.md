[toc]

# 1：介绍

- 本项目是 golang 使用 pulsar 的 demo;
- pulsar 介绍:https://mp.weixin.qq.com/s/wYb9UQWYJf8ay1DbLXaZDQ

# 2: 安装好 Pulsar 和 PulsarManager

- pulsar
- pulsar-manager
- pulsarctl

# 3: pulsar go

- 您可以使用 Pulsar Go 客户端 来创建使用 Go 语言的 Pulsar 生产者（producer） 、 消费者（consumer） 和 readers ;
- https://pulsar.apache.org/docs/zh-CN/next/client-libraries-go/

## 3.1 安装 go 工具包

```go
//
go get -u "github.com/apache/pulsar-client-go/pulsar"
import "github.com/apache/pulsar-client-go/pulsar"
```

## 3.2 连接 URL

- 要使用 client 连接到 Pulsar，你需要指定 Pulsar 协议的 url;
- Pulsar 协议 URL 使用 pulsar scheme 来指定被连接的集群:`pulsar://localhost:6650`;
- 如果你有多个 broker，你可以使用下面的方法设置 URl:`pulsar://localhost:6550,localhost:6651,localhost:6652`;
- 如果你使用 TLS 认证，那么 URL 应该是这样的:`pulsar+ssl://pulsar.us-west.example.com:6651`;

## 3.2 创建客户端

```go
import (
    "log"
    "time"
    "github.com/apache/pulsar-client-go/pulsar"
)

func main() {
    client, err := pulsar.NewClient(pulsar.ClientOptions{
        // URL: "pulsar://localhost:6650,localhost:6651,localhost:6652",
        URL: "pulsar://172.17.0.1:6650",
        OperationTimeout:  30 * time.Second, // 操作超时
        ConnectionTimeout: 30 * time.Second, // 建立TCP连接超时
    })
    if err != nil {
        log.Fatalf("Could not instantiate Pulsar client: %v", err)
    }

    defer client.Close()
}
```

## 3.3 Producers

- Pulsar producers publish messages to Pulsar topics. You can configure Go producers using a ProducerOptions object. Here's an example:

```go
producer, err := client.CreateProducer(pulsar.ProducerOptions{
    Topic: "my-topic",
})

if err != nil {
    log.Fatal(err)
}

_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
    Payload: []byte("hello"),
})

defer producer.Close()

if err != nil {
    fmt.Println("Failed to publish message", err)
}
fmt.Println("Published message")
```

- **producer operations** : Pulsar Go producers have the following methods available:
  |Method|说明|Return type|
  |:---|:---|:---|
  |`Topic()`||`string`|
  |`Name()`||`string`|
  |`Send(context.Context, *ProducerMessage)()`||`(MessageID, error)`|
  |`SendAsync(context.Context, *ProducerMessage, func(MessageID, *ProducerMessage, error))()`|||
  |`LastSequenceID()()`||`int64`|
  |`Flush()`||`error`|
  |`Close()`|||

## 3.4 Consumers

- Pulsar consumers subscribe to one or more Pulsar topics and listen for incoming messages produced on that topic/those topics. You can configure Go consumers using a ConsumerOptions object. Here's a basic example that uses channels:
- 消费者操作:
  |Method|说明|Return type|
  |:--|:--|:--|
  |`Subscription()`|Returns the consumer's subscription name|`string`|
  |`Unsubcribe()`||`error`|
  |`Receive(context.Context)`||`(Message, error)`|
  |`Chan()`||`<-chan ConsumerMessage`|
  |`Ack(Message)`| Acknowledges a message to the Pulsar broker||
  |`AckID(MessageID)`|Acknowledges a message to the Pulsar broker by message ID||
  |`ReconsumeLater(msg Message, delay time.Duration)`|||
  |`Nack(Message)`|||
  |`NackID(MessageID)`|||
  |`Seek(msgID MessageID)`|||
  |`SeekByTime(time time.Time)`|||
  |`Close()`||`error`|
  |`Name`||`string`|

## 3.5 Reader

# 4: vegeta

```go
// 压测
echo "GET http://localhost:8082/produce" | vegeta -cpus=8 attack -rate=500 -connections=10 -duration=10s | tee result.bin | vegeta report
```
