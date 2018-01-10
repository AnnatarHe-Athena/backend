package services

import (
	"context"
	"log"

	"github.com/douban-girls/backend/model"
	pb "github.com/douban-girls/backend/proto"
)

type UserServer struct {}

func (this *UserServer) GetList(ctx context.Context, in *pb.PaginationRequest) (*pb.UserList, error) {
}
func (this *UserServer) Add(ctx context.Context, in *pb.UserItem) (*pb.UserItem, error) {
}
func (this *UserServer) Auth(ctx context.Context, in *pb.AuthRequest) (*pb.UserItem, error) {
}
func (this *UserServer) Get(ctx context.Context, in *pb.UserItem) (*pb.UserItem, error) {
}
func (this *UserServer) Remove(ctx context.Context, in *pb.UserItem) (*pb.UserItem, error) {
}
func (this *UserServer) Destroy(ctx context.Context, in *pb.UserItem) (*pb.CommonBoolReply, error) {
}
func (this *UserServer) Update(ctx context.Context, in *pb.UserItem) (*pb.UserItem, error) {
}
