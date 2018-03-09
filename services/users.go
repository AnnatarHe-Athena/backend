package services

import (
	"context"
	"log"

	"github.com/douban-girls/backend/model"
	pb "github.com/douban-girls/backend/proto"
)

type UserServer struct{}

func (this *UserServer) GetList(ctx context.Context, in *pb.PaginationRequest) (*pb.UserList, error) {
	return nil, nil
}
func (this *UserServer) Add(ctx context.Context, in *pb.UserItem) (*pb.UserItem, error) {
	return nil, nil
}
func (this *UserServer) Auth(ctx context.Context, in *pb.AuthRequest) (*pb.UserItem, error) {
	log.Println("user auth rpc service")
	email := in.GetEmail()
	pwd := in.GetPwd()

	user := &model.User{
		Email: email,
		Pwd:   pwd,
	}

	err := user.Auth()
	if err != nil {
		return nil, err
	}
	log.Println("auth pass user: ", *user)
	pbUser := user.ConvertToProto()
	return pbUser, err
}
func (this *UserServer) Get(ctx context.Context, in *pb.UserItem) (*pb.UserItem, error) {
	userID := in.GetId()

	user := &model.User{
		ID: int(userID),
	}

	err := user.Find()
	result := user.ConvertToProto()

	return result, err
}
func (this *UserServer) Remove(ctx context.Context, in *pb.UserItem) (*pb.UserItem, error) {
	return nil, nil
}
func (this *UserServer) Destroy(ctx context.Context, in *pb.UserItem) (*pb.CommonBoolReply, error) {
	return nil, nil
}
func (this *UserServer) Update(ctx context.Context, in *pb.UserItem) (*pb.UserItem, error) {
	return nil, nil
}
