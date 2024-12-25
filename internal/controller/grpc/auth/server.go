package grpcauth

import (
	"context"
	pb "cyberball-auth/gen/auth"
	"fmt"
)

var _ pb.AuthServer = (*AuthServer)(nil)

// AuthServer - структура для обработки RPC-методов, реализующая интерфейс pb.AuthServer
type AuthServer struct {
	pb.UnimplementedAuthServer
}

// Register - регистрация нового пользователя
func (h *AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// TODO - тут будет регистрация нового пользователя (не логика)
	//userID, err := h.auth.Register(req.Email, req.Password)
	//
	//if err != nil {
	//	return nil, err
	//}

	// Для примера пока просто выводим в консоль, что все работает
	fmt.Printf("Register Request - Email: %s, Password: %s\n", req.Email, req.Password)
	return &pb.RegisterResponse{UserId: "123"}, nil
}

// Login - авторизация пользователя
func (h *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// TODO - тут будет авторизация пользователя (не логика)
	//accessToken, refreshToken, err := h.auth.Login(req.Email, req.Password)
	//if err != nil {
	//	return nil, err
	//}
	// Для примера пока просто выводим в консоль, что все работает
	fmt.Printf("Login Request - Email: %s, Password: %s\n", req.Email, req.Password)
	return &pb.LoginResponse{AccessToken: "123", RefreshToken: "456"}, nil
}

// RefreshToken - обновление токена
func (h *AuthServer) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	// TODO - тут будет обновление токена (не логика)
	//accessToken, err := h.auth.RefreshToken(req.RefreshToken)
	//if err != nil {
	//	return nil, err
	//}
	// Для примера пока просто выводим в консоль, что все работает
	fmt.Printf("RefreshToken Request - RefreshToken: %s\n", req.RefreshToken)
	return &pb.RefreshTokenResponse{AccessToken: "123"}, nil
}

// ValidateToken - проверка токена
func (h *AuthServer) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	// TODO - тут будет проверка токена (не логика)
	//valid, err := h.auth.ValidateToken(req.AccessToken)
	//if err != nil {
	//	return nil, err
	//}
	// Для примера пока просто выводим в консоль, что все работает
	fmt.Printf("ValidateToken Request - AccessToken: %s\n", req.AccessToken)
	return &pb.ValidateTokenResponse{Valid: true}, nil
}

// Logout - выход из системы
func (h *AuthServer) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	// TODO - тут будет выход из системы (не логика)
	//err := h.auth.Logout(req.AccessToken)
	//if err != nil {
	//	return nil, err
	//}
	// Для примера пока просто выводим в консоль, что все работает
	fmt.Printf("Logout Request - AccessToken: %s\n", req.AccessToken)
	return &pb.LogoutResponse{}, nil
}
