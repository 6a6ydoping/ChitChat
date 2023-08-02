package service

import (
	"context"
	"github.com/6a6ydoping/ChitChat/internal/entity"
)

type UserService interface {
	// User
	CreateUser(ctx context.Context, u *entity.User) error
	Login(ctx context.Context, username, password string) (string, error)
	VerifyToken(token string) (int64, error)
}
