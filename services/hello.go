package services

import (
	"context"

	pb "github.com/douban-girls/backend/proto"
	logs "github.com/sirupsen/logrus"
)

type HelloServer struct{}

func (s *HelloServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	logs.Info("recive a rpc request in HelloServer")

	name := in.GetName()

	return &pb.HelloReply{
		Message: "hello " + name,
	}, nil
}
