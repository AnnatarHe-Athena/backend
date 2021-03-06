package services

import (
	"context"
	"log"

	"github.com/douban-girls/backend/model"
	pb "github.com/douban-girls/backend/proto"
	"github.com/douban-girls/backend/utils"
)

type CellServer struct{}

func (this *CellServer) GetList(ctx context.Context, in *pb.PaginationRequest) (*pb.CellsReply, error) {
	log.Println("CellServer.GetList: ", in)
	permission := 2
	if in.GetHiddenOnly() {
		permission = 3
	}
	cells, err := model.CellsFetchAll(in.GetFrom(), in.GetTake(), in.GetOffset(), int32(permission))
	cells.EncodeImageURL()

	reply := cells.ConvertToProtoType()

	return &pb.CellsReply{
		Cells: reply,
	}, err
}

// Add 单条插入数据库
func (c *CellServer) Add(ctx context.Context, in *pb.CellItem) (*pb.CellItemReply, error) {
	cell := &model.Cell{
		Img:        in.GetImg(),
		Text:       in.GetText(),
		Permission: int(in.GetPermission()),
		Cate:       int(in.GetCate()),
		FromID:     in.GetFromID(),
		FromURL:    in.GetFromURL(),
		Md5:        utils.GenPassword(in.GetImg()),
	}
	log.Println(in)
	// 塞入数据完成后，需要调用 tags 服务，存入对应的 tag 中
	err := cell.Save()
	log.Println("saved cell:", *cell)
	cellsToReturn := cell.ConvertToProtoType()
	return &pb.CellItemReply{
		Cell: cellsToReturn,
	}, err
}

func (this *CellServer) Remove(ctx context.Context, in *pb.CellItem) (*pb.CommonBoolReply, error) {
	log.Println(in.GetId())
	cell := &model.Cell{ID: int(in.GetId())}
	result := cell.Remove(false)
	log.Println("ended cell remove")
	return &pb.CommonBoolReply{
		Success: result,
		// TODO: 添加报错
		Errors: []*pb.Error{},
	}, nil
}

func (this *CellServer) Destroy(ctx context.Context, in *pb.CellItem) (*pb.CommonBoolReply, error) {
	// 需要在 API 层调用七牛的服务，删掉这个真实文件
	cell := &model.Cell{ID: int(in.GetId())}
	result := cell.Remove(true)
	err := &pb.Error{
		Code: 50000,
		Msg:  "error when remove item",
	}
	return &pb.CommonBoolReply{
		Success: result,
		// TODO: 添加报错
		Errors: []*pb.Error{err},
	}, nil
}
