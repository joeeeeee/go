package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

const HelloServiceName = "store"

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

	go func() {
		for  {
			DialHelloService("tcp", "localhost:8088").Watch(30, &reply)
			fmt.Println("watch:", reply)
		}
	}()

	time.Sleep(time.Second * 1)
	var setReply string
	err := DialHelloService("tcp", "localhost:8088").Set([2]string{"name", "111"}, &setReply)
	//err := DialHelloService("tcp", "localhost:8088").Get("name", &reply)
	if err != nil {
		log.Fatal("dial error:")
	}
	// 给对方提供RPC服务
	fmt.Println("reply:", setReply)

	time.Sleep(time.Second * 2)
	DialHelloService("tcp", "localhost:8088").Set([2]string{"name", "wuyanzhous"}, &setReply)

	// 给对方提供RPC服务
	fmt.Println("reply:", setReply)

	time.Sleep(time.Second * 2)
	DialHelloService("tcp", "localhost:8088").Set([2]string{"name", "wuyanzhousss"}, &setReply)

	// 给对方提供RPC服务
	fmt.Println("reply:", setReply)

	time.Sleep(time.Second * 100)
}

func (p *HelloServiceClient) Set(kv [2]string, reply *string) error {
	return p.Call(HelloServiceName+".Set", kv, &reply)
}

func (p *HelloServiceClient) Get(key string, reply *string) error {
	return p.Call(HelloServiceName+".Get", key, &reply)
}

func (p *HelloServiceClient) Watch(time int, reply *string) error {
	return p.Call(HelloServiceName+".Watch", time, &reply)
}
