package main

//1:首先理一下思路，使用代码去注册其他api到go-micro体系中，
// 我们就需要在代码中模拟出这样的json数据，并发送给我们的micro Registry服务

//2:首先我们构造结构体，构建一个这样的结构体，然后把需要的数据填充进去序列化成json
//发送给micro Registry就完成了我们的需求

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"

	"log"

	myhttp "github.com/micro/v2/go-plugins/client/http"
)

func main() {
	etcdReg := etcd.NewRegistry(registry.Addrs("106.12.72.181:23791"))

	mySelector := selector.NewSelector(
		selector.Registry(etcdReg),
		selector.SetStrategy(selector.RoundRobin),
	)
	getClient := myhttp.NewClient(client.Selector(mySelector), client.ContentType("application/json"))

	//1创建request
	req := getClient.NewRequest("api.jtthink.com.test", "/v1/test", map[string]string{}) //这里的request
	//2创建response
	var rsp map[string]interface{}                         //var rsp map[string]string这里这么写也可以，因为我们的返回值是{"data":"test"}，所以都对的上
	err := getClient.Call(context.Background(), req, &rsp) //将返回值映射到map中
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rsp)
}
