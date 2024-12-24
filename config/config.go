package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"time"
)

type (
	// Config - структура конфига проекта
	Config struct {
		App            AppConfig   `yaml:"app"`    // Инфа о приложении
		GRPC           GRPCConfig  `yaml:"grpc"`   // Инфа по gRPC сервера
		Token          TokenConfig `yaml:"token"`  // Инфа по токену
		Log            LogConfig   `yaml:"logger"` // Уровень логгирования
		PG             PGConfig    // Данные по Postgres
		MigrationsPath string      `env:"MIGRATIONS_PATH"` // путь к миграциям
	}
	AppConfig struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}

	GRPCConfig struct {
		Port    int `yaml:"port"`
		Timeout int `yaml:"timeout"`
	}

	LogConfig struct {
		Level string `yaml:"level"`
	}

	PGConfig struct {
		User     string `env:"PG_USER"`
		Password string `env:"PG_PASSWORD"`
		Host     string `env:"PG_HOST"`
		Port     int    `env:"PG_PORT"`
	}

	TokenConfig struct {
		Secret     string        `env:"TOKEN_SECRET"`
		AccessTTL  time.Duration `yaml:"accessTTL"`
		RefreshTTL time.Duration `yaml:"refreshTTL"`
	}
)

func NewConfig() (*Config, error) {

	// Загружаем переменные окружения из .env файла
	if err := godotenv.Load("../.env"); err != nil {
		return nil, err
	}

	// Создаем конфигурацию
	cfg := &Config{}

	err := cleanenv.ReadConfig("../config/config.yaml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	// Загружаем конфигурацию с использованием cleanenv
	if err := cleanenv.ReadEnv(cfg); err != nil {
		log.Println("Error loading environment variables:", err)
		return nil, err
	}

	return cfg, nil
}
