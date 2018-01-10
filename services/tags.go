package services

import (
	"context"
	"log"

	"github.com/douban-girls/backend/model"
	pb "github.com/douban-girls/backend/proto"
)


type TagServer struct {}

func (this *TagServer) GetAll(ctx context.Context, in *pb.PaginationRequest) (*pb.TagList, error) {
}
func (this *TagServer) Get(ctx context.Context, in *pb.TagItem) (*pb.TagItem, error) {
}
func (this *TagServer) Add(ctx context.Context, in *pb.TagItem) (*pb.TagItem, error) {
}
func (this *TagServer) GetCells(ctx context.Context, in *pb.PaginationRequest) (*pb.CellsReply, error) {
}