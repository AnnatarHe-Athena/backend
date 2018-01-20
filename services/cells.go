package services

import (
	"context"
	"errors"
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
func (this *CellServer) Add(ctx context.Context, in *pb.CellItem) (*pb.CellsReply, error) {
	cs := model.Cells{
		&model.Cell{
			Img:        in.GetImg(),
			Text:       in.GetText(),
			Permission: in.GetPermission(),
			Cate:       in.GetCate(),
			FromID:     in.GetFromID(),
			FromURL:    in.GetFromURL(),
		},
	}
	// 塞入数据完成后，需要调用 tags 服务，存入对应的 tag 中
	err := cs.Save()
	return &pb.CellsReply{
		Cells: cs,
	}, err

}
func (this *CellServer) Remove(ctx context.Context, in *pb.CellItem) (*pb.CommonBoolReply, error) {
	cell := &model.Cell{ID: in.GetId()}
	result := cell.Remove()
	return &pb.CommonBoolReply{
		Success: result,
		Errors:  []error{errors.New("error when delete item")},
	}, nil
}

func (this *CellServer) Destroy(ctx context.Context, in *pb.CellItem) (*pb.CommonBoolReply, error) {
	// 需要在 API 层调用七牛的服务，删掉这个真实文件
	cell := &model.Cell{ID: in.GetId()}
	result := cell.Remove(true)
	return nil, &pb.CommonBoolReply{
		Success: result,
		Errors:  []error{errors.New("error when delete item")},
	}
}
