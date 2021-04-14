package delivery_http_account

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang-backend-example/domain"
	account_usecase_mock "golang-backend-example/services/account/usecase/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"

	"github.com/bxcodec/faker"
)

func TestAccountHandler_HandlerGetAccountList(t *testing.T) {
	var mockAccount domain.Account
	logger := logrus.New()

	err := faker.FakeData(&mockAccount)
	assert.NoError(t, err)

	mockAccountUseCase := new(account_usecase_mock.AccountMockUseCase)

	mockAccounts := make([]domain.Account, 0)
	mockAccounts = append(mockAccounts, mockAccount)

	mockAccountUseCase.On("List", mock.Anything, mock.AnythingOfType("domain.AccountListParam")).Return(mockAccounts, nil).Once()

	r := gin.Default()
	r = NewAccountHandler(r, mockAccountUseCase, logger)

	req, err := http.NewRequest(http.MethodGet, "/account?limit=10&offset=0&search=&order=asc", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)
	assert.Equal(t, 200, rec.Code)
	mockAccountUseCase.AssertExpectations(t)
}

func TestAccountHandler_HandlerGetAccountByAccountNumber(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var mockAccount domain.Account
		logger := logrus.New()

		err := faker.FakeData(&mockAccount)
		assert.NoError(t, err)

		mockAccountUseCase := new(account_usecase_mock.AccountMockUseCase)

		mockAccountUseCase.On("GetByAccountNumber", mock.Anything, mock.AnythingOfType("int")).Return(domain.DetailByAccountNumberResponse{}, nil).Once()

		r := gin.Default()
		r = NewAccountHandler(r, mockAccountUseCase, logger)

		req, err := http.NewRequest(http.MethodGet, "/account/1001", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockAccountUseCase.AssertExpectations(t)
	})

	t.Run("Account-not-found", func(t *testing.T) {
		var mockAccount domain.Account
		logger := logrus.New()

		err := faker.FakeData(&mockAccount)
		assert.NoError(t, err)

		mockAccountUseCase := new(account_usecase_mock.AccountMockUseCase)

		mockAccountUseCase.On("GetByAccountNumber", mock.Anything, mock.AnythingOfType("int")).Return(domain.DetailByAccountNumberResponse{}, sql.ErrNoRows).Once()

		r := gin.Default()
		r = NewAccountHandler(r, mockAccountUseCase, logger)

		req, err := http.NewRequest(http.MethodGet, "/account/1", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)
		mockAccountUseCase.AssertExpectations(t)
	})
}

func TestAccountHandler_HandlerAccountStore(t *testing.T) {
	var mockAccount domain.Account
	logger := logrus.New()

	err := faker.FakeData(&mockAccount)
	assert.NoError(t, err)

	mockAccountUseCase := new(account_usecase_mock.AccountMockUseCase)

	mockAccountUseCase.On("Store", mock.Anything, mock.AnythingOfType("*domain.Account")).Return(nil).Once()

	r := gin.Default()
	r = NewAccountHandler(r, mockAccountUseCase, logger)

	reqBody, err := json.Marshal(mockAccount)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/account", bytes.NewBuffer(reqBody))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusCreated, rec.Code)
	mockAccountUseCase.AssertExpectations(t)
}

func TestAccountHandler_HandlerAccountUpdate(t *testing.T) {
	var mockAccount domain.Account
	logger := logrus.New()

	err := faker.FakeData(&mockAccount)
	assert.NoError(t, err)

	mockAccountUseCase := new(account_usecase_mock.AccountMockUseCase)

	mockAccountUseCase.On("Update", mock.Anything, mock.AnythingOfType("*domain.Account")).Return(nil).Once()

	r := gin.Default()
	r = NewAccountHandler(r, mockAccountUseCase, logger)

	reqBody, err := json.Marshal(mockAccount)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPut, "/account", bytes.NewBuffer(reqBody))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockAccountUseCase.AssertExpectations(t)
}

func TestAccountHandler_HandlerAccountDelete(t *testing.T) {
	var mockAccount domain.Account
	logger := logrus.New()

	err := faker.FakeData(&mockAccount)
	assert.NoError(t, err)

	mockAccountUseCase := new(account_usecase_mock.AccountMockUseCase)

	mockAccountUseCase.On("Delete", mock.Anything, mock.AnythingOfType("*domain.Account")).Return(nil).Once()

	r := gin.Default()
	r = NewAccountHandler(r, mockAccountUseCase, logger)

	reqBody, err := json.Marshal(mockAccount)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodDelete, "/account", bytes.NewBuffer(reqBody))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockAccountUseCase.AssertExpectations(t)
}
