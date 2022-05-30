package main

import (
	"context"
	"fmt"
	"github.com/joe/iam/go-practise/grpc/grpc/advance/Token/pkg"
	pb "github.com/joe/iam/go-practise/grpc/grpc/advance/ca/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	auth := pkg.Authentication{
		User:     "gopher",
		Password: "password",
	}

	conn, err := grpc.Dial("localhost:8087", grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth))

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
