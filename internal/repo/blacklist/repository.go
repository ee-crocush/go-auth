package blacklist

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Убедимся, что repository реализует интерфейс Repository
var _ Repository = (*repository)(nil)

// Repository - интерфейс для работы с blacklist
type Repository interface {
	// AddToBlacklist - добавление токена в blacklist
	AddToBlacklist(ctx context.Context, token string) error
	// IsTokenBlacklisted - проверка, находится ли токен в blacklist
	IsTokenBlacklisted(ctx context.Context, token string) (bool, error)
}

// repository - структура репозитория для работы с PostgreSQL
type repository struct {
	db *pgxpool.Pool
}

// NewRepository - конструктор для repository
func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

// AddToBlacklist - добавление токена в blacklist
func (r *repository) AddToBlacklist(ctx context.Context, token string) error {
	query := `INSERT INTO blacklist (access_token) VALUES ($1)`
	_, err := r.db.Exec(ctx, query, token)

	return err
}

// IsTokenBlacklisted - проверка, находится ли токен в blacklist
func (r *repository) IsTokenBlacklisted(ctx context.Context, token string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM blacklist WHERE access_token = $1)`
	err := r.db.QueryRow(ctx, query, token).Scan(&exists)

	if err != nil {
		return false, err
	}
	return exists, nil
}
