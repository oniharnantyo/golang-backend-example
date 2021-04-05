package domain

import (
	"context"
	"linkaja-test/util"
)

type (
	Account struct {
		AccountNumber  int `json:"account_number"`
		CustomerNumber int `json:"customer_number"`
		Balance        int `json:"balance"`
	}

	AccountListParam struct {
		util.Filter
	}

	DetailByAccountNumberResponse struct {
		AccountNumber int    `json:"account_number"`
		CustomerName  string `json:"customer_name"`
		Balance       int    `json:"balance"`
	}

	TransferParam struct {
		ToAccountNumber string `json:"to_account_number"`
		Amount          int    `json:"amount"`
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
	}

	AccountRepository interface {
		List(ctx context.Context, param AccountListParam) ([]Account, error)
		GetByAccountNumber(ctx context.Context, accountNumber int) (Account, error)
		Store(ctx context.Context, a *Account) error
		Update(ctx context.Context, a *Account) error
		Delete(ctx context.Context, a *Account) error
	}
)
