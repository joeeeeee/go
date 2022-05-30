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
	err2 := RegisterHelloService(new(HelloService))

	if err2 != nil {
		log.Fatal("register error :", err2)
	}

	// 建立TCP链接
	listen, err := net.Listen("tcp", ":8088")

	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}


	for {
		conn, err := listen.Accept()

		if err != nil {
			log.Fatal("Accept error:", err)
		}

		// 给对方提供RPC服务
		go rpc.ServeConn(conn)
	}

}
