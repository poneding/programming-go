package handler

import (
	"context"
	"user/internal/repository"
	service "user/internal/service/pb"
	"user/pkg/e"
)

type UserService struct {
	service.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) UserLogin(ctx context.Context, req *service.UserRequest) (*service.UserDetailResponse, error) {
	var u repository.User
	resp := new(service.UserDetailResponse)

	resp.Code = e.Success

	if err := u.ShowUserInfo(req); err != nil {
		resp.Code = e.Error
		return resp, err
	}

	resp.UserDetail = repository.BuildUser(u)
	return resp, nil
}

func (*UserService) UserRegister(ctx context.Context, req *service.UserRequest) (*service.UserDetailResponse, error) {
	var user repository.User
	resp := new(service.UserDetailResponse)
	resp.Code = e.Success
	if err := user.Create(req); err != nil {
		resp.Code = e.Error
		return resp, err
	}
	resp.UserDetail = repository.BuildUser(user)
	return resp, nil
}

func (*UserService) UserLogout(ctx context.Context, req *service.UserRequest) (*service.UserDetailResponse, error) {
	return new(service.UserDetailResponse), nil
}
