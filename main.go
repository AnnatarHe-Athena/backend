package main

import (
	"log"
	"net"

	"github.com/douban-girls/backend/cfg"
	"github.com/douban-girls/backend/model"
	pb "github.com/douban-girls/backend/proto"
	"github.com/douban-girls/backend/services"
	raven "github.com/getsentry/raven-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.SetPrefix("backend | ")
	lis, err := net.Listen("tcp", cfg.CONFIG.Port)
	onAppStart()
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCellsServer(s, &services.CellServer{})
	pb.RegisterHelloServer(s, &services.HelloServer{})
	pb.RegisterCollectionsServer(s, &services.CollectionServer{})
	pb.RegisterTagsServer(s, &services.TagServer{})
	pb.RegisterUsersServer(s, &services.UserServer{})
	pb.RegisterVersionsServer(s, &services.VersionServer{})
	pb.RegisterCategoriesServer(s, &services.CategoryServer{})
	pb.RegisterMiscServer(s, &services.MiscServer{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("backend server running on ", cfg.CONFIG.Port, " in ", cfg.CONFIG.IsDev, " mode")
}

func onAppStart() {
	if !cfg.CONFIG.IsDev {
		raven.SetDSN(cfg.CONFIG.RavenDSN)
	}
	model.DatabaseInit()
}
