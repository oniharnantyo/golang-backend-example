package domain

import (
	"context"

	"github.com/oniharnantyo/golang-backend-example/util"
)

type (
	Account struct {
		AccountNumber  int    `json:"account_number"`
		CustomerNumber int    `json:"customer_number"`
		Balance        int    `json:"balance"`
		Email          string `json:"email"`
		Password       string `json:"-"`
	}

	AccountListParam struct {
		util.Filter
	}

	AccountLoginParam struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	TransferParam struct {
		ToAccountNumber string `json:"to_account_number"`
		Amount          int    `json:"amount"`
	}

	DetailByAccountNumberResponse struct {
		AccountNumber int    `json:"account_number"`
		CustomerName  string `json:"customer_name"`
		Balance       int    `json:"balance"`
	}

	LoginResponse struct {
		Token string `json:"token"`
	}
)

type (
	AccountUseCase interface {
		List(ctx context.Context, param AccountListParam) ([]Account, error)
		GetByAccountNumber(ctx context.Context, accountNumber int) (DetailByAccountNumberResponse, error)
		Store(ctx context.Context, a *Account) error
		Update(ctx context.Context, a *Account) error
		Delete(ctx context.Context, a *Account) error
		Transfer(ctx context.Context, fromAccountNumber int, param TransferParam) error

		Login(ctx context.Context, param AccountLoginParam) (LoginResponse, error)
	}

	AccountRepository interface {
		List(ctx context.Context, param AccountListParam) ([]Account, error)
		GetByAccountNumber(ctx context.Context, accountNumber int) (Account, error)
		GetByEmail(ctx context.Context, email string) (Account, error)
		Store(ctx context.Context, a *Account) error
		Update(ctx context.Context, a *Account) error
		Delete(ctx context.Context, a *Account) error
	}
)
