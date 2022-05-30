package main

import (
	"context"
	"github.com/joe/iam/go-practise/grpc/grpc/advance/Token/pkg"
	pb "github.com/joe/iam/go-practise/grpc/grpc/advance/ca/proto"
	"google.golang.org/grpc"
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


type HelloService struct {
	auth pkg.Authentication
}

func (p *HelloService) Hello(ctx context.Context,reply *pb.String) (*pb.String, error) {

	if err := p.auth.Auth(ctx); err != nil {
		return nil, err
	}

	return &pb.String{Value: "hello:" + reply.GetValue()}, nil
}

func main() {
	server := grpc.NewServer()

	listen, err := net.Listen("tcp", ":8087")

	pb.RegisterHelloServiceServer(server, new(HelloService))

	err = server.Serve(listen)

	if err != nil {
		log.Fatal(err)
	}
}




