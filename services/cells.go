package services

import (
	"context"
	"log"

	"github.com/douban-girls/backend/model"
	pb "github.com/douban-girls/backend/proto"
)

type CellServer struct{}

func (s *CellServer) GetList(ctx context.Context, in *pb.CellsRequest) (*pb.CellsReply, error) {
	log.Println("CellServer.GetList: ", in)
	cells, err := model.CellsFetchAll(in.GetFrom(), in.GetTake(), in.GetOffset(), 2)

	reply := cells.ConvertToProtoType()

	return &pb.CellsReply{
		Cells: reply,
	}, err
}
