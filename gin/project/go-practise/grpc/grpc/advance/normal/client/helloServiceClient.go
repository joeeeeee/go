package main

import (
	"context"
	"fmt"
	pb "github.com/joe/iam/go-practise/grpc/grpc/advance/ca/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("../../server.crt", "server.io")

	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial("localhost:8087", grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatal("dial error : ", err)
	} else {
		log.Println("dial success")
	}

	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)

	reply, err := client.Hello(context.Background(), &pb.String{Value: "hello golang"})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
