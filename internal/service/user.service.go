package service

import "go_ecommerce_backend/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetInforUserService() string {
	return us.userRepo.GetInforUserRepo()
}
