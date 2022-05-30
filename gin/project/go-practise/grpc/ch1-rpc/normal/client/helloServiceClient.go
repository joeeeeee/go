package main

import (
	"fmt"
	"log"
	"net/rpc"
)

const HelloServiceName = "pkg.helloService"


type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

var _ HelloServiceInterface = (*HelloServiceClient) (nil)


type HelloServiceClient struct {
	*rpc.Client
}

func DialHelloService(network string, address string) *HelloServiceClient {
	client, err := rpc.Dial(network, address)
	if err != nil {
		log.Fatal("dialing error:", err)
	}

	return &HelloServiceClient{Client: client}
}



func main() {

	var reply string
	err := DialHelloService("tcp", "localhost:8088").Hello("client", &reply)
	if err != nil {
		log.Fatal("dial error:")
	}
	// 给对方提供RPC服务
	fmt.Printf(reply)
}


func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Call(HelloServiceName+".Hello", request, &reply)
}