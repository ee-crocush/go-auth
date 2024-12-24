package tests

import (
	"cyberball-auth/config"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	//"time"
)

// Подготовка .env файла для теста
func TestNewConfig(t *testing.T) {
	// Загружаем конфигурацию
	cfg, err := config.NewConfig()

	assert.NoError(t, err, "config error")
	// Печатаем конфигурацию для отладки

	assert.NotEmpty(t, cfg.App.Name, "Имя приложения не может быть пустым")
	assert.NotEmpty(t, cfg.App.Version, "Версия приложения не может быть пустым")

	assert.NotEmpty(t, cfg.GRPC.Port, "Порт gRPC не может быть пустым")
	assert.NotEmpty(t, cfg.GRPC.Timeout, "Таймаут gRPC не может быть пустым")

	assert.NotEmpty(t, cfg.Log.Level, "Уровень логирования не может быть пустым")

	assert.NotEmpty(t, cfg.Token.Secret, "JWT Secret не может быть пустым")
	assert.NotEmpty(t, cfg.Token.AccessTTL, "JWT SAccessTTL не может быть пустым")
	assert.NotEmpty(t, cfg.Token.RefreshTTL, "JWT RefreshTTL не может быть пустым")

	assert.NotEmpty(t, cfg.MigrationsPath, "Путь к миграциям не может быть пустым")

	assert.NotEmpty(t, cfg.PG.Host, "PG Host не может быть пустым")
	assert.NotEmpty(t, cfg.PG.Port, "PG Port не может быть пустым")
	assert.NotEmpty(t, cfg.PG.User, "PG User не может быть пустым")
	assert.NotEmpty(t, cfg.PG.Password, "PG Password не может быть пустым")

	fmt.Printf("%+v\n", cfg)
}
