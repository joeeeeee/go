package main

import (
	"context"
	"fmt"
	pb "github.com/joe/iam/go-practise/grpc/grpc/pub/proto"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewPublishServiceClient(conn)

	subscribe, err := client.Subscribe(context.Background(), &pb.String{Value: "golang"})

	for {
		recv, err := subscribe.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("end of file")
				break
			}
			log.Fatal("recv error:", err)
		}

		fmt.Println(recv.GetValue())
	}



}
