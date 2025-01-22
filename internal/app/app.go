// Package app configures and runs application.
package app

import (
	"context"
	"fmt"
	"go-auth/config"
	"go-auth/gen/auth"
	"go-auth/internal/adapter/database"
	"go-auth/internal/adapter/token"
	"go-auth/internal/controller/grpc/auth"
	"go-auth/internal/repo/blacklist"
	"go-auth/internal/repo/user"
	"go-auth/internal/usecase/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"log"
	"net"
)

// Run - запускает приложение
func Run(cfg *config.Config, devMode bool) {
	// Инициализация дефолтного логгера
	logger := log.Default()
	// Подключение к базе данных
	dbpool, err := database.New(context.Background(), *cfg)
	if err != nil {
		logger.Fatalf("Unable to create connection pool: %v", err)
	}
	defer dbpool.Close()
	logger.Printf("Database connection established")

	// Создаем репозитории
	userRepo := user.NewRepository(dbpool)
	blacklistRepo := blacklist.NewRepository(dbpool)

	// Создаем сервис работы с токенами
	tokenSvc, err := token.New(cfg.Token.Secret, cfg.Token.AccessTTL, cfg.Token.RefreshTTL)
	if err != nil {
		logger.Fatalf("Failed to initialize token service: %v", err)
	}

	// Создаем слой usecase
	authUseCase := usecase.NewAuthUseCase(userRepo, blacklistRepo, tokenSvc)

	// Создаем gRPC-сервер
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpcLogStreamInterceptor),
		grpc.UnaryInterceptor(grpcLogUnaryInterceptor),
	)

	// Создаем и регистрируем gRPC-сервис Auth
	authController := grpcauth.NewAuthServer(authUseCase)
	auth.RegisterAuthServer(grpcServer, authController)

	// Слушаем порт gRPC
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
	if err != nil {
		logger.Fatalf("Failed to listen on port %d: %v", cfg.GRPC.Port, err)
	}

	logger.Printf("Starting gRPC server on port %d\n", cfg.GRPC.Port)
	// Запускаем gRPC-сервер
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatalf("Failed to serve gRPC server: %v", err)
	}
}

// Интерсепторы для логирования
func grpcLogStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	logger := log.Default()
	logger.Printf("gRPC Stream called: %s from %s", info.FullMethod, ss.Context().Value("peer").(*peer.Peer).Addr.String())
	return handler(srv, ss)
}

func grpcLogUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	logger := log.Default()
	peerInfo, _ := peer.FromContext(ctx)
	//if !ok {
	//	logger.Printf("gRPC Unary called: %s from UNKNOWN", info.FullMethod)
	//} else {
	//	//logger.Printf("gRPC Unary called: %s from %s", info.FullMethod, peerInfo.Addr.String())
	//}

	resp, err := handler(ctx, req)

	if err != nil {
		logger.Printf("gRPC Unary response error: %s, method: %s, error: %v", peerInfo.Addr.String(), info.FullMethod, err)
	} else {
		logger.Printf("gRPC Unary response: %s, method: %s, response: %v", peerInfo.Addr.String(), info.FullMethod, resp)
	}

	return resp, err
}
