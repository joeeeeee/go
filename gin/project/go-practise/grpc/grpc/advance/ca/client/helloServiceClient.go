package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	pb "github.com/joe/iam/go-practise/grpc/grpc/advance/ca/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

var tlsServerName =  "server.io"

func main() {

	crets, err := tls.LoadX509KeyPair("../../client.crt", "../../client.key")

	if err != nil {
		log.Fatal("tls load err", err)
	}

	certPool := x509.NewCertPool()

	ca, err := ioutil.ReadFile("../../ca.crt")

	if err != nil {
		log.Fatal(err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("fail to append ca certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates:       []tls.Certificate{crets},
		ServerName:         tlsServerName, // NOTE: this is required!
		RootCAs:            certPool,
	})

	conn, err := grpc.Dial("localhost:8098", grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)

	reply, err := client.Hello(context.Background(), &pb.String{Value: "golang"})

	if err != nil {
		log.Fatal(err)
	}


	fmt.Println(reply)

}
