*************** Golang中优秀的消息队列NSQ基础安装及使用 ******************
//https://blog.csdn.net/sd653159/article/details/83624661
1 ： 背景介绍
2 ： nsq 服务介绍

    （1）nsqlookupd : 主要负责服务发现，负责nsqd的心跳,状态监测 。给客户端,nsqadmin提供nsqd地址与状态
    
    ====启动命令：nsqlookupd
    
    （2）nsqd : 负责接收消息，存储队列和将消息发送给客户端。nsqd可以多机器部署，当你使用客户端像一个topic发送消息的时候，可以分配多个nsqd地址，
    消息会随机的分配到各个nsqd上，nsqd优先把消息存储到内存channel中，当内存channel满了之后，则把消息写入磁盘文件。nsqd监听了2个tcp端口，
    一个用来服务客户端，一个用来提供http的接口
    
    ====启动命令：nsqd --lookupd-tcp-address=127.0.0.1:4160
    
    （3）nsqadmin : 是一个web管理页面
    
    ====启动命令：nsqadmin --lookupd-http-address=127.0.0.1:4161
    
    启动之后，可以通过http://127.0.0.1:4171可以访问这个管理页面

3 ： nsq使用的是pub/sub模式 (发布/订阅，生产者/消费者)。我们可以先发布一个主题到nsq,然后所有订阅的服务器就会异步的从这里读取主题的内容

    Topic：发布的主题名字

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
  github.com/nsqio/go-nsq


5 :案例笔记
    
    (1)：安装nsq
    
        nsqlookupd
        
        nsqd --lookupd-tcp-address=127.0.0.1:4160
        
        nsqadmin --lookupd-http-address=127.0.0.1:4161
    
        通过4171端口，在web UI查看
        
    (2):写代码 publish(发布)，subscribe(订阅)       
-----------------------------------------------------------------------------------
介绍： https://segmentfault.com/a/1190000009194607

NSQ是由四个重要组件构成：   

    nsqd：一个负责接收、排队、转发消息到客户端的守护进程。   
    
    nsqlookupd：管理拓扑信息，并提供最终一致性的发现服务的守护进程。   
    
    nsqadmin：一套Web用户界面，可实时查看集群的统计数据和执行各种各样的管理任务。 
    
    utilities：常见基础功能、数据流处理工具，如nsq_stat、nsq_tail、nsq_to_file、nsq_to_http、nsq_to_nsq、to_nsq。


为了达到高效的分布式消息服务，NSQ实现了合理、智能的权衡，从而使得其能够完全适用于生产环境中，具体内容如下： 支持消息内存队列的大小设置，默认完全持久化（值为0），消息既可以可持久到磁盘，也可以保存在内存中 保证消息至少传递一次，以确保消息可以最终成功发送 收到的消息是无序的， 实现了松散订购 发现服务nsqlookupd具有最终一致性，消费者最终能够找到所有Topic生产者

分布式的实时消息平台NSQ

从上图可以看出，单个nsqd可以有多个Topic，每个Topic又可以有多个Channel。Channel能够接收Topic所有消息的副本，从而实现了多播分发；而Channel上的每个消息被分发给它的一个订阅者，从而实现负载均衡。

