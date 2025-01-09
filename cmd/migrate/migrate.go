package main

import (
	"cyberball-auth/config"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Путь к миграциям и строка подключения к бд
	migrationsPath := fmt.Sprintf("file://%s", cfg.Migrations.Path)
	dbURL := cfg.PG.MigrationsURL()

	// Создаем объект миграции
	m, err := migrate.New(migrationsPath, dbURL)
	if err != nil {
		log.Fatalf("Failed to create migrations: %v", err)
	}

	// Выполняем команду, переданную как аргумент (up, down)
	if len(os.Args) < 2 {
		log.Fatal("Usage: migrate <command>\nAvailable commands: up, down")
	}
	command := os.Args[1]
	switch command {
	case "up":
		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("Failed to apply migrations: %v", err)
		}
		log.Printf("Migrations applied successfully!")
	case "down":
		if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("Failed to rollback back migrations: %v", err)
		}
		log.Printf("Migrations rolled back successfully!")

	default:
		log.Fatalf("Unknown command: %s\nAvailable commands: up, down", command)
	}
}
