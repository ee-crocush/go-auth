package usecase

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go-auth/internal/entity"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestRegister(t *testing.T) {
	mockUser := new(MockUserRepository)
	mockBlacklist := new(MockBlacklistRepository)
	mockJWT := new(MockTokenService)

	repo := NewAuthUseCase(mockUser, mockBlacklist, mockJWT)

	t.Run("success", func(t *testing.T) {
		email := "test@example.com"
		password := "password123"

		// Мок репозитория возвращает nil (пользователя еще нет)
		mockUser.On("GetUserByEmail", mock.Anything, email).Return(nil, nil)

		// Переопределяем ожидание создания пользователя в репозитории
		mockUser.On("CreateUser", mock.Anything, mock.Anything).Return("new-user-id", nil)

		// Проверяем результат
		userID, err := repo.Register(email, password)

		require.NoError(t, err)
		require.NotEmpty(t, userID)

		mockUser.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockBlackListRepo := new(MockBlacklistRepository)
	mockTokenService := new(MockTokenService)

	authUseCase := NewAuthUseCase(mockUserRepo, mockBlackListRepo, mockTokenService)
	email := "test@example.com"
	password := "password123"

	t.Run("success", func(t *testing.T) {
		// Генерация хэша пароля
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		require.NoError(t, err)

		mockUser := &entity.User{
			ID:       "user-id",
			Email:    email,
			Password: string(hashedPassword), // Используем свежесгенерированный хэш
			IsActive: true,
		}

		mockUserRepo.On("GetUserByEmail", mock.Anything, email).Return(mockUser, nil)
		mockTokenService.On("GenerateAccessToken", mockUser).Return("access_token", nil)
		mockTokenService.On("GenerateRefreshToken", mockUser).Return("refresh_token", nil)

		accessToken, refreshToken, err := authUseCase.Login(email, password)

		require.NoError(t, err)
		require.Equal(t, "access_token", accessToken)
		require.Equal(t, "refresh_token", refreshToken)

		mockUserRepo.AssertExpectations(t)
		mockTokenService.AssertExpectations(t)
	})
	t.Run("invalid credentials", func(t *testing.T) {
		// Генерация хэша пароля
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		require.NoError(t, err)

		mockUser := &entity.User{
			ID:       "user-id",
			Email:    email,
			Password: string(hashedPassword), // Используем свежесгенерированный хэш
			IsActive: true,
		}

		mockUserRepo.On("GetUserByEmail", mock.Anything, email).Return(mockUser, nil)

		_, _, err = authUseCase.Login(email, "dfgdfgd")
		require.ErrorIs(t, err, ErrInvalidCredentials)

		mockUserRepo.AssertExpectations(t)
	})
}

func TestRefreshToken(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockBlacklistRepo := new(MockBlacklistRepository)
	mockTokenService := new(MockTokenService)

	authUseCase := NewAuthUseCase(mockUserRepo, mockBlacklistRepo, mockTokenService)

	t.Run("success", func(t *testing.T) {
		refreshToken := "valid-refresh-token"
		newAccessToken := "new-access-token"

		// Настраиваем моки
		mockBlacklistRepo.On("IsTokenBlacklisted", mock.Anything, refreshToken).Return(false, nil)
		mockTokenService.On("ValidateToken", refreshToken).Return(true, nil)
		mockTokenService.On("RefreshAccessToken", refreshToken).Return(newAccessToken, nil)

		// Тестируем
		token, err := authUseCase.RefreshToken(refreshToken)
		require.NoError(t, err)
		require.Equal(t, newAccessToken, token)

		// Проверяем вызовы
		mockBlacklistRepo.AssertExpectations(t)
		mockTokenService.AssertExpectations(t)
	})

	t.Run("invalid token", func(t *testing.T) {
		refreshToken := "invalid-refresh-token"

		mockBlacklistRepo.On("IsTokenBlacklisted", mock.Anything, refreshToken).Return(false, nil)
		mockTokenService.On("ValidateToken", refreshToken).Return(false, errors.New("invalid token"))

		token, err := authUseCase.RefreshToken(refreshToken)
		require.Error(t, err)
		require.Empty(t, token)
		require.Contains(t, err.Error(), "invalid refresh token")

		mockBlacklistRepo.AssertExpectations(t)
		mockTokenService.AssertExpectations(t)
	})

	t.Run("blacklisted token", func(t *testing.T) {
		refreshToken := "blacklisted-token"

		mockBlacklistRepo.On("IsTokenBlacklisted", mock.Anything, refreshToken).Return(true, nil)

		token, err := authUseCase.RefreshToken(refreshToken)
		require.Error(t, err)
		require.Empty(t, token)
		require.True(t, errors.Is(err, ErrBlackListed), "expected error to be ErrBlackListed")

		mockBlacklistRepo.AssertExpectations(t)
	})
}

func TestValidateToken(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockBlacklistRepo := new(MockBlacklistRepository)
	mockTokenService := new(MockTokenService)

	authUseCase := NewAuthUseCase(mockUserRepo, mockBlacklistRepo, mockTokenService)

	t.Run("success", func(t *testing.T) {
		token := "valid-access-token"

		mockTokenService.On("ValidateToken", token).Return(true, nil)
		mockBlacklistRepo.On("IsTokenBlacklisted", mock.Anything, token).Return(false, nil)

		valid, err := authUseCase.ValidateToken(token)
		require.NoError(t, err)
		require.True(t, valid)

		mockTokenService.AssertExpectations(t)
		mockBlacklistRepo.AssertExpectations(t)
	})

	t.Run("invalid token", func(t *testing.T) {
		token := "invalid-access-token"

		mockTokenService.On("ValidateToken", token).Return(false, nil)
		mockBlacklistRepo.On("IsTokenBlacklisted", mock.Anything, token).Return(false, nil)

		valid, err := authUseCase.ValidateToken(token)
		require.NoError(t, err)
		require.False(t, valid)

		mockTokenService.AssertExpectations(t)
	})
}

func TestLogout(t *testing.T) {
	// Устанавливаем мокированные объекты
	mockBlacklistRepo := new(MockBlacklistRepository)

	authUseCase := &auth{
		blacklist: mockBlacklistRepo,
	}

	t.Run("success", func(t *testing.T) {
		accessToken := "valid-access-token"

		mockBlacklistRepo.On("AddToBlacklist", mock.Anything, accessToken).Return(nil)

		err := authUseCase.Logout(accessToken)
		require.NoError(t, err)

		mockBlacklistRepo.AssertExpectations(t)
	})
}
