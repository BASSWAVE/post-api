package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

const postgresUrl = "postgres://user:pass@db:5432/postgres?sslmode=disable"

func NewPool() (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), postgresUrl)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
