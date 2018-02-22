package services

import (
	"context"

	pb "github.com/douban-girls/backend/proto"
)

type TagServer struct{}

func (this *TagServer) GetAll(ctx context.Context, in *pb.PaginationRequest) (*pb.TagList, error) {
	return nil, nil
}
func (this *TagServer) Get(ctx context.Context, in *pb.TagItem) (*pb.TagItem, error) {
	return nil, nil
}
func (this *TagServer) Add(ctx context.Context, in *pb.TagItem) (*pb.TagItem, error) {
	return nil, nil
}
func (this *TagServer) GetCells(ctx context.Context, in *pb.PaginationRequest) (*pb.CellsReply, error) {
	return nil, nil
}
