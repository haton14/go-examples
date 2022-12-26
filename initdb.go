package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func initDB() (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig("postgres://postgres:postgres@localhost:15432/postgres?sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("pgxpool.ParseConfig(): err %w", err)
	}

	pool, err := pgxpool.NewWithConfig(context.TODO(), config)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.NewWithConfig(): err %w", err)
	}

	if err := pool.Ping(context.TODO()); err != nil {
		return nil, fmt.Errorf("pgxpool.Pint(): err %w", err)
	}
	return pool, nil
}
