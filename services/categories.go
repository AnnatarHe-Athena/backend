package services

import (
	"context"
	"errors"

	"github.com/douban-girls/backend/model"
	pb "github.com/douban-girls/backend/proto"
)

type CategoryServer struct{}

func (this *CategoryServer) GetList(ctx context.Context, in *pb.PaginationRequest) (*pb.CategoryReply, error) {
	cates, err := model.FetchAllCategories()
	result := cates.Convert()
	return &pb.CategoryReply{
		Categories: result,
	}, err
}
func (this *CategoryServer) Add(ctx context.Context, in *pb.CategoryItem) (*pb.CellsReply, error) {
	// 不再支持，请使用新的 tag 类
	return nil, errors.New("deprecated method, instead of tags service plz")

}
