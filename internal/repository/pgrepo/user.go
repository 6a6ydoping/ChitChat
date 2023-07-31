package pgrepo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/6a6ydoping/ChitChat/internal/entity"
	"strconv"
	"time"
)

func (p *Postgres) CreateUser(ctx context.Context, u *entity.User) error {
	query := fmt.Sprintf(`
			INSERT INTO %s (
			                username, -- 1 
			                password -- 2
			                )
			VALUES ($1, $2)
			`, usersTable)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := p.Pool.Exec(ctx, query, u.Username, u.Password)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) Login(ctx context.Context, username, password string) (string, error) {
	query := fmt.Sprintf(`SELECT id, username, password FROM %s WHERE username = $1`, usersTable)
	row := p.Pool.QueryRow(ctx, query, username)
	var dbUser entity.User
	err := row.Scan(&dbUser.ID, &dbUser.Username, &dbUser.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("Invalid username or password")
		}
		return "", err
	}
	if dbUser.Password != password {
		return "", fmt.Errorf("Invalid username or password")
	}

	return strconv.Itoa(int(dbUser.ID)), nil
}

func (p *Postgres) GetUser(ctx context.Context, username string) (*entity.User, error) {
	return nil, nil
}
