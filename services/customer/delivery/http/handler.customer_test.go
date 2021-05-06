package delivery_http_customer

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/oniharnantyo/golang-backend-example/domain"
	customer_usecase_mock "github.com/oniharnantyo/golang-backend-example/services/customer/usecase/mock"

	"github.com/sirupsen/logrus"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"

	"github.com/bxcodec/faker"
)

func TestCustomerHandler_HandlerGetCustomerList(t *testing.T) {
	var mockCustomer domain.Customer
	logger := logrus.New()

	err := faker.FakeData(&mockCustomer)
	assert.NoError(t, err)

	mockCustomerUseCase := new(customer_usecase_mock.CustomerMockUseCase)

	mockCustomers := make([]domain.Customer, 0)
	mockCustomers = append(mockCustomers, mockCustomer)

	mockCustomerUseCase.On("List", mock.Anything, mock.AnythingOfType("domain.CustomerListParam")).Return(mockCustomers, nil).Once()

	r := gin.Default()
	r = NewCustomerHandler(r, mockCustomerUseCase, logger)

	req, err := http.NewRequest(http.MethodGet, "/customer?limit=10&offset=0&search=&order=asc", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)
	assert.Equal(t, 200, rec.Code)
	mockCustomerUseCase.AssertExpectations(t)
}

func TestCustomerHandler_HandlerGetCustomerByCustomerNumber(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var mockCustomer domain.Customer
		logger := logrus.New()

		err := faker.FakeData(&mockCustomer)
		assert.NoError(t, err)

		mockCustomerUseCase := new(customer_usecase_mock.CustomerMockUseCase)

		mockCustomerUseCase.On("GetByCustomerNumber", mock.Anything, mock.AnythingOfType("int")).Return(mockCustomer, nil).Once()

		r := gin.Default()
		r = NewCustomerHandler(r, mockCustomerUseCase, logger)

		req, err := http.NewRequest(http.MethodGet, "/customer/1", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockCustomerUseCase.AssertExpectations(t)
	})

	t.Run("Customer-not-found", func(t *testing.T) {
		var mockCustomer domain.Customer
		logger := logrus.New()

		err := faker.FakeData(&mockCustomer)
		assert.NoError(t, err)

		mockCustomerUseCase := new(customer_usecase_mock.CustomerMockUseCase)

		mockCustomerUseCase.On("GetByCustomerNumber", mock.Anything, mock.AnythingOfType("int")).Return(domain.Customer{}, sql.ErrNoRows).Once()

		r := gin.Default()
		r = NewCustomerHandler(r, mockCustomerUseCase, logger)

		req, err := http.NewRequest(http.MethodGet, "/customer/1", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)
		mockCustomerUseCase.AssertExpectations(t)
	})
}

func TestCustomerHandler_HandlerCustomerStore(t *testing.T) {
	var mockCustomer domain.Customer
	logger := logrus.New()

	err := faker.FakeData(&mockCustomer)
	assert.NoError(t, err)

	mockCustomerUseCase := new(customer_usecase_mock.CustomerMockUseCase)

	mockCustomerUseCase.On("Store", mock.Anything, mock.AnythingOfType("*domain.Customer")).Return(mockCustomer, nil).Once()

	r := gin.Default()
	r = NewCustomerHandler(r, mockCustomerUseCase, logger)

	reqBody, err := json.Marshal(mockCustomer)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/customer", bytes.NewBuffer(reqBody))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusCreated, rec.Code)
	mockCustomerUseCase.AssertExpectations(t)
}

func TestCustomerHandler_HandlerCustomerUpdate(t *testing.T) {
	var mockCustomer domain.Customer
	logger := logrus.New()

	err := faker.FakeData(&mockCustomer)
	assert.NoError(t, err)

	mockCustomerUseCase := new(customer_usecase_mock.CustomerMockUseCase)

	mockCustomerUseCase.On("Update", mock.Anything, mock.AnythingOfType("*domain.Customer")).Return(mockCustomer, nil).Once()

	r := gin.Default()
	r = NewCustomerHandler(r, mockCustomerUseCase, logger)

	reqBody, err := json.Marshal(mockCustomer)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPut, "/customer", bytes.NewBuffer(reqBody))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockCustomerUseCase.AssertExpectations(t)
}

func TestCustomerHandler_HandlerCustomerDelete(t *testing.T) {
	var mockCustomer domain.Customer
	logger := logrus.New()

	err := faker.FakeData(&mockCustomer)
	assert.NoError(t, err)

	mockCustomerUseCase := new(customer_usecase_mock.CustomerMockUseCase)

	mockCustomerUseCase.On("Delete", mock.Anything, mock.AnythingOfType("*domain.Customer")).Return(mockCustomer, nil).Once()

	r := gin.Default()
	r = NewCustomerHandler(r, mockCustomerUseCase, logger)

	reqBody, err := json.Marshal(mockCustomer)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodDelete, "/customer", bytes.NewBuffer(reqBody))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockCustomerUseCase.AssertExpectations(t)
}
