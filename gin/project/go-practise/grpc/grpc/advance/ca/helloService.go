package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	pb "github.com/joe/iam/go-practise/grpc/grpc/advance/ca/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
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

	certificate, err := tls.LoadX509KeyPair("../server.crt", "../server.key")
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert, // NOTE: this is optional!
		ClientCAs:    certPool,
	})

	server := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterHelloServiceServer(server, new(HelloService))

	listen, err := net.Listen("tcp", ":5000")

	if err != nil {
		log.Fatal("listen error: ", err)
	}

	err = server.Serve(listen)

	if err != nil {
		log.Fatal(err)
	}
}
