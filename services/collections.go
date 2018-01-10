package services

import (
	"context"
	"log"

	"github.com/douban-girls/backend/model"
	pb "github.com/douban-girls/backend/proto"
)

type CollectionServer struct {}


func (this *CollectionServer) GetList(ctx context.Context, in *pb.PaginationRequest) (*pb.CellsReply, error) {
}
func (this *CollectionServer) Add(ctx context.Context, in *pb.CollectionItem) (*pb.CommonBoolReply, error) {
}
func (this *CollectionServer) Remove(ctx context.Context, in *pb.CollectionItem) (*pb.CommonBoolReply, error) {
}