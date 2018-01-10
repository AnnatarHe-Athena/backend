package services

import (
	"context"
	"log"

	"github.com/douban-girls/backend/model"
	pb "github.com/douban-girls/backend/proto"
)

type CellServer struct{}

func (this *CellServer) GetList(ctx context.Context, in *pb.PaginationRequest) (*pb.CellsReply, error) {
	log.Println("CellServer.GetList: ", in)
	cells, err := model.CellsFetchAll(in.GetFrom(), in.GetTake(), in.GetOffset(), 2)

	reply := cells.ConvertToProtoType()

	return &pb.CellsReply{
		Cells: reply,
	}, err
}
func (this *CellServer) Add(ctx context.Context, in *pb.CellItem) (*pb.CellItem, error) {
}
func (this *CellServer) Remove(ctx context.Context, in *pb.CellItem) (*pb.CommonBoolReply, error) {
}

func (this *CellServer) Destroy(ctx context.Context, in *pb.CellItem) (*pb.CommonBoolReply, error) {
}