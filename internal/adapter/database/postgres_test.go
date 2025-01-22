package database

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go-auth/config"
	"os"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	// Создаем контекст
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Загружаем переменные окружения из .env файла
	if err := godotenv.Load("../../../.env"); err != nil {
		t.Fatal("failed to load .env file")
	}

	// Заглушка для конфигурации
	cfg := config.Config{
		PG: config.PGConfig{
			Host:        os.Getenv("PG_HOST"),
			Port:        5433,
			User:        os.Getenv("PG_USER"),
			Password:    os.Getenv("PG_PASSWORD"),
			DbName:      os.Getenv("PG_DBNAME"),
			MaxConns:    5,
			ConnTimeout: 3,
		},
	}

	// Пытаемся создать подключение
	db, err := New(ctx, cfg)
	if err != nil {
		t.Fatalf("failed to create PostgreSQL connection pool: %v", err)
	}

	// Убедимся, что подключение успешно создано
	if db == nil {
		t.Fatal("expected a valid connection pool, got nil")
	}

	// Закрываем подключение после теста
	defer db.Close()

	// Проверяем возможность выполнения запроса
	err = db.Ping(ctx)
	if err != nil {
		t.Fatalf("failed to ping PostgreSQL: %v", err)
	}

	assert.NotNil(t, db)
}
