package repository_customer_mock

import (
	"context"
	"linkaja-test/domain"

	"github.com/stretchr/testify/mock"
)

type AccountMockRepository struct {
	mock.Mock
}

func (c *AccountMockRepository) List(ctx context.Context, param domain.AccountListParam) ([]domain.Account, error) {
	args := c.Called(ctx, param)
	result := args.Get(0)

	return result.([]domain.Account), args.Error(1)
}

func (c *AccountMockRepository) GetByAccountNumber(ctx context.Context, customerNumber int) (domain.Account, error) {
	args := c.Called(ctx, customerNumber)
	result := args.Get(0)

	return result.(domain.Account), args.Error(1)
}

func (c *AccountMockRepository) Store(ctx context.Context, a *domain.Account) error {
	args := c.Called(ctx, a)

	return args.Error(0)
}

func (c *AccountMockRepository) Update(ctx context.Context, a *domain.Account) error {
	args := c.Called(ctx, a)

	return args.Error(0)
}

func (c *AccountMockRepository) Delete(ctx context.Context, a *domain.Account) error {
	args := c.Called(ctx, a)

	return args.Error(0)
}
