package usecase

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"

	"github.com/stretchr/testify/assert"

	"github.com/oniharnantyo/golang-backend-example/domain"

	"github.com/stretchr/testify/mock"

	auth_repository_mock "github.com/oniharnantyo/golang-backend-example/services/auth/repository/mock"
)

func TestNewAuthUseCase(t *testing.T) {
	ctx := context.Background()

	authRepo := new(auth_repository_mock.AuthMockRepository)

	account := domain.Account{
		AccountNumber:  1,
		CustomerNumber: 1,
		Balance:        1000,
		Email:          "mail@gmail.com",
		Password:       "secret",
	}

	t.Run("Success", func(t *testing.T) {
		authRepo.On("Set", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("time.Duration")).Return(nil).Once()
		authRepo.On("Set", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("time.Duration")).Return(nil).Once()

		authUseCase := NewAuthUseCase(authRepo, "secret", 10, "refreshsecret", 15)

		token, err := authUseCase.CreateAuth(ctx, account)
		assert.NoError(t, err)
		assert.NotNil(t, token)

		authRepo.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		authRepo.On("Set", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("time.Duration")).Return(redis.Nil)

		authUseCase := NewAuthUseCase(authRepo, "secret", 10, "refreshsecret", 15)

		token, err := authUseCase.CreateAuth(ctx, account)
		assert.Error(t, err)
		assert.Equal(t, domain.Auth{}, token)

		authRepo.AssertExpectations(t)
	})
}
