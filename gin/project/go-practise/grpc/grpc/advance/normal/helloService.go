package main

import (
	"context"
	pb "github.com/joe/iam/go-practise/grpc/grpc/advance/ca/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
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

func (p *HelloService) Hello(ctx context.Context,reply *pb.String) (*pb.String, error) {
	return &pb.String{Value: "hello:" + reply.GetValue()}, nil
}

func main() {

	creds, err := credentials.NewServerTLSFromFile("../server.crt", "../server.key")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer(grpc.Creds(creds))

	reflection.Register(server)

	listen, err := net.Listen("tcp", ":8087")

	pb.RegisterHelloServiceServer(server, new(HelloService))

	err = server.Serve(listen)

	if err != nil {
		log.Fatal(err)
	}

}
