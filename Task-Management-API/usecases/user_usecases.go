package usecases

import (
	"2025-internship-backend-tasks/Task-Management-API/domain/user"
	"context"
)

type UserUsecase struct {
	userRepo user.UserRepository
}

func NewUserUsecase(userRepo user.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (u *UserUsecase) RegisterUser(ctx context.Context, usr *user.User) error {
	return u.userRepo.Create(ctx, usr)
}

func (u *UserUsecase) GetUserByUsername(ctx context.Context, username string) (*user.User, error) {
	return u.userRepo.FindByUsername(ctx, username)
}

func (u *UserUsecase) UsernameExists(ctx context.Context, username string) (bool, error) {
	return u.userRepo.UsernameExists(ctx, username)
}
