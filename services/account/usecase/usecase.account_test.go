package usecase

import (
	"context"
	"database/sql"
	"linkaja-test/domain"
	repository_account_mock "linkaja-test/services/account/repository/mock"
	repository_customer_mock "linkaja-test/services/customer/repository/mock"
	"testing"

	"github.com/pkg/errors"

	"github.com/stretchr/testify/assert"

	"github.com/sirupsen/logrus"

	"github.com/stretchr/testify/mock"
)

func TestAccountUseCase_List(t *testing.T) {
	logger := logrus.New()

	mockAccountRepo := new(repository_account_mock.AccountMockRepository)
	mockCustomerRepo := new(repository_customer_mock.CustomerMockRepository)

	customersData := []domain.Account{
		{
			AccountNumber:  555001,
			CustomerNumber: 1001,
			Balance:        10000,
		},
		{
			AccountNumber:  555002,
			CustomerNumber: 1002,
			Balance:        15000,
		},
	}

	t.Run("Success", func(t *testing.T) {
		mockAccountRepo.On("List", mock.Anything, mock.AnythingOfType("domain.AccountListParam")).Return(customersData, nil).Once()

		customerUseCase := NewAccountUseCase(mockAccountRepo, mockCustomerRepo, logger)

		cDatas, err := customerUseCase.List(context.Background(), domain.AccountListParam{})
		assert.NoError(t, err)
		assert.NotNil(t, cDatas)

		mockAccountRepo.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		mockAccountRepo.On("List", mock.Anything, mock.AnythingOfType("domain.AccountListParam")).Return([]domain.Account{}, errors.New("Unexpected")).Once()

		customerUseCase := NewAccountUseCase(mockAccountRepo, mockCustomerRepo, logger)

		cDatas, err := customerUseCase.List(context.Background(), domain.AccountListParam{})
		assert.Error(t, err)
		assert.Nil(t, cDatas)

		mockAccountRepo.AssertExpectations(t)
	})

}

func TestAccountUseCase_GetByAccountNumber(t *testing.T) {
	logger := logrus.New()

	mockAccountRepo := new(repository_account_mock.AccountMockRepository)
	mockCustomerRepo := new(repository_customer_mock.CustomerMockRepository)

	accountData := domain.Account{
		AccountNumber:  555001,
		CustomerNumber: 1001,
		Balance:        10000,
	}

	customerData := domain.Customer{
		CustomerNumber: 1001,
		Name:           "Bob",
	}

	t.Run("Success", func(t *testing.T) {
		mockAccountRepo.On("GetByAccountNumber", mock.Anything, mock.AnythingOfType("int")).Return(accountData, nil).Once()
		mockCustomerRepo.On("GetByCustomerNumber", mock.Anything, mock.AnythingOfType("int")).Return(customerData, nil).Once()

		customerUseCase := NewAccountUseCase(mockAccountRepo, mockCustomerRepo, logger)

		cData, err := customerUseCase.GetByAccountNumber(context.Background(), 1001)
		assert.NoError(t, err)
		assert.NotNil(t, cData)

		mockAccountRepo.AssertExpectations(t)
	})

	t.Run("Failed-Data-Not-Exists", func(t *testing.T) {
		mockAccountRepo.On("GetByAccountNumber", mock.Anything, mock.AnythingOfType("int")).Return(domain.Account{}, sql.ErrNoRows).Once()
		mockCustomerRepo.On("GetByCustomerNumber", mock.Anything, mock.AnythingOfType("int")).Return(domain.Customer{}, nil).Once()

		customerUseCase := NewAccountUseCase(mockAccountRepo, mockCustomerRepo, logger)

		cData, err := customerUseCase.GetByAccountNumber(context.Background(), 0)
		assert.Error(t, err)
		assert.Equal(t, cData, domain.DetailByAccountNumberResponse{})

		mockAccountRepo.AssertExpectations(t)
	})
}

func TestAccountUseCase_Store(t *testing.T) {
	logger := logrus.New()

	mockAccountRepo := new(repository_account_mock.AccountMockRepository)
	mockCustomerRepo := new(repository_customer_mock.CustomerMockRepository)

	accountData := domain.Account{
		AccountNumber:  555001,
		CustomerNumber: 1001,
		Balance:        10000,
	}

	t.Run("Success", func(t *testing.T) {
		mockAccountRepo.On("Store", mock.Anything, mock.AnythingOfType("*domain.Account")).Return(nil).Once()

		customerUseCase := NewAccountUseCase(mockAccountRepo, mockCustomerRepo, logger)

		err := customerUseCase.Store(context.Background(), &accountData)
		assert.NoError(t, err)

		mockAccountRepo.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		mockAccountRepo.On("Store", mock.Anything, mock.AnythingOfType("*domain.Account")).Return(errors.New("Unexpected")).Once()

		customerUseCase := NewAccountUseCase(mockAccountRepo, mockCustomerRepo, logger)

		err := customerUseCase.Store(context.Background(), &accountData)
		assert.Error(t, err)

		mockAccountRepo.AssertExpectations(t)
	})
}

func TestAccountUseCase_Update(t *testing.T) {
	logger := logrus.New()

	mockAccountRepo := new(repository_account_mock.AccountMockRepository)
	mockCustomerRepo := new(repository_customer_mock.CustomerMockRepository)

	customerData := domain.Account{
		AccountNumber:  555001,
		CustomerNumber: 1001,
		Balance:        10000,
	}

	t.Run("Success", func(t *testing.T) {
		mockAccountRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Account")).Return(nil).Once()

		customerUseCase := NewAccountUseCase(mockAccountRepo, mockCustomerRepo, logger)

		err := customerUseCase.Update(context.Background(), &customerData)
		assert.NoError(t, err)

		mockAccountRepo.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		mockAccountRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Account")).Return(errors.New("Unexpected")).Once()

		customerUseCase := NewAccountUseCase(mockAccountRepo, mockCustomerRepo, logger)

		err := customerUseCase.Update(context.Background(), &customerData)
		assert.Error(t, err)

		mockAccountRepo.AssertExpectations(t)
	})
}

func TestAccountUseCase_Delete(t *testing.T) {
	logger := logrus.New()

	mockAccountRepo := new(repository_account_mock.AccountMockRepository)
	mockCustomerRepo := new(repository_customer_mock.CustomerMockRepository)

	customerData := domain.Account{
		AccountNumber:  555001,
		CustomerNumber: 1001,
		Balance:        10000,
	}

	t.Run("Success", func(t *testing.T) {
		mockAccountRepo.On("Delete", mock.Anything, mock.AnythingOfType("*domain.Account")).Return(nil).Once()

		customerUseCase := NewAccountUseCase(mockAccountRepo, mockCustomerRepo, logger)

		err := customerUseCase.Delete(context.Background(), &customerData)
		assert.NoError(t, err)

		mockAccountRepo.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		mockAccountRepo.On("Delete", mock.Anything, mock.AnythingOfType("*domain.Account")).Return(errors.New("Unexpected")).Once()

		customerUseCase := NewAccountUseCase(mockAccountRepo, mockCustomerRepo, logger)

		err := customerUseCase.Delete(context.Background(), &customerData)
		assert.Error(t, err)

		mockAccountRepo.AssertExpectations(t)
	})
}

func TestAccountUseCase_Transfer(t *testing.T) {
	logger := logrus.New()

	mockAccountRepo := new(repository_account_mock.AccountMockRepository)
	mockCustomerRepo := new(repository_customer_mock.CustomerMockRepository)

	accountSenderData := domain.Account{
		AccountNumber:  555001,
		CustomerNumber: 1001,
		Balance:        10000,
	}

	accountReceiverData := domain.Account{
		AccountNumber:  555002,
		CustomerNumber: 1002,
		Balance:        15000,
	}

	transferParam := domain.TransferParam{
		ToAccountNumber: "555002",
		Amount:          1000,
	}

	t.Run("Success", func(t *testing.T) {
		mockAccountRepo.On("GetByAccountNumber", mock.Anything, 555001).Return(accountSenderData, nil).Once()
		mockAccountRepo.On("GetByAccountNumber", mock.Anything, 555002).Return(accountReceiverData, nil).Once()
		mockAccountRepo.On("Update", mock.Anything, &accountSenderData).Return(nil).Once()
		mockAccountRepo.On("Update", mock.Anything, &accountReceiverData).Return(nil).Once()

		accountSenderData.Balance = accountSenderData.Balance - transferParam.Amount
		accountReceiverData.Balance = accountReceiverData.Balance + transferParam.Amount

		customerUseCase := NewAccountUseCase(mockAccountRepo, mockCustomerRepo, logger)

		err := customerUseCase.Transfer(context.Background(), accountSenderData.AccountNumber, transferParam)
		assert.NoError(t, err)
		assert.Equal(t, 9000, accountSenderData.Balance)
		assert.Equal(t, 16000, accountReceiverData.Balance)

	})

	t.Run("Account-sender-not-exists", func(t *testing.T) {
		mockAccountRepo.On("GetByAccountNumber", mock.Anything, mock.AnythingOfType("int")).Return(domain.Account{}, sql.ErrNoRows).Once()

		customerUseCase := NewAccountUseCase(mockAccountRepo, mockCustomerRepo, logger)

		err := customerUseCase.Transfer(context.Background(), accountSenderData.AccountNumber, transferParam)
		assert.Error(t, err)

		mockAccountRepo.AssertExpectations(t)
	})

	t.Run("Account-receiver-not-exists", func(t *testing.T) {
		mockAccountRepo.On("GetByAccountNumber", mock.Anything, mock.AnythingOfType("int")).Return(accountSenderData, nil).Once()
		mockAccountRepo.On("GetByAccountNumber", mock.Anything, mock.AnythingOfType("int")).Return(domain.Account{}, sql.ErrNoRows).Once()

		customerUseCase := NewAccountUseCase(mockAccountRepo, mockCustomerRepo, logger)

		err := customerUseCase.Transfer(context.Background(), accountSenderData.AccountNumber, transferParam)
		assert.Error(t, err)

		mockAccountRepo.AssertExpectations(t)
	})

	t.Run("Insufficient-balance", func(t *testing.T) {
		mockAccountRepo.On("GetByAccountNumber", mock.Anything, 555001).Return(accountSenderData, nil).Once()
		mockAccountRepo.On("GetByAccountNumber", mock.Anything, 555002).Return(accountReceiverData, nil).Once()
		mockAccountRepo.On("Update", mock.Anything, &accountSenderData).Return(nil).Once()
		mockAccountRepo.On("Update", mock.Anything, &accountReceiverData).Return(nil).Once()

		customerUseCase := NewAccountUseCase(mockAccountRepo, mockCustomerRepo, logger)

		transferParam.Amount = 100000
		err := customerUseCase.Transfer(context.Background(), accountSenderData.AccountNumber, transferParam)
		assert.Error(t, err)
	})
}
