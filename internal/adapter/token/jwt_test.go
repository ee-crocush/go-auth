package token

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go-auth/internal/entity"
	"testing"
	"time"
)

func TestGenerateAccessToken(t *testing.T) {
	// Параметры для создания
	svc, err := baseToken()
	require.NoError(t, err)

	//	Тестовый пользователь
	user := &entity.User{ID: "testid"}
	token, err := svc.GenerateAccessToken(user)
	require.NoError(t, err)

	assert.NotEmpty(t, token)
}

func TestGenerateRefreshToken(t *testing.T) {
	// Параметры для создания
	svc, err := baseToken()
	require.NoError(t, err)

	//	Тестовый пользователь
	user := &entity.User{ID: "testid"}
	token, err := svc.GenerateRefreshToken(user)
	require.NoError(t, err)

	assert.NotEmpty(t, token)
}

func TestValidateToken(t *testing.T) {
	// Параметры для создания
	svc, err := baseToken()
	require.NoError(t, err)

	//	Тестовый пользователь
	user := &entity.User{ID: "testid"}
	token, err := svc.GenerateAccessToken(user)
	require.NoError(t, err)

	valid, err := svc.ValidateToken(token)
	require.NoError(t, err)
	assert.True(t, valid)
}

func TestValidateTokenExpired(t *testing.T) {
	// Параметры для создания
	svc, err := baseExpiredToken()
	require.NoError(t, err)

	//	Тестовый пользователь
	user := &entity.User{ID: "testid"}
	token, err := svc.GenerateAccessToken(user)
	require.NoError(t, err)
	// Ожидаем просрочки токена
	time.Sleep(2 * time.Second)

	valid, err := svc.ValidateToken(token)
	assert.Error(t, err)
	assert.False(t, valid)
}

func TestRefreshToken(t *testing.T) {
	// Параметры для создания
	svc, err := baseToken()
	require.NoError(t, err)

	//	Тестовый пользователь
	user := &entity.User{ID: "testid"}
	token, err := svc.GenerateRefreshToken(user)
	require.NoError(t, err)

	newToken, err := svc.RefreshAccessToken(token)
	require.NoError(t, err)
	assert.NotEmpty(t, newToken)
}

func TestRefreshTokenExpired(t *testing.T) {
	// Параметры для создания
	svc, err := baseExpiredToken()
	require.NoError(t, err)

	//	Тестовый пользователь
	user := &entity.User{ID: "testid"}
	token, err := svc.GenerateRefreshToken(user)
	require.NoError(t, err)

	// Ожидаем просрочки токена
	time.Sleep(2 * time.Second)

	newToken, err := svc.RefreshAccessToken(token)
	valid, err := svc.ValidateToken(newToken)
	assert.Error(t, err)
	assert.False(t, valid)
}

func baseToken() (JWTToken, error) {
	secret := "testsecret"
	accessTTL := 10 * time.Minute
	refreshTTL := 24 * time.Hour
	return New(secret, accessTTL, refreshTTL)
}

func baseExpiredToken() (JWTToken, error) {
	secret := "testsecret"
	accessTTL := -1 * time.Minute
	refreshTTL := 24 * time.Hour
	return New(secret, accessTTL, refreshTTL)
}
