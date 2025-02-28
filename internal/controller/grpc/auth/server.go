package grpcauth

import (
	"context"
	pb "go-auth/gen/auth"
	"go-auth/internal/usecase/auth"
	"go-auth/pkg/validator"
)

var _ pb.AuthServer = (*AuthServer)(nil)

// AuthServer - структура для обработки RPC-методов, реализующая интерфейс pb.AuthServer
type AuthServer struct {
	pb.UnimplementedAuthServer
	auth usecase.AuthUseCase
}

// RequestValidation - структура для валидации запроса
type RequestValidation struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6,max=128"`
}

// NewAuthServer - конструктор для AuthServer
func NewAuthServer(auth usecase.AuthUseCase) *AuthServer {
	return &AuthServer{auth: auth}
}

// Register - регистрация нового пользователя
func (s *AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// Подготовка структуры для валидации
	validateReq := &RequestValidation{
		Email:    req.Email,
		Password: req.Password,
	}
	if err := validator.ValidateRequest(validateReq); err != nil {
		return nil, err
	}

	userID, err := s.auth.Register(req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResponse{UserId: userID}, nil
}

// Login - авторизация пользователя
func (s *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// Подготовка структуры для валидации
	validateReq := &RequestValidation{
		Email:    req.Email,
		Password: req.Password,
	}
	if err := validator.ValidateRequest(validateReq); err != nil {
		return nil, err
	}

	accessToken, refreshToken, err := s.auth.Login(req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// RefreshToken - обновление токена
func (s *AuthServer) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	accessToken, err := s.auth.RefreshToken(req.RefreshToken)
	if err != nil {
		return nil, err
	}
	return &pb.RefreshTokenResponse{AccessToken: accessToken}, nil
}

// ValidateToken - проверка токена
func (s *AuthServer) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	valid, err := s.auth.ValidateToken(req.AccessToken)
	if err != nil {
		return nil, err
	}
	return &pb.ValidateTokenResponse{Valid: valid}, nil
}

// Logout - выход из системы
func (s *AuthServer) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	err := s.auth.Logout(req.AccessToken)
	if err != nil {
		return nil, err
	}
	return &pb.LogoutResponse{}, nil
}
