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
	listener, err := net.Listen("tcp", ":8088")

	if err != nil {
		log.Fatal("Listen tcp error:", err)
	}

	clientChan := make(chan *rpc.Client)

	go func() {
		for {
			conn, err := listener.Accept()

			if err != nil {
				clientChan <- rpc.NewClient(conn)
			}
		}
	}()

	doClientWork(clientChan)

}

func doClientWork(clientChan <-chan *rpc.Client) {
	client := <-clientChan

	defer client.Close()
	var reply string

	err := client.Call(HelloServiceName+".Hello", "hello", &reply)

	if err != nil {
		log.Fatal("client call error :", err)
	}

	fmt.Print("reply", reply)
}


func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Call(HelloServiceName+".Hello", request, &reply)
}