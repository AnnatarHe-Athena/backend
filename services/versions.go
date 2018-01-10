package services

import (
	"context"
	"log"

	"github.com/douban-girls/backend/model"
	pb "github.com/douban-girls/backend/proto"
)

type VersionServer struct {}


func (this *VersionServer) GetAll(ctx context.Context, in *pb.PlatformRequest) (*pb.VersionList, error) {
}
func (this *VersionServer) Get(ctx context.Context, in *pb.PlatformRequest) (*pb.VersionList, error) {
}