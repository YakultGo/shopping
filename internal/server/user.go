package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	userPb "shopping/api/user"
	"shopping/internal/data/model"
	"shopping/internal/data/query"
	"time"
)

type UserServer struct {
	userPb.UnimplementedUserServer
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

// GetUser 根据用户ID获取用户信息
func (us *UserServer) GetUser(ctx context.Context, req *userPb.GetUserRequest) (*userPb.GetUserResponse, error) {
	u := query.User
	user, err := u.WithContext(ctx).Where(u.ID.Eq(req.Id)).First()
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "[GetUser] user not found")
	}
	resp := &userPb.GetUserResponse{
		Id:        user.ID,
		Name:      user.Name,
		Telephone: user.Telephone,
	}
	if user.Address != nil {
		resp.Address = *user.Address
	}
	if user.Birthday != nil {
		resp.Birthday = user.Birthday.String()
	}
	return resp, nil
}

// CreateUser 创建用户
func (us *UserServer) CreateUser(ctx context.Context, req *userPb.CreateUserRequest) (*userPb.CreateUserResponse, error) {
	u := query.User
	user := &model.User{
		Name:      req.Name,
		Telephone: req.Telephone,
	}
	err := u.WithContext(ctx).Create(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "[CreateUser] create user failed")
	}
	return &userPb.CreateUserResponse{
		Id: user.ID,
	}, nil
}
func (us *UserServer) UpdateUser(ctx context.Context, req *userPb.UpdateUserRequest) (*userPb.UpdateUserResponse, error) {
	u := query.User
	// 将时间转换为time.Time类型
	birthday, err := time.Parse("2006-01-02", req.Birthday)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "[UpdateUser] birthday format error")
	}
	result, err := u.WithContext(ctx).Where(u.ID.Eq(req.Id)).Updates(model.User{
		Name:      req.Name,
		Telephone: req.Telephone,
		Address:   &req.Address,
		Birthday:  &birthday,
		Password:  &req.Password,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "[UpdateUser] update user failed")
	}
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "[UpdateUser] user not found")
	}
	return &userPb.UpdateUserResponse{Id: req.Id}, nil
}

func (us *UserServer) DeleteUser(ctx context.Context, req *userPb.DeleteUserRequest) (*userPb.DeleteUserResponse, error) {
	u := query.User
	result, err := u.WithContext(ctx).Where(u.ID.Eq(req.Id)).Delete()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "[DeleteUser] delete user failed")
	}
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "[DeleteUser] user not found")
	}
	return &userPb.DeleteUserResponse{Id: req.Id}, nil
}
