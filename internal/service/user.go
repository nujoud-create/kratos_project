package service

import (
	"context"

	v1 "KratosNew/api/user/v1"
	"KratosNew/internal/data"
)

func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserReply, error) {
	user := &data.User{
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
	}

	if err := s.uc.Create(user); err != nil {
		return nil, err
	}

	return &v1.CreateUserReply{
		User: &v1.User{
			Id:      user.ID,
			Name:    user.Name,
			Email:   user.Email,
			Phone:   user.Phone,
			Address: user.Address,
		},
	}, nil
}

func (s *UserService) GetUsers(ctx context.Context, req *v1.GetUsersRequest) (*v1.GetUsersReply, error) {
	users, err := s.uc.GetAll()
	if err != nil {
		return nil, err
	}

	result := make([]*v1.User, 0, len(users))

	for _, user := range users {
		result = append(result, &v1.User{
			Id:      user.ID,
			Name:    user.Name,
			Email:   user.Email,
			Phone:   user.Phone,
			Address: user.Address,
		})
	}

	return &v1.GetUsersReply{
		Users: result,
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserReply, error) {
	user, err := s.uc.GetByID(req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.GetUserReply{
		User: &v1.User{
			Id:      user.ID,
			Name:    user.Name,
			Email:   user.Email,
			Phone:   user.Phone,
			Address: user.Address,
		},
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserReply, error) {
	user := &data.User{
		ID:      req.Id,
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
	}

	if err := s.uc.Update(user); err != nil {
		return nil, err
	}

	return &v1.UpdateUserReply{
		User: &v1.User{
			Id:      user.ID,
			Name:    user.Name,
			Email:   user.Email,
			Phone:   user.Phone,
			Address: user.Address,
		},
	}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserReply, error) {
	if err := s.uc.Delete(req.Id); err != nil {
		return nil, err
	}

	return &v1.DeleteUserReply{}, nil
}
