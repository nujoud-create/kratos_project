package biz

import "KratosNew/internal/data"

type UserUsecase struct {
	repo *data.UserRepo
}

func NewUserUsecase(repo *data.UserRepo) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uc *UserUsecase) Create(user *data.User) error {
	return uc.repo.Create(user)
}

func (uc *UserUsecase) GetAll() ([]data.User, error) {
	return uc.repo.GetAll()
}

func (uc *UserUsecase) GetByID(id uint64) (*data.User, error) {
	return uc.repo.GetByID(id)
}

func (uc *UserUsecase) Update(user *data.User) error {
	return uc.repo.Update(user)
}

func (uc *UserUsecase) Delete(id uint64) error {
	return uc.repo.Delete(id)
}