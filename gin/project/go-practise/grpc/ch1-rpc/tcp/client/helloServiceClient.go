package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
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
	conn, err := net.Dial(network, address)
	if err != nil {
		log.Fatal("dialing error:", err)
	}

	return &HelloServiceClient{Client: rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))}
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