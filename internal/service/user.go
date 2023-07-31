package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/6a6ydoping/ChitChat/internal/entity"
	"github.com/6a6ydoping/ChitChat/pkg/util"
)

func (m *Manager) CreateUser(ctx context.Context, u *entity.User) error {
	hashedPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword
	fmt.Println(u)
	err = m.Repository.CreateUser(ctx, u)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) Login(ctx context.Context, username, password string) (string, error) {
	user, err := m.Repository.GetUser(ctx, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("user not found")
		}

		return "", fmt.Errorf("get user err: %w", err)
	}

	err = util.CheckPassword(password, user.Password)
	if err != nil {
		return "", fmt.Errorf("incorrect password: %w", err)
	}

	accessToken, err := m.Token.GenerateJWT(user.ID, m.Config.Token.TimeToLive)
	if err != nil {
		return "", fmt.Errorf("create token err: %w", err)
	}

	return accessToken, nil
}

func (m *Manager) VerifyToken(token string) (int64, error) {
	//payload, err := m.Token.ValidateToken(token)
	//if err != nil {
	//	return 0, fmt.Errorf("validate token err: %w", err)
	//}
	//
	//return payload.UserID, nil
	return 0, nil
}
