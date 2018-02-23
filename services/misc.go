package services

import (
	"context"

	"github.com/douban-girls/backend/model"
	pb "github.com/douban-girls/backend/proto"
)

type MiscServer struct{}

func (m *MiscServer) Get(ctx context.Context, in *pb.NullRequest) (*pb.MiscReturnMsg, error) {
	userCount := model.FetchUserCount()
	cellCount := model.FetchCellCount()
	return &pb.MiscReturnMsg{
		UserCount: userCount,
		CellCount: cellCount,
	}, nil
}
