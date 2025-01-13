package grpcauth

import (
	"context"
	"testing"

	pb "cyberball-auth/gen/auth"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegister(t *testing.T) {
	mockAuth := new(MockAuth)
	server := NewAuthServer(mockAuth)

	ctx := context.Background()
	req := &pb.RegisterRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	mockAuth.On("Register", req.Email, req.Password).Return("123", nil)

	res, err := server.Register(ctx, req)
	require.NoError(t, err)
	assert.Equal(t, "123", res.UserId)

	mockAuth.AssertExpectations(t)
}

func TestRegisterValidationError(t *testing.T) {
	mockAuth := new(MockAuth)
	server := NewAuthServer(mockAuth)

	ctx := context.Background()
	req := &pb.RegisterRequest{
		Email:    "invalid",
		Password: "123",
	}
	res, err := server.Register(ctx, req)
	require.Error(t, err)
	assert.Nil(t, res)

}

func TestLogin(t *testing.T) {
	mockAuth := new(MockAuth)
	server := NewAuthServer(mockAuth)

	ctx := context.Background()
	req := &pb.LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	mockAuth.On("Login", req.Email, req.Password).Return("access_token", "refresh_token", nil)

	res, err := server.Login(ctx, req)
	require.NoError(t, err)
	assert.Equal(t, "access_token", res.AccessToken)
	assert.Equal(t, "refresh_token", res.RefreshToken)

	mockAuth.AssertExpectations(t)
}

func TestLoginValidationError(t *testing.T) {
	mockAuth := new(MockAuth)
	server := NewAuthServer(mockAuth)

	ctx := context.Background()
	req := &pb.LoginRequest{
		Email:    "invalid",
		Password: "123",
	}

	mockAuth.On("Login", req.Email, req.Password).Return("", "", errors.New("invalid credentials"))

	res, err := server.Login(ctx, req)
	require.Error(t, err)
	assert.Nil(t, res)
}

func TestRefreshToken(t *testing.T) {
	mockAuth := new(MockAuth)
	server := NewAuthServer(mockAuth)

	ctx := context.Background()
	req := &pb.RefreshTokenRequest{
		RefreshToken: "refreshToken12345",
	}

	mockAuth.On("RefreshToken", req.RefreshToken).Return("access_token", nil)

	res, err := server.RefreshToken(ctx, req)
	require.NoError(t, err)
	assert.Equal(t, "access_token", res.AccessToken)

	mockAuth.AssertExpectations(t)
}

func TestValidateToken(t *testing.T) {
	mockAuth := new(MockAuth)
	server := NewAuthServer(mockAuth)

	ctx := context.Background()
	req := &pb.ValidateTokenRequest{
		AccessToken: "accessToken12345",
	}

	mockAuth.On("ValidateToken", req.AccessToken).Return(true, nil)

	res, err := server.ValidateToken(ctx, req)
	require.NoError(t, err)
	assert.True(t, res.Valid)

	mockAuth.AssertExpectations(t)
}

func TestLogout(t *testing.T) {
	mockAuth := new(MockAuth)
	server := NewAuthServer(mockAuth)

	ctx := context.Background()
	req := &pb.LogoutRequest{
		AccessToken: "accessToken12345",
	}

	mockAuth.On("Logout", req.AccessToken).Return(nil)

	res, err := server.Logout(ctx, req)
	require.NoError(t, err)
	assert.Equal(t, &pb.LogoutResponse{}, res)

	mockAuth.AssertExpectations(t)
}
