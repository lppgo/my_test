/**
 * @Author: lucas
 * @Description:
 * @File:  sidecar.go
 * @Version: 1.0.0
 * @Date: 2020/8/18 10:51
 */
package sidecar

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type JSONRequest struct {
	Jsonrpc string
	Method  string
	Params  []*Service
	Id      int
}

func NewJSONRequest(service *Service, endpoint string) *JSONRequest {
	return &JSONRequest{Jsonrpc: "2.0", Method: endpoint, Params: []*Service{service}, Id: 1}
}

type Service struct {
	Name  string
	Nodes []*ServiceNode
}
type ServiceNode struct {
	Id      string //服务ID,不能重复
	Port    int
	Address string
}

func NewService(name string) *Service {
	return &Service{Name: name, Nodes: make([]*ServiceNode, 0)}
}
func NewServiceNode(id string, port int, address string) *ServiceNode {
	return &ServiceNode{Id: id, Port: port, Address: address}
}
func (this *Service) AddNode(id string, port int, address string) {
	this.Nodes = append(this.Nodes, NewServiceNode(id, port, address))
}

var RegistryURI = "http://localhost:8000"

func requestRegistry(jsonrequest *JSONRequest) error { //关键代码。用来请求注册器
	b, err := json.Marshal(jsonrequest)
	if err != nil {
		log.Fatal(err)
		return err
	}
	rsp, err := http.Post(RegistryURI, "application/json", bytes.NewReader(b)) //发送Post请求到Registry的地址带上我们的json数据，就可以成功注册啦
	if err != nil {
		return err
	}
	defer rsp.Body.Close()
	res, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(res)) //打印出结果
	return nil
}

func UnRegService(service *Service) error {
	return requestRegistry(NewJSONRequest(service, "Registry.Deregister")) //反注册只需要改一下endpoint为Registry.Deregister就可以了
}
func RegService(service *Service) error {
	return requestRegistry(NewJSONRequest(service, "Registry.Register"))
}
