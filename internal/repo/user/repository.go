package user

import (
	"context"
	"cyberball-auth/internal/entity"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Убедимся, что repository реализует интерфейс Repository
var _ Repository = (*repository)(nil)

type Repository interface {
	// CreateUser - создание нового пользователя
	CreateUser(ctx context.Context, user *entity.User) (string, error)
	// GetUserByEmail - получение пользователя по email
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	// GetUserById - получение пользователя по id
	GetUserById(ctx context.Context, id string) (*entity.User, error)
}

// repository - репозиторий для работы с PostgreSQL
type repository struct {
	db *pgxpool.Pool
}

// NewRepository - конструктор создания репозитория для работы с PostgreSQL
func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *entity.User) (string, error) {
	userID := uuid.New().String()

	query := `INSERT INTO users (id, email, password) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, query, userID, user.Email, user.Password)
	if err != nil {
		return "", err
	}

	return userID, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `SELECT id, email, password, is_active FROM users WHERE email = $1`
	row := r.db.QueryRow(ctx, query, email)

	var user entity.User

	if err := row.Scan(&user.ID, &user.Email, &user.Password, &user.IsActive); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) GetUserById(ctx context.Context, id string) (*entity.User, error) {
	query := `SELECT id, email, password, is_active FROM users WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)

	var user entity.User

	if err := row.Scan(&user.ID, &user.Email, &user.Password, &user.IsActive); err != nil {
		return nil, err
	}
	return &user, nil
}
