package repository

import (
	"context"
	"github.com/6a6ydoping/ChitChat/internal/entity"
)

type Repository interface {
	// User
	CreateUser(ctx context.Context, u *entity.User) error
	GetUser(ctx context.Context, username string) (*entity.User, error)
}
