package repository_customer_mock

import (
	"context"

	"github.com/oniharnantyo/golang-backend-example/domain"

	"github.com/stretchr/testify/mock"
)

type CustomerMockRepository struct {
	mock.Mock
}

func (c *CustomerMockRepository) List(ctx context.Context, param domain.CustomerListParam) ([]domain.Customer, error) {
	args := c.Called(ctx, param)
	result := args.Get(0)

	return result.([]domain.Customer), args.Error(1)
}

func (c *CustomerMockRepository) GetByCustomerNumber(ctx context.Context, customerNumber int) (domain.Customer, error) {
	args := c.Called(ctx, customerNumber)
	result := args.Get(0)

	return result.(domain.Customer), args.Error(1)
}

func (c *CustomerMockRepository) Store(ctx context.Context, a *domain.Customer) error {
	args := c.Called(ctx, a)

	return args.Error(0)
}

func (c *CustomerMockRepository) Update(ctx context.Context, a *domain.Customer) error {
	args := c.Called(ctx, a)

	return args.Error(0)
}

func (c *CustomerMockRepository) Delete(ctx context.Context, a *domain.Customer) error {
	args := c.Called(ctx, a)

	return args.Error(0)
}
