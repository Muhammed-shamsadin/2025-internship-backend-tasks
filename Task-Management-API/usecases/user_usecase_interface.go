package usecases

import (
	"2025-internship-backend-tasks/Task-Management-API/domain/user"
	"context"
)

type UserUsecaseInterface interface {
	RegisterUser(ctx context.Context, usr *user.User) error
	GetUserByUsername(ctx context.Context, username string) (*user.User, error)
	UsernameExists(ctx context.Context, username string) (bool, error)
	LoginUser(ctx context.Context, username, password string) (*user.User, error)
}
