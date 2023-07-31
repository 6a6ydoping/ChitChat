package service

import (
	"github.com/6a6ydoping/ChitChat/internal/config"
	"github.com/6a6ydoping/ChitChat/internal/repository"
	"github.com/6a6ydoping/ChitChat/pkg/jwttoken"
)

type Manager struct {
	Repository repository.Repository
	Token      *jwttoken.Token
	Config     *config.Config
}

func New(repository repository.Repository, token *jwttoken.Token, config *config.Config) *Manager {
	return &Manager{
		Repository: repository,
		Token:      token,
		Config:     config,
	}
}
