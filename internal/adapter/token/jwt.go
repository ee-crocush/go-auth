package token

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go-auth/internal/entity"
	"time"
)

var _ JWTToken = (*jwtToken)(nil)

// Ошибки токенов
var (
	ErrMissingSecret       = errors.New("missing JWT_SECRET in config")
	ErrInvalidToken        = errors.New("invalid JWT token")
	ErrAccessTokenExpired  = errors.New("access token expired")
	ErrRefreshTokenExpired = errors.New("refresh token expired")
)

// JWTToken - интерфейс для работы с токенами
type JWTToken interface {
	// GenerateAccessToken - генерация access токена
	GenerateAccessToken(user *entity.User) (string, error)
	// GenerateRefreshToken - генерация refresh токена
	GenerateRefreshToken(user *entity.User) (string, error)
	// ValidateToken - валидация токена
	ValidateToken(token string) (bool, error)
	// RefreshAccessToken - обновление access токена
	RefreshAccessToken(refreshToken string) (string, error)
}

// service - реализация интерфейса Service
type jwtToken struct {
	secret     string
	accessTTL  time.Duration
	refreshTTL time.Duration
}

// New - конструктор создает новый экземпляр Service
func New(secret string, accessTTL, refreshTTL time.Duration) (JWTToken, error) {
	if secret == "" {
		return nil, ErrMissingSecret
	}
	return &jwtToken{
		secret:     secret,
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
	}, nil
}

// GenerateAccessToken - генерация access токена
func (s *jwtToken) GenerateAccessToken(user *entity.User) (string, error) {
	return s.generateToken(user.ID, s.accessTTL)
}

// GenerateRefreshToken - генерация refresh токена
func (s *jwtToken) GenerateRefreshToken(user *entity.User) (string, error) {
	return s.generateToken(user.ID, s.refreshTTL)
}

// ValidateToken - валидация токена
func (s *jwtToken) ValidateToken(token string) (bool, error) {
	claims, err := s.parseToken(token)
	if err != nil {
		return false, err
	}

	// Проверяем срок действия токена
	if claims.ExpiresAt.Before(time.Now()) {
		return false, ErrAccessTokenExpired
	}

	return true, nil
}

// RefreshAccessToken - обновление access токена
func (s *jwtToken) RefreshAccessToken(refreshToken string) (string, error) {
	claims, err := s.parseToken(refreshToken)
	if err != nil {
		return "", fmt.Errorf("failed to parse refresh token: %w", err)
	}

	// Проверяем срок действия refresh токена
	if claims.ExpiresAt.Before(time.Now()) {
		return "", ErrRefreshTokenExpired
	}

	// Генерируем новый access токен
	return s.generateToken(claims.Subject, s.accessTTL)
}

// generateToken - вспомогательный метод для генерации токена
func (s *jwtToken) generateToken(userID string, ttl time.Duration) (string, error) {
	claims := &jwt.RegisteredClaims{
		Issuer:    userID,
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secret))
}

// parseToken - парсинг токена и валидация
func (s *jwtToken) parseToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}
	if claims.Subject == "" {
		return nil, ErrInvalidToken
	}
	return claims, nil
}
