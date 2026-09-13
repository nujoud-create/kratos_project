package biz

type UserRepo interface{}

type UserUsecaseImpl struct {
	repo UserRepo
}

func NewUserUsecaseWithRepo(repo UserRepo) *UserUsecaseImpl {
	return &UserUsecaseImpl{
		repo: repo,
	}
}
