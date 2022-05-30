package main

import (
	"log"
	"net"
	"net/rpc"
)

const HelloServiceName = "pkg.helloService"

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(srv HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, srv)
}


type HelloService struct {}

func (p *HelloService) Hello (request string, reply *string) error {
	* reply = "hello :" + request

	return nil
}

func main() {
	// 将对象类型中所有满足RPC规则的对象方法注册为RPC函数
	//rpc.RegisterName("HelloService", new (HelloService))
	RegisterHelloService(new(HelloService))

	for {
		conn, err := net.Dial("tcp", ":8088")

		if err != nil {
			log.Fatal("Accept error:", err)
		}

		rpc.ServeConn(conn)
	}


}
