package auth_usecase_mock

import (
	"context"

	"github.com/oniharnantyo/golang-backend-example/domain"
	"github.com/stretchr/testify/mock"
)

type AuthMockUseCase struct {
	mock.Mock
}

func (a AuthMockUseCase) CreateAuth(ctx context.Context, account domain.Account) (domain.Auth, error) {
	args := a.Called(ctx, account)

	return args.Get(0).(domain.Auth), args.Error(1)
}
