package domain

import (
	"context"

	"github.com/oniharnantyo/golang-backend-example/util"
)

type Customer struct {
	CustomerNumber int    `json:"account_number"`
	Name           string `json:"name"`
}

type CustomerListParam struct {
	util.Filter
}

type (
	CustomerUseCase interface {
		List(ctx context.Context, param CustomerListParam) ([]Customer, error)
		GetByCustomerNumber(ctx context.Context, accountNumber int) (Customer, error)
		Store(ctx context.Context, a *Customer) error
		Update(ctx context.Context, a *Customer) error
		Delete(ctx context.Context, a *Customer) error
	}

	CustomerRepository interface {
		List(ctx context.Context, param CustomerListParam) ([]Customer, error)
		GetByCustomerNumber(ctx context.Context, customerNumber int) (Customer, error)
		Store(ctx context.Context, a *Customer) error
		Update(ctx context.Context, a *Customer) error
		Delete(ctx context.Context, a *Customer) error
	}
)
