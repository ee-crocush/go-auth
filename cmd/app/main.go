package main

import (
	"cyberball-auth/config"
	app "cyberball-auth/internal/app"
	"flag"
	"log"
)

func main() {

	// Определяем флаг
	devMode := flag.Bool("dev", false, "Run server in development mode with gRPC reflection enabled")
	flag.Parse()

	// Загружаем конфигурацию
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatal(err)
	}

	//	Запускаем
	app.Run(cfg, *devMode)
}
