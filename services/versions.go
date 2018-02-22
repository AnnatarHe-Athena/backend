package services

import (
	"context"

	pb "github.com/douban-girls/backend/proto"
)

type VersionServer struct{}

func (this *VersionServer) GetAll(ctx context.Context, in *pb.PlatformRequest) (*pb.VersionList, error) {
	return nil, nil
}
func (this *VersionServer) Get(ctx context.Context, in *pb.PlatformRequest) (*pb.VersionList, error) {
	return nil, nil
}
