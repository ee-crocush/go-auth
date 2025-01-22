package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"go-auth/config"
	"time"
)

// New - функция для создания подключения к базе данных
func New(ctx context.Context, cfg config.Config) (*pgxpool.Pool, error) {
	url := cfg.PG.URL()

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("failed to parse PostgreSQL DSN: %w", err)
	}

	poolConfig.MaxConns = cfg.PG.MaxConns
	poolConfig.ConnConfig.ConnectTimeout = cfg.PG.ConnTimeout * time.Second

	db, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create PostgreSQL connection pool: %w", err)
	}
	return db, nil
}
