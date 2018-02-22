package services

import (
	"context"
	"log"

	pb "github.com/douban-girls/backend/proto"
)

type HelloServer struct{}

func (s *HelloServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("recive a rpc request in HelloServer")

	name := in.GetName()

	return &pb.HelloReply{
		Message: "hello " + name,
	}, nil
}
