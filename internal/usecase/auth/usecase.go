package usecase

import (
	"context"
	"cyberball-auth/internal/adapter/token"
	"cyberball-auth/internal/entity"
	"cyberball-auth/internal/repo/blacklist"
	"cyberball-auth/internal/repo/user"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrExistingUser       = errors.New("email already in use")
	ErrUserNotActive      = errors.New("user account is not active")
	ErrBlackListed        = errors.New("access token is blacklisted")
	ErrMinLengthPswd      = errors.New("password length must be between 6 and 128 characters")
)

var _ AuthUseCase = (*auth)(nil)

// AuthUseCase - интерфейс для аутентификации
type AuthUseCase interface {
	// Register - регистрация нового пользователя
	Register(email string, password string) (userId string, err error)
	// Login - авторизация пользователя
	Login(email string, password string) (accessToken, refreshToken string, err error)
	// RefreshToken - обновление токена
	RefreshToken(refreshToken string) (accessToken string, err error)
	// ValidateToken - проверка токена
	ValidateToken(accessToken string) (valid bool, err error)
	// Logout - выход из системы
	Logout(accessToken string) (err error)
}

// Auth - структура для аутентификации, пока пустая
type auth struct {
	userRepo     user.Repository
	blacklist    blacklist.Repository
	tokenService token.JWTToken
}

// NewAuthUseCase - конструктор для auth
func NewAuthUseCase(userRepo user.Repository, blacklist blacklist.Repository, tokenSvc token.JWTToken) AuthUseCase {
	return &auth{
		userRepo:     userRepo,
		blacklist:    blacklist,
		tokenService: tokenSvc,
	}
}

// Register - регистрация нового пользователя
func (uc *auth) Register(email string, password string) (string, error) { // Проверка сложности и длины пароля
	if len(password) < 6 || len(password) > 128 {
		return "", ErrMinLengthPswd
	}

	// Проверяем, что пользователя с таким email не существует
	existingUser, err := uc.userRepo.GetUserByEmail(context.Background(), email)
	if err == nil && existingUser != nil {
		return "", ErrExistingUser
	}
	// Хэшируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// Генерация ID пользователя
	userID := uuid.New().String()
	// Создаем нового пользователя
	newUser := &entity.User{
		ID:        userID,
		Email:     email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		IsActive:  true, // Пользователь активен по умолчанию
	}
	return uc.userRepo.CreateUser(context.Background(), newUser)
}

// Login - авторизация пользователя
func (uc *auth) Login(email string, password string) (string, string, error) {
	// Получаем пользователя
	curUser, err := uc.userRepo.GetUserByEmail(context.Background(), email)
	if err != nil || curUser == nil {
		return "", "", ErrInvalidCredentials
	}

	// Проверяем пароль
	if err := bcrypt.CompareHashAndPassword([]byte(curUser.Password), []byte(password)); err != nil {
		return "", "", ErrInvalidCredentials
	}

	// Проверяем, активен ли пользователь
	if !curUser.IsActive {
		return "", "", ErrUserNotActive
	}

	//Генерируем токены
	accessToken, err := uc.tokenService.GenerateAccessToken(curUser)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := uc.tokenService.GenerateRefreshToken(curUser)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// RefreshToken - обновление токена
func (uc *auth) RefreshToken(refreshToken string) (string, error) {
	// Проверяем валидность refreshToken
	isValid, err := uc.checkToken(refreshToken)
	if !isValid || err != nil {
		return "", fmt.Errorf("invalid refresh token: %w", err)
	}

	// Генерация нового access токена
	newAccessToken, err := uc.tokenService.RefreshAccessToken(refreshToken)
	if err != nil {
		return "", fmt.Errorf("failed to generate access token: %w", err)
	}
	return newAccessToken, nil
}

// ValidateToken - проверка токена
func (uc *auth) ValidateToken(accessToken string) (bool, error) {
	return uc.checkToken(accessToken)
}

// checkToken - проверка токена
func (uc *auth) checkToken(token string) (bool, error) {
	// Проверяем, находится ли токен в blacklist
	isBlacklisted, err := uc.blacklist.IsTokenBlacklisted(context.Background(), token)
	if err != nil {
		return false, fmt.Errorf("failed to check blacklist: %w", err)
	}
	if isBlacklisted {
		return false, ErrBlackListed
	}

	// Проверяем валидность токена
	isValid, err := uc.tokenService.ValidateToken(token)
	if err != nil {
		return false, err
	}
	return isValid, nil
}

// Logout - выход из системы
func (uc *auth) Logout(accessToken string) error {
	// Добавляем токен в blacklist
	err := uc.blacklist.AddToBlacklist(context.Background(), accessToken)
	if err != nil {
		return fmt.Errorf("failed to add token to blacklist: %w", err)
	}
	return nil
}
