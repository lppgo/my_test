*************** Golang中优秀的消息队列NSQ基础安装及使用 ******************
//https://blog.csdn.net/sd653159/article/details/83624661


1 ： 背景介绍

2 ： nsq 服务介绍

    （1）nsqlookupd : 主要负责服务发现，负责nsqd的心跳,状态监测 。给客户端,nsqadmin提供nsqd地址与状态
    ====启动命令：nsqlookupd

    （2）nsqd : 负责接收消息，存储队列和将消息发送给客户端。nsqd可以多机器部署，当你使用客户端像一个topic发送消息的时候，
    可以分配多个nsqd地址，消息会随机的分配到各个nsqd上，nsqd优先把消息存储到内存channel中，当内存channel满了之后，
    则把消息写入磁盘文件。nsqd监听了2个tcp端口，一个用来服务客户端，一个用来提供http的接口
    ====启动命令：nsqd --lookupd-tcp-address=127.0.0.1:4160
    
    （3）nsqadmin : 是一个web管理页面
    ====启动命令：nsqadmin --lookupd-http-address=127.0.0.1:4161
    启动之后，可以通过http://127.0.0.1:4171可以访问这个管理页面

3 ： nsq使用的是pub/sub模式 (发布/订阅，生产者/消费者)。我们可以先发布一个主题到nsq,然后所有订阅的服务器就会异步的从这里读取主题的内容

    Topic(左上角)：发布的主题名字

    NSQd Host：Nsq主机服务地址

    Channel：消息通道

    Depth：消息积压量

    In-flight：已经投递但是还未消费掉的消息

    Deferred：没有消费掉的延时消息

    Messages：服务器启动之后，总共接收到的消息量

    Connections：通道里面客户端的订阅数

    TimeOut：超时时间内没有被响应的消息数

    Memory + Disk：储存在内存和硬盘中总共的消息数


4: 在代码中发布主题内容，然后通过订阅Topic去异步读取消息
    go get github.com/nsqio/go-nsq


5: (1)先创建一个主题，并发布100条消息
