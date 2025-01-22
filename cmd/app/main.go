package main

import (
	"flag"
	"go-auth/config"
	app "go-auth/internal/app"
	"log"
)

func main() {
	// Определяем флаг
	devMode := flag.Bool("dev", false, "Run server in development mode")
	flag.Parse()

	// Загружаем конфигурацию
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Инициализируем логгер
	//appLogger := logger.New(*devMode)
	//defer appLogger.Sync()
	//// Настраиваем gRPC логгер
	//appLogger.ReplaceGrpcLogger()

	//	Запускаем
	app.Run(cfg, *devMode)
}
