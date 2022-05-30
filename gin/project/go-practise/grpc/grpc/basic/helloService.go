package main

import (
	"context"
	pb "github.com/joe/iam/go-practise/grpc/grpc/basic/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *pb.String) (*pb.String, error) {
	reply := &pb.String{Value: "Hello : " + args.GetValue()}

	return reply, nil
}

func (p *HelloServiceImpl) Channel (stream pb.HelloService_ChannelServer) error {
	for {
		recv, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				return nil
			}
			log.Fatal("recv error", err)
		}

		err = stream.Send(&pb.String{Value: recv.GetValue()})

		if err != nil {
			log.Fatal("send error", err)
		}
	}
}

func main() {
	grpcServer := grpc.NewServer()

	pb.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	listen, err := net.Listen("tcp", ":8089")

	if err != nil {
		log.Fatal("listen tcp error :", err)
	}

	grpcServer.Serve(listen)

}
