package main

import (
	"log"
	"net"

	pb "github.com/douban-girls/backend/proto"
	"github.com/douban-girls/backend/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":9988"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCellsServer(s, &services.CellServer{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	println("back end server running on ", port)
}
