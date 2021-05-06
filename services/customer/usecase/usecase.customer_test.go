package usecase

import (
	"context"
	"database/sql"
	"testing"

	"github.com/oniharnantyo/golang-backend-example/domain"
	repository_customer_mock "github.com/oniharnantyo/golang-backend-example/services/customer/repository/mock"

	"github.com/pkg/errors"

	"github.com/stretchr/testify/assert"

	"github.com/sirupsen/logrus"

	"github.com/stretchr/testify/mock"
)

func TestCustomerUseCase_List(t *testing.T) {
	logger := logrus.New()

	mockCustomerRepo := new(repository_customer_mock.CustomerMockRepository)

	customersData := []domain.Customer{
		{
			CustomerNumber: 1001,
			Name:           "Bob Martin",
		},
		{
			CustomerNumber: 1002,
			Name:           "Linus Torvalds",
		},
	}

	t.Run("Success", func(t *testing.T) {
		mockCustomerRepo.On("List", mock.Anything, mock.AnythingOfType("domain.CustomerListParam")).Return(customersData, nil).Once()

		customerUseCase := NewCustomerUseCase(mockCustomerRepo, logger)

		cDatas, err := customerUseCase.List(context.Background(), domain.CustomerListParam{})
		assert.NoError(t, err)
		assert.NotNil(t, cDatas)

		mockCustomerRepo.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		mockCustomerRepo.On("List", mock.Anything, mock.AnythingOfType("domain.CustomerListParam")).Return([]domain.Customer{}, errors.New("Unexpected")).Once()

		customerUseCase := NewCustomerUseCase(mockCustomerRepo, logger)

		cDatas, err := customerUseCase.List(context.Background(), domain.CustomerListParam{})
		assert.Error(t, err)
		assert.Nil(t, cDatas)

		mockCustomerRepo.AssertExpectations(t)
	})

}

func TestCustomerUseCase_GetByCustomerNumber(t *testing.T) {
	logger := logrus.New()

	mockCustomerRepo := new(repository_customer_mock.CustomerMockRepository)

	customerData := domain.Customer{
		CustomerNumber: 1001,
		Name:           "Bob Martin",
	}

	t.Run("Success", func(t *testing.T) {
		mockCustomerRepo.On("GetByCustomerNumber", mock.Anything, mock.AnythingOfType("int")).Return(customerData, nil).Once()

		customerUseCase := NewCustomerUseCase(mockCustomerRepo, logger)

		cData, err := customerUseCase.GetByCustomerNumber(context.Background(), 1001)
		assert.NoError(t, err)
		assert.NotNil(t, cData)

		mockCustomerRepo.AssertExpectations(t)
	})

	t.Run("Failed-Data-Not-Exists", func(t *testing.T) {
		mockCustomerRepo.On("GetByCustomerNumber", mock.Anything, mock.AnythingOfType("int")).Return(domain.Customer{}, sql.ErrNoRows).Once()

		customerUseCase := NewCustomerUseCase(mockCustomerRepo, logger)

		cData, err := customerUseCase.GetByCustomerNumber(context.Background(), 0)
		assert.Error(t, err)
		assert.Equal(t, cData, domain.Customer{})

		mockCustomerRepo.AssertExpectations(t)
	})
}

func TestCustomerUseCase_Store(t *testing.T) {
	logger := logrus.New()

	mockCustomerRepo := new(repository_customer_mock.CustomerMockRepository)

	customerData := domain.Customer{
		CustomerNumber: 1001,
		Name:           "Bob Martin",
	}

	t.Run("Success", func(t *testing.T) {
		mockCustomerRepo.On("Store", mock.Anything, mock.AnythingOfType("*domain.Customer")).Return(nil).Once()

		customerUseCase := NewCustomerUseCase(mockCustomerRepo, logger)

		err := customerUseCase.Store(context.Background(), &customerData)
		assert.NoError(t, err)

		mockCustomerRepo.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		mockCustomerRepo.On("Store", mock.Anything, mock.AnythingOfType("*domain.Customer")).Return(errors.New("Unexpected")).Once()

		customerUseCase := NewCustomerUseCase(mockCustomerRepo, logger)

		err := customerUseCase.Store(context.Background(), &customerData)
		assert.Error(t, err)

		mockCustomerRepo.AssertExpectations(t)
	})
}

func TestCustomerUseCase_Update(t *testing.T) {
	logger := logrus.New()

	mockCustomerRepo := new(repository_customer_mock.CustomerMockRepository)

	customerData := domain.Customer{
		CustomerNumber: 1001,
		Name:           "Bob Martin",
	}

	t.Run("Success", func(t *testing.T) {
		mockCustomerRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Customer")).Return(nil).Once()

		customerUseCase := NewCustomerUseCase(mockCustomerRepo, logger)

		err := customerUseCase.Update(context.Background(), &customerData)
		assert.NoError(t, err)

		mockCustomerRepo.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		mockCustomerRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Customer")).Return(errors.New("Unexpected")).Once()

		customerUseCase := NewCustomerUseCase(mockCustomerRepo, logger)

		err := customerUseCase.Update(context.Background(), &customerData)
		assert.Error(t, err)

		mockCustomerRepo.AssertExpectations(t)
	})
}

func TestCustomerUseCase_Delete(t *testing.T) {
	logger := logrus.New()

	mockCustomerRepo := new(repository_customer_mock.CustomerMockRepository)

	customerData := domain.Customer{
		CustomerNumber: 1001,
		Name:           "Bob Martin",
	}

	t.Run("Success", func(t *testing.T) {
		mockCustomerRepo.On("Delete", mock.Anything, mock.AnythingOfType("*domain.Customer")).Return(nil).Once()

		customerUseCase := NewCustomerUseCase(mockCustomerRepo, logger)

		err := customerUseCase.Delete(context.Background(), &customerData)
		assert.NoError(t, err)

		mockCustomerRepo.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		mockCustomerRepo.On("Delete", mock.Anything, mock.AnythingOfType("*domain.Customer")).Return(errors.New("Unexpected")).Once()

		customerUseCase := NewCustomerUseCase(mockCustomerRepo, logger)

		err := customerUseCase.Delete(context.Background(), &customerData)
		assert.Error(t, err)

		mockCustomerRepo.AssertExpectations(t)
	})
}
