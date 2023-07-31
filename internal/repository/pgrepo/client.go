package pgrepo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"net/url"
)

const usersTable = "users"

type Postgres struct {
	host     string
	username string
	password string
	port     string
	dbName   string
	Pool     *pgxpool.Pool
	SQLDB    *sql.DB
}

func New(opts ...Option) (*Postgres, error) {
	p := new(Postgres)

	for _, opt := range opts {
		opt(p)
	}

	q := url.Values{}
	q.Add("sslmode", "disable")

	u := url.URL{
		Scheme:   "postgresql",
		User:     url.UserPassword(p.username, p.password),
		Host:     fmt.Sprintf("%s:%s", p.host, p.port),
		Path:     p.dbName,
		RawQuery: q.Encode(),
	}

	poolConfig, err := pgxpool.ParseConfig(u.String())
	if err != nil {
		return nil, fmt.Errorf("pgxpool parse config err: %w", err)
	}

	p.Pool, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("pgxpool connect err: %w", err)
	}

	return p, nil
}

func NewSQL(opts ...Option) (*Postgres, error) {
	p := new(Postgres)

	for _, opt := range opts {
		opt(p)
	}

	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", p.host, p.port, p.username, p.password, p.dbName)
	db, err := sql.Open("postgres", connect)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	p.SQLDB = db

	return p, nil
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
