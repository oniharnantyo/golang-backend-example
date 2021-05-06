package mock

import (
	"context"

	"github.com/oniharnantyo/golang-backend-example/domain"

	"github.com/stretchr/testify/mock"
)

type CustomerMockUseCase struct {
	mock.Mock
}

func (c *CustomerMockUseCase) List(ctx context.Context, param domain.CustomerListParam) ([]domain.Customer, error) {
	args := c.Called(ctx, param)
	result := args.Get(0)

	return result.([]domain.Customer), args.Error(1)
}

func (c *CustomerMockUseCase) GetByCustomerNumber(ctx context.Context, customerNumber int) (domain.Customer, error) {
	args := c.Called(ctx, customerNumber)
	result := args.Get(0)

	return result.(domain.Customer), args.Error(1)
}

func (c *CustomerMockUseCase) Store(ctx context.Context, a *domain.Customer) error {
	args := c.Called(ctx, a)

	return args.Error(1)
}

func (c *CustomerMockUseCase) Update(ctx context.Context, a *domain.Customer) error {
	args := c.Called(ctx, a)

	return args.Error(1)
}

func (c *CustomerMockUseCase) Delete(ctx context.Context, a *domain.Customer) error {
	args := c.Called(ctx, a)

	return args.Error(1)
}
