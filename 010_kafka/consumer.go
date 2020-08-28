package kafka

import (
	"flag"
	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
	"github.com/gogf/gf/os/glog"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Sarama configuration options
var (
	brokers  = ""
	version  = ""
	group    = ""
	topics   = ""
	assignor = "" //Consumer group partition assignment strategy (range, roundrobin, sticky)
	oldest   = true
	verbose  = false
)

func init() {
	flag.StringVar(&brokers, "brokers", "", "Kafka bootstrap brokers to connect to, as a comma separated list")
	flag.StringVar(&group, "group", "", "Kafka consumer group definition")
	flag.StringVar(&version, "version", "2.1.1", "Kafka cluster version")
	flag.StringVar(&topics, "topics", "", "Kafka topics to be consumed, as a comma separated list")
	flag.StringVar(&assignor, "assignor", "range", "Consumer group partition assignment strategy (range, roundrobin, sticky)")
	flag.BoolVar(&oldest, "oldest", true, "Kafka consumer consume initial offset from oldest")
	flag.BoolVar(&verbose, "verbose", false, "Sarama logging")
	flag.Parse()

	if len(brokers) == 0 {
		panic("no Kafka bootstrap brokers defined, please set the -brokers flag")
	}
	if len(topics) == 0 {
		panic("no topics given to be consumed, please set the -topics flag")
	}
	if len(group) == 0 {
		panic("no Kafka consumer group defined, please set the -group flag")
	}
}

// NewClusterConfig
func NewClusterConfig() *cluster.Config{
	log.Println("New cluster config ...")
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	//config.Version = sarama.V0_10_1_0

	switch assignor {
		case "sticky":
			config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
		case "roundrobin":
			config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
		case "range":
			config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
		default:
			log.Panicf("Unrecognized consumer group partition assignor: %s", assignor)
	}

	if oldest {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}
	return config
}

func ConsumerTopics(){
	brokers:=[]string{}
	topics:=[]string{}
	clusterConfig:=NewClusterConfig()
	consumer, err := cluster.NewConsumer(brokers, group, topics, clusterConfig)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()
	// consume errors
	go ConsumerExceptionHanding(consumer)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	//wg := sync.WaitGroup{}
	for {
		//
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				//wg.Add(1)
				//defer wg.Done()
				switch msg.Topic {
				case "test":
					processTopicLog(msg, consumer)
				case "LOGIN":
					//Login(msg.Value)
					processTopicLog(msg, consumer)
				case "BIND_PHONE":
					//BindPhone(msg.Value)
					processTopicLog(msg, consumer)
				case "YwdOnlineInfo":
					processTopicLog(msg, consumer)
					//EveryOnlineDataProcess(msg)
				case "YwdProduction":
					processTopicLog(msg, consumer)
					//ProductionDataProcess(msg)
				case "YwdTask":
					processTopicLog(msg, consumer, )
					//TaskDataProcess(msg)
				default:
					glog.Errorf("此topic没有消费topic is:%s, topic value is:%s, offset is:%d ", msg.Topic, msg.Value, msg.Offset)
				}
				consumer.MarkOffset(msg, "") // mark message as processed
				//wg.Wait()
			}
		case <-signals:
			return
		}
	}
}


func processTopicLog(msg *sarama.ConsumerMessage, consumer *cluster.Consumer) {
	glog.Infof("-----> Topic:%s <-----, Partition:%d, Offset:%d,\nKey:%s,\nValue:%s", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
}
//ConsumerExceptionHanding 接收消费者返回错误或者通知
func ConsumerExceptionHanding(consumer *cluster.Consumer) {
	for {
		select {
		case err := <-consumer.Errors():
			if err != nil {
				glog.Errorf("consumer.Errors: %s", err.Error())
			}
		case notify := <-consumer.Notifications():
			glog.Infof("consumer.Notifications: %s", notify)
		default:
		}
	}
}
