/**
 * @Author: Tiffa
 * @Description: etcd 常用操作
 * @File:  main_etcd.go
 * @Version: 1.0.0
 * @Date: 2020/6/3 16:23
 */
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

//type Client struct {
//	Cluster //向集群里增加etcd服务端节点之类，属于管理员操作
//	KV		//我们主要使用的功能，即K-V键值库的操作
//	Lease   //租约相关操作，比如申请一个TTL=10秒的租约（应用给key可以实现键值的自动过期）
//	Watcher //观察订阅，从而监听最新的数据变化
//	Auth    //管理etcd的用户和权限，属于管理员操作
//	Maintenance  //维护etcd,比如主动迁移etcd的leader节点，属于管理员操作
//	conn *grpc.ClientConn  //
//	cfg           Config   //
//	creds         grpccredentials.TransportCredentials
//	resolverGroup *endpoint.ResolverGroup
//	mu            *sync.RWMutex
//	ctx    context.Context
//	cancel context.CancelFunc
//	// Username is a user name for authentication.
//	Username string
//	// Password is a password for authentication.
//	Password        string
//	authTokenBundle credentials.Bundle
//	callOpts []grpc.CallOption
//	lg *zap.Logger
//}

func main() {
	fmt.Println("----------------------------1: 连接客户端---------------------------------")
	conf := clientv3.Config{
		Endpoints:        []string{"localhost:2379"}, //etcd的多个节点服务地址
		AutoSyncInterval: 0,
		DialTimeout:      5 * time.Second, //创建client的首次连接超时时间，这里传了5秒
	}
	cli, err := clientv3.New(conf)
	if err != nil {
		panic("连接失败:" + err.Error())
	}
	defer cli.Close()

	//2：我们通过方法clientv3.NewKV()来获得KV接口的实现（实现中内置了错误重试机制）：
	kv := clientv3.NewKV(cli)

	//3：我们通过kv操作etcd中的数据
	fmt.Println("-------------------------------------- Put --------------------------------------")
	putResp, err := kv.Put(context.TODO(), "/test/key1", "hello etcd !")
	if err != nil {
		fmt.Println("Put操作错误:", err.Error())
	}
	fmt.Printf("Put操作成功，putResp:%+v\n", putResp)
	_, _ = kv.Put(context.TODO(), "/test/key2", "Hello World!")
	_, _ = kv.Put(context.TODO(), "/testspam", "spam")

	fmt.Println("-------------------------------------- Get --------------------------------------")
	getResp, err := kv.Get(context.TODO(), "/test/key2")
	if err != nil {
		fmt.Println("Get操作错误:", err.Error())
	}
	fmt.Println(getResp.Kvs)
	//fmt.Printf("Get操作成功，getResp:%+v", getResp)
	getResp, err = kv.Get(context.TODO(), "/test/", clientv3.WithPrefix()) //获取前缀1
	if err != nil {
		fmt.Println("Get操作(WithPrefix)错误:", err.Error())
	}
	//fmt.Println(getResp.Kvs)
	getResp, err = kv.Get(context.TODO(), "/test", clientv3.WithPrefix()) //获取前缀2
	if err != nil {
		fmt.Println("Get操作(WithPrefix)错误:", err.Error())
	}

	fmt.Println("-------------------------------------- Lease ------------------------------------")
	//Grant：分配一个租约。
	//Revoke：释放一个租约。
	//TimeToLive：获取剩余TTL时间。
	//Leases：列举所有etcd中的租约。
	//KeepAlive：自动定时的续约某个租约。
	//KeepAliveOnce：为某个租约续约一次。
	//Close：释放当前客户端建立的所有租约。

	//当我们实现服务注册时，需要主动给Lease进行续约，通常是以小于TTL的间隔循环调用Lease的KeepAliveOnce()方法对租约进行续期，
	//一旦某个服务节点出错无法完成租约的续期，等key过期后客户端即无法在查询服务时获得对应节点的服务，这样就通过租约到期实现了服务的错误隔离。
	lease := clientv3.NewLease(cli)
	grantResp, err := lease.Grant(context.TODO(), 10)
	if err != nil {
		fmt.Println("lease.Grant错误:", err.Error())
	}
	_, _ = kv.Put(context.TODO(), "/test/vanish", "vanish in 10 s", clientv3.WithLease(grantResp.ID))

	//ch, err := lease.KeepAlive(context.TODO(), grantResp.ID)
	//_, _ = lease.KeepAliveOnce(context.TODO(), grantResp.ID)

	fmt.Println("-------------------------------------- Op ---------------------------------------")
	ops := []clientv3.Op{
		clientv3.OpPut("/put-key", "123"),
		clientv3.OpGet("/get-key"),
		clientv3.OpPut("/put-key", "456"),
	}
	for _, op := range ops {
		if _, err := cli.Do(context.TODO(), op); err != nil {
			fmt.Println("cli.Do操作错误:", err.Error())
		}
	}
	fmt.Println("-------------------------------------- Txn事务 ----------------------------------")
	//etcd中的事务是原子执行的，只支持if ... then ... else ... 这种表达式
	//If(满足条件) Then(执行若干Op) Else(执行若干Op)
	kv.Put(context.TODO(), "/k1", "6")
	gr, _ := kv.Get(context.TODO(), "/k1")
	for key, val := range gr.Kvs {
		fmt.Printf("index:%d  Value:%s  Version:%d\n", key, val.Value, val.Version)
	}

	txn := kv.Txn(context.TODO())
	txn.If(
		clientv3.Compare(clientv3.Version("/k1"), ">", 1),
	).Then(
		clientv3.OpPut("/k2", "100"),
	).Else(
		clientv3.OpPut("/k3", "99"),
	).Commit()

	fmt.Println("-------------------------------------- Watch -----------------------------------")
	//Watch 用于监听某个值的变化
	//Watch的典型应用场景是应用于系统配置的热加载,我们可以在系统读取到存储在etcd key中的配置后，用Watch监听key的变化
	appConfig := &AppConfig{}
	confKey := "/config_key"
	confObj, _ := json.Marshal(appConfig)
	kv.Put(context.TODO(), confKey, string(confObj))
	watchConfig(cli, "/config_key", &appConfig)
	select {}
	fmt.Println("-------------------------------------- etcd  执行完毕 ... -----------------------")
}

type AppConfig struct {
	Addr string
	Port string
	Name string
}

func watchConfig(cli *clientv3.Client, key string, obj interface{}) {
	watchCh := cli.Watch(context.TODO(), key)
	go func() {
		for res := range watchCh {
			value := res.Events[0].Kv.Value
			if err := json.Unmarshal(value, obj); err != nil {
				fmt.Println("now", time.Now(), "watchConfig err", err)
				continue
			}
			fmt.Println("now", time.Now(), "watchConfig", obj)
		}
	}()

}
