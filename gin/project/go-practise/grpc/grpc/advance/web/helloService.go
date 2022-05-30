package main

import (
	"context"
	pb "github.com/joe/iam/go-practise/grpc/grpc/advance/ca/proto"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
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

	mux := http.NewServeMux()

	hanlder := h2c.NewHandler(mux, &http2.Server{})

	server := &http.Server{
		Addr:              ":3999",
		Handler:           hanlder,
	}

	server.ListenAndServe()

}
