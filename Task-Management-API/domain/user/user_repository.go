package user

import (
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByUsername(ctx context.Context, username string) (*User, error)
	FindByID(ctx context.Context, id string) (*User, error)
	UsernameExists(ctx context.Context, username string) (bool, error)
}
