package main

import (
	"context"
	"fmt"
	pb "github.com/joe/iam/go-practise/grpc/grpc/pub/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewPublishServiceClient(conn)

	reply, err := client.Publish(context.Background(), &pb.String{Value: "golang: hello golang"})


	fmt.Println(reply.GetValue())
}
