package account_usecase_mock

import (
	"context"

	"github.com/oniharnantyo/golang-backend-example/domain"

	"github.com/stretchr/testify/mock"
)

type AccountMockUseCase struct {
	mock.Mock
}

func (c *AccountMockUseCase) List(ctx context.Context, param domain.AccountListParam) ([]domain.Account, error) {
	args := c.Called(ctx, param)
	result := args.Get(0)

	return result.([]domain.Account), args.Error(1)
}

func (c *AccountMockUseCase) GetByAccountNumber(ctx context.Context, customerNumber int) (domain.DetailByAccountNumberResponse, error) {
	args := c.Called(ctx, customerNumber)
	result := args.Get(0)

	return result.(domain.DetailByAccountNumberResponse), args.Error(1)
}

func (c *AccountMockUseCase) Store(ctx context.Context, a *domain.Account) error {
	args := c.Called(ctx, a)

	return args.Error(0)
}

func (c *AccountMockUseCase) Update(ctx context.Context, a *domain.Account) error {
	args := c.Called(ctx, a)

	return args.Error(0)
}

func (c *AccountMockUseCase) Delete(ctx context.Context, a *domain.Account) error {
	args := c.Called(ctx, a)

	return args.Error(0)
}

func (c *AccountMockUseCase) Transfer(ctx context.Context, fromAccountNumber int, a domain.TransferParam) error {
	args := c.Called(ctx, fromAccountNumber, a)

	return args.Error(0)
}

func (c *AccountMockUseCase) Login(ctx context.Context, param domain.AccountLoginParam) (domain.LoginResponse, error) {
	args := c.Called(ctx, param)
	result := args.Get(0)

	return result.(domain.LoginResponse), args.Error(1)
}
