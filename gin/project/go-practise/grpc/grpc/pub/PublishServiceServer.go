package main

import (
	"context"
	"fmt"
	pb "github.com/joe/iam/go-practise/grpc/grpc/pub/proto"
	"github.com/moby/moby/pkg/pubsub"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
	"time"
)

type PublishServiceImpl struct {
	Publisher *pubsub.Publisher
}

func NewPublishServiceImpl() *PublishServiceImpl {
	return &PublishServiceImpl{
		Publisher: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

func (p *PublishServiceImpl) Publish(ctx context.Context, args *pb.String) (*pb.String, error) {
	p.Publisher.Publish(args.GetValue())

	fmt.Println("Publish:", args.GetValue())

	return &pb.String{Value: "SUCCESS"}, nil
}

func (p *PublishServiceImpl) Subscribe(args *pb.String, stream pb.PublishService_SubscribeServer) error {

	topic := p.Publisher.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, args.GetValue()) {
				return true
			}
		}
		return false
	})
	for t := range topic {
		err := stream.Send(&pb.String{Value: t.(string)})
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	server := grpc.NewServer()

	//注册服务
	pb.RegisterPublishServiceServer(server, NewPublishServiceImpl())

	// 监听端口
	listen, err := net.Listen("tcp", ":8090")

	if err != nil {
		log.Fatal(err)
	}

	server.Serve(listen)
}
