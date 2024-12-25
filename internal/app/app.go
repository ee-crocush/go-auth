// Package app configures and runs application.
package app

import (
	"cyberball-auth/config"
	pb "cyberball-auth/gen/auth"
	authgrpc "cyberball-auth/internal/controller/grpc/auth"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// Run - запускает приложение
func Run(cfg *config.Config, devMode bool) {
	// Создаем gRPC-сервер
	grpcServer := grpc.NewServer()
	// Создаем сервис Auth
	authService := &authgrpc.AuthServer{}
	// Регистрируем сервис Auth
	pb.RegisterAuthServer(grpcServer, authService)
	// Включаем рефлексию только в режиме разработки
	if devMode {
		reflection.Register(grpcServer)
		log.Println("Development mode: gRPC reflection enabled")
	} else {
		log.Println("Production mode: gRPC reflection disabled")
	}

	// Слушаем порт gRPC
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))

	if err != nil {
		log.Fatalf("Failed to listen +on port %d: %v", cfg.GRPC.Port, err)
	}

	log.Printf("Starting gRPC server on port %d\n", cfg.GRPC.Port)

	// Запускаем gRPC-сервер
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
