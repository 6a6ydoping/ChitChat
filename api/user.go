package api

import "github.com/6a6ydoping/ChitChat/internal/entity"

type RegisterRequest struct {
	entity.User
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
