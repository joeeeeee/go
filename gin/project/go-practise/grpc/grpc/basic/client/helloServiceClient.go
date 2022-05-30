package main

import (
	"context"
	"fmt"
	pb "github.com/joe/iam/go-practise/grpc/grpc/basic/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:8089", grpc.WithInsecure())

	if err != nil {
		log.Fatal("dial err", err)
	}

	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)

	stream, err := client.Channel(context.Background())

	if err != nil {
		log.Fatal("Call Channel error:", err)
	}


	go func() {
		for  {
			err2 := stream.Send(&pb.String{Value: "wuyanzhou"})
			if err2 != nil {
				log.Fatal(err2)
			}
			time.Sleep(time.Second)
		}
	}()




	for {
		recv, err2 := stream.Recv()

		if err2 != nil {
			log.Fatal("recv error:", err2)
		}

		fmt.Println(recv.GetValue())
	}


}
