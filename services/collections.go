package services

import (
	"context"

	"github.com/douban-girls/backend/model"
	pb "github.com/douban-girls/backend/proto"
)

type CollectionServer struct{}

func (this *CollectionServer) GetList(ctx context.Context, in *pb.PaginationRequest) (*pb.CellsReply, error) {
	cells, err := model.FetchUserCollectionBy(in.GetUserId(), in.GetFrom(), in.GetOffset())
	reply := cells.ConvertToProtoType()

	return &pb.CellsReply{Cells: reply}, err
}
func (this *CollectionServer) Add(ctx context.Context, in *pb.CollectionItem) (*pb.CommonBoolReply, error) {
	collections := model.Collections{
		&model.Collection{Owner: in.GetOwner(), Cell: in.GetCell()},
	}
	err := collections.Save()
	return nil, &pb.CommonBoolReply{
		Success: err == nil,
		Errors:  []error{err},
	}, err

}
func (this *CollectionServer) Remove(ctx context.Context, in *pb.CollectionItem) (*pb.CommonBoolReply, error) {
	// TODO
}
