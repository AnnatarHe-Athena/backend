package main

import (
	"log"
	"net"

	"github.com/douban-girls/backend/cfg"
	pb "github.com/douban-girls/backend/proto"
	"github.com/douban-girls/backend/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":9988"

func main() {
	log.SetPrefix("backend")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCellsServer(s, &services.CellServer{})
	pb.RegisterHelloServer(s, &services.HelloServer{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("backend server running on ", port, " in ", cfg.CONFIG.IsDev, " mode")
}
