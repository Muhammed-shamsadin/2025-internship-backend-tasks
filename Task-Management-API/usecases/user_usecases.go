package usecases

import (
	"2025-internship-backend-tasks/Task-Management-API/domain/user"
	"2025-internship-backend-tasks/Task-Management-API/infrastructure"
	"context"
	"errors"
)

type UserUsecase struct {
	userRepo user.UserRepository
}

func NewUserUsecase(userRepo user.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (u *UserUsecase) RegisterUser(ctx context.Context, usr *user.User) error {
	exists, err := u.userRepo.UsernameExists(ctx, usr.Username)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("username already exists")
	}
	return u.userRepo.Create(ctx, usr)
}

func (u *UserUsecase) GetUserByUsername(ctx context.Context, username string) (*user.User, error) {
	return u.userRepo.FindByUsername(ctx, username)
}

func (u *UserUsecase) UsernameExists(ctx context.Context, username string) (bool, error) {
	return u.userRepo.UsernameExists(ctx, username)
}

func (u *UserUsecase) LoginUser(ctx context.Context, username, password string) (*user.User, error) {
	usr, err := u.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if !infrastructure.CheckPasswordHash(password, usr.Password) {
		return nil, errors.New("invalid credentials")
	}

	return usr, nil
}
