package services

import (
	"context"
	"log"

	"github.com/douban-girls/backend/model"
	pb "github.com/douban-girls/backend/proto"
)

type CollectionServer struct{}

func (this *CollectionServer) GetList(ctx context.Context, in *pb.PaginationRequest) (*pb.CellsReply, error) {

	user := &model.User{
		ID: int(in.GetUserId()),
	}
	cells, err := user.Collections(int(in.GetFrom()), int(in.GetTake()))
	cells.EncodeImageURL()

	reply := cells.ConvertToProtoType()

	return &pb.CellsReply{Cells: reply}, err
}
func (this *CollectionServer) Add(ctx context.Context, in *pb.CollectionItem) (*pb.CommonBoolReply, error) {
	log.Println("collection rpc got: ", *in)
	collections := model.Collections{
		&model.Collection{Owner: int(in.GetOwner()), Cell: int(in.GetCell())},
	}
	err := collections.Save()
	log.Println(collections)
	return &pb.CommonBoolReply{
		Success: err == nil,
		Errors:  nil,
	}, err

}
func (this *CollectionServer) Remove(ctx context.Context, in *pb.CollectionItem) (*pb.CommonBoolReply, error) {
	// TODO
	return nil, nil
}
