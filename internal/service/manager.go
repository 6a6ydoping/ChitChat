package service

import (
	"github.com/6a6ydoping/ChitChat/internal/config"
	"github.com/6a6ydoping/ChitChat/internal/repository"
	"github.com/6a6ydoping/ChitChat/pkg/jwttoken"
	"github.com/6a6ydoping/ChitChat/pkg/ws"
)

type Manager struct {
	Repository repository.Repository
	Token      *jwttoken.Token
	Config     *config.Config
	Dispatcher *ws.Dispatcher
}

func New(repository repository.Repository, token *jwttoken.Token, config *config.Config, dispatcher *ws.Dispatcher) *Manager {
	return &Manager{
		Repository: repository,
		Token:      token,
		Config:     config,
		Dispatcher: dispatcher,
	}
}
