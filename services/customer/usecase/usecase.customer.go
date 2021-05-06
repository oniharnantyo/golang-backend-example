package usecase

import (
	"context"
	"database/sql"

	"github.com/oniharnantyo/golang-backend-example/domain"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type customerUseCase struct {
	customerRepository domain.CustomerRepository
	logger             *logrus.Logger
}

func (c customerUseCase) List(ctx context.Context, param domain.CustomerListParam) ([]domain.Customer, error) {
	customers, err := c.customerRepository.List(ctx, param)
	if err != nil {
		c.logger.Errorf("customerUseCase/List/List :%v", err)
		return nil, err
	}

	return customers, nil
}

func (c customerUseCase) GetByCustomerNumber(ctx context.Context, accountNumber int) (domain.Customer, error) {
	customer, err := c.customerRepository.GetByCustomerNumber(ctx, accountNumber)
	if err != nil {
		c.logger.Errorf("customerUseCase/GetByCustomerNumber/GetByCustomerNumber :%v", err)
		if errors.Cause(err) == sql.ErrNoRows {
			return domain.Customer{}, errors.New("Customer not found")
		}
		return domain.Customer{}, err
	}

	return customer, nil
}

func (c customerUseCase) Store(ctx context.Context, a *domain.Customer) error {
	err := c.customerRepository.Store(ctx, a)
	if err != nil {
		c.logger.Errorf("customerUseCase/Store/Store :%v", err)
		return err
	}

	return nil
}

func (c customerUseCase) Update(ctx context.Context, a *domain.Customer) error {
	err := c.customerRepository.Update(ctx, a)
	if err != nil {
		c.logger.Errorf("customerUseCase/Update/Update :%v", err)
		return err
	}

	return nil
}

func (c customerUseCase) Delete(ctx context.Context, a *domain.Customer) error {
	err := c.customerRepository.Delete(ctx, a)
	if err != nil {
		c.logger.Errorf("customerUseCase/Update/Update :%v", err)
		return err
	}

	return nil
}

func NewCustomerUseCase(c domain.CustomerRepository, log *logrus.Logger) domain.CustomerUseCase {
	return &customerUseCase{
		customerRepository: c,
		logger:             log,
	}
}
