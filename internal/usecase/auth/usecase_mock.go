package usecase

import (
	"context"
	"cyberball-auth/internal/entity"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUserById(ctx context.Context, id string) (*entity.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.User), args.Error(1)
}
func (m *MockUserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user *entity.User) (string, error) {
	args := m.Called(ctx, user)
	return args.String(0), args.Error(1)
}

type MockBlacklistRepository struct {
	mock.Mock
}

func (m *MockBlacklistRepository) AddToBlacklist(ctx context.Context, token string) error {
	args := m.Called(ctx, token)
	return args.Error(0)
}

func (m *MockBlacklistRepository) IsTokenBlacklisted(ctx context.Context, token string) (bool, error) {
	args := m.Called(ctx, token)
	return args.Bool(0), args.Error(1)
}

type MockTokenService struct {
	mock.Mock
}

func (m *MockTokenService) GenerateAccessToken(user *entity.User) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}

func (m *MockTokenService) GenerateRefreshToken(user *entity.User) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}

func (m *MockTokenService) RefreshAccessToken(refreshToken string) (string, error) {
	args := m.Called(refreshToken)
	return args.String(0), args.Error(1)
}

func (m *MockTokenService) ValidateToken(token string) (bool, error) {
	args := m.Called(token)
	return args.Bool(0), args.Error(1)
}
