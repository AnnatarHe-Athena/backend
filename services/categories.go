package services

import (
	"context"
	"log"

	"github.com/douban-girls/backend/model"
	pb "github.com/douban-girls/backend/proto"
)

type CategoryServer struct {}

func (this *CategoryServer) GetList(ctx context.Context, in *pb.PaginationRequest) (*pb.CategoryReply, error) {
	
}
func (this *CategoryServer) Add(ctx context.Context, in *pb.CategoryItem) (*pb.CellsReply, error) {

}


