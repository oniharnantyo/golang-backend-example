package auth_repository_mock

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

type (
	AuthMockRepository struct {
		mock.Mock
	}
)

func (a *AuthMockRepository) Get(ctx context.Context, id string) (string, error) {
	args := a.Called(ctx, id)

	return args.Get(0).(string), args.Error(1)
}

func (a *AuthMockRepository) Set(ctx context.Context, key, value string, expire time.Duration) error {
	args := a.Called(ctx, key, value, expire)

	return args.Error(0)
}
