package server

import (
	"context"

	v1 "KratosNew/api/user/v1"
	"KratosNew/internal/biz"
)

type UserService struct {
	uc *biz.UserUsecase
}

func NewGRPCServer(uc *biz.UserUsecase) *UserService {
	return &UserService{
		uc: uc,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserReply, error) {
	return nil, nil
}
