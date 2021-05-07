package usecase

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/oniharnantyo/golang-backend-example/domain"

	"golang.org/x/crypto/bcrypt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type accountUseCase struct {
	authUseCase        domain.AuthUseCase
	accountRepository  domain.AccountRepository
	customerRepository domain.CustomerRepository
	logger             *logrus.Logger
}

func (c accountUseCase) List(ctx context.Context, param domain.AccountListParam) ([]domain.Account, error) {
	accounts, err := c.accountRepository.List(ctx, param)
	if err != nil {
		c.logger.Errorf("accountUseCase/List/List :%v", err)
		return nil, errors.Wrap(err, "accountUseCase/List/List")
	}

	return accounts, nil
}

func (c accountUseCase) GetByAccountNumber(ctx context.Context, accountNumber int) (domain.DetailByAccountNumberResponse, error) {
	account, err := c.accountRepository.GetByAccountNumber(ctx, accountNumber)
	if err != nil {
		c.logger.Errorf("accountUseCase/GetByAccountNumber/GetByAccountNumber :%v", err)
		return domain.DetailByAccountNumberResponse{}, err
	}

	customer, err := c.customerRepository.GetByCustomerNumber(ctx, account.CustomerNumber)
	if err != nil {
		c.logger.Errorf("accountUseCase/GetByAccountNumber/GetByCustomerNumber :%v", err)
		return domain.DetailByAccountNumberResponse{}, err
	}

	return domain.DetailByAccountNumberResponse{
		AccountNumber: account.AccountNumber,
		CustomerName:  customer.Name,
		Balance:       account.Balance,
	}, nil
}

func (c accountUseCase) Store(ctx context.Context, a *domain.Account) error {
	err := c.accountRepository.Store(ctx, a)
	if err != nil {
		c.logger.Errorf("accountUseCase/Store/Store :%v", err)
		return err
	}

	return nil
}

func (c accountUseCase) Update(ctx context.Context, a *domain.Account) error {
	err := c.accountRepository.Update(ctx, a)
	if err != nil {
		c.logger.Errorf("accountUseCase/Update/Update :%v", err)
		return err
	}

	return nil
}

func (c accountUseCase) Delete(ctx context.Context, a *domain.Account) error {
	err := c.accountRepository.Delete(ctx, a)
	if err != nil {
		c.logger.Errorf("accountUseCase/Update/Update :%v", err)
		return err
	}

	return nil
}

func (c accountUseCase) Transfer(ctx context.Context, fromAccountNumber int, param domain.TransferParam) error {
	senderAccount, err := c.accountRepository.GetByAccountNumber(ctx, fromAccountNumber)
	if err != nil {
		c.logger.Errorf("accountUseCase/Transfer/GetSenderAccountByAccountNumber :%v", err)
		if errors.Cause(err) == sql.ErrNoRows {
			return errors.New("Sender account not found")
		}
		return err
	}

	toAccountNumber, err := strconv.Atoi(param.ToAccountNumber)
	if err != nil {
		c.logger.Errorf("accountUseCase/Transfer/parserToAccountNumber :%v", err)
		return err
	}

	receiverAccount, err := c.accountRepository.GetByAccountNumber(ctx, toAccountNumber)
	if err != nil {
		c.logger.Errorf("accountUseCase/Transfer/GetByReceiverAccountAccountNumber :%v", err)
		if errors.Cause(err) == sql.ErrNoRows {
			return errors.New("Receiver account not found")
		}
		return err
	}

	// Validate sending to the same account as the sender
	if senderAccount.AccountNumber == receiverAccount.AccountNumber {
		return errors.New("Sender and receiver is same account")
	}

	// Validate sender account balance
	if senderAccount.Balance < param.Amount {
		c.logger.Errorf("accountUseCase/Transfer/validateBalance :%v", errors.New("Insufficient balance"))
		return errors.New("Insufficient balance")
	}

	senderAccount.Balance = senderAccount.Balance - param.Amount
	err = c.accountRepository.Update(ctx, &senderAccount)
	if err != nil {
		c.logger.Errorf("accountUseCase/Transfer/senderAccount/BeginTx :%v", err)
		return err
	}

	receiverAccount.Balance = receiverAccount.Balance + param.Amount
	err = c.accountRepository.Update(ctx, &receiverAccount)
	if err != nil {
		c.logger.Errorf("accountUseCase/Transfer/receiverAccount/UpdateTx :%v", err)
		return err
	}

	return nil
}

func (c accountUseCase) Login(ctx context.Context, param domain.AccountLoginParam) (domain.LoginResponse, error) {
	account, err := c.accountRepository.GetByEmail(ctx, param.Email)
	if err != nil {
		c.logger.Errorf("accountUseCase/Transfer/receiverAccount/GetByEmail :%v", err)
		return domain.LoginResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(param.Password))
	if err != nil {
		c.logger.Errorf("accountUseCase/Transfer/receiverAccount/CompareHashAndPassword :%v", err)
		return domain.LoginResponse{}, err
	}

	token, err := c.authUseCase.CreateAuth(ctx, account)
	if err != nil {
		c.logger.Errorf("accountUseCase/Transfer/receiverAccount/CreateToken :%v", err)
		return domain.LoginResponse{}, err
	}

	return domain.LoginResponse{token.AccessToken}, nil
}

func NewAccountUseCase(au domain.AuthUseCase, a domain.AccountRepository, c domain.CustomerRepository, log *logrus.Logger) domain.AccountUseCase {
	return &accountUseCase{
		authUseCase:        au,
		accountRepository:  a,
		customerRepository: c,
		logger:             log,
	}
}
