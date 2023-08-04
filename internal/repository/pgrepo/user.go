package pgrepo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/6a6ydoping/ChitChat/internal/entity"
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

func (p *Postgres) GetUser(ctx context.Context, username string) (*entity.User, error) {
	query := fmt.Sprintf(`SELECT id, username, password FROM %s WHERE username = $1`, usersTable)
	row := p.Pool.QueryRow(ctx, query, username)
	var dbUser entity.User
	err := row.Scan(&dbUser.ID, &dbUser.Username, &dbUser.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Invalid username or password")
		}
		return nil, err
	}

	return &dbUser, nil
}

func (p *Postgres) UpdateUser(ctx context.Context, user *entity.User) error {
	query := fmt.Sprintf(`
		UPDATE %s
		SET username = $2, password = $3
		WHERE id = $1
	`, usersTable)

	_, err := p.Pool.Exec(ctx, query, user.ID, user.Username, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) DeleteUser(ctx context.Context, userID int64) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, usersTable)

	_, err := p.Pool.Exec(ctx, query, userID)
	if err != nil {
		return err
	}

	return nil
}
