package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"time"
)

type (
	// Config - структура конфига проекта
	Config struct {
		App        AppConfig        `yaml:"app"`    // Инфа о приложении
		GRPC       GRPCConfig       `yaml:"grpc"`   // Инфа по gRPC сервера
		Token      TokenConfig      `yaml:"token"`  // Инфа по токену
		Log        LogConfig        `yaml:"logger"` // Уровень логгирования
		PG         PGConfig         // Данные по Postgres
		Migrations MigrationsConfig `yaml:"migrations"` // путь к миграциям
	}
	// AppConfig - структура конфига приложения
	AppConfig struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}
	// GRPCConfig - структура конфига gRPC
	GRPCConfig struct {
		Port    int `yaml:"port"`
		Timeout int `yaml:"timeout"`
	}
	// LogConfig - структура конфига логгирования
	LogConfig struct {
		Level string `yaml:"level"`
	}
	// PGConfig - структура конфига базы данных
	PGConfig struct {
		User        string        `env:"PG_USER"`
		Password    string        `env:"PG_PASSWORD"`
		Host        string        `env:"PG_HOST"`
		Port        int           `env:"PG_PORT"`
		DbName      string        `env:"PG_DBNAME"`
		MaxConns    int32         `env:"DB_MAX_CONNS"`
		ConnTimeout time.Duration `env:"DB_CONN_TIMEOUT"`
	}
	// TokenConfig - структура конфига токена
	TokenConfig struct {
		Secret     string        `env:"TOKEN_SECRET"`
		AccessTTL  time.Duration `yaml:"accessTTL"`
		RefreshTTL time.Duration `yaml:"refreshTTL"`
	}
	// MigrationsConfig - структура конфига миграций
	MigrationsConfig struct {
		Path string `yaml:"path"`
	}
)

// URL формирует строку подключения к PostgreSQL
func (p PGConfig) URL() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		p.User,
		p.Password,
		p.Host,
		p.Port,
		p.DbName,
	)
}

// MigrationsURL формирует строку подключения к PostgreSQL для миграции
func (p PGConfig) MigrationsURL() string {
	return fmt.Sprintf("pgx5://%s:%s@%s:%d/%s?sslmode=disable",
		p.User,
		p.Password,
		p.Host,
		p.Port,
		p.DbName,
	)
}

// NewConfig - конструктор для создания Config
func NewConfig() (*Config, error) {
	// Создаем конфигурацию
	cfg := &Config{}

	// Загружаем конфигурацию с использованием cleanenv
	if err := cleanenv.ReadConfig("./config/config.yaml", cfg); err != nil {
		log.Println("Error loading environment variables:", err)
		return nil, err
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		log.Println("Error loading environment variables:", err)
		return nil, err
	}

	return cfg, nil
}
