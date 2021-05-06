package delivery_http_account

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oniharnantyo/golang-backend-example/domain"
	account_usecase_mock "github.com/oniharnantyo/golang-backend-example/services/account/usecase/mock"

	"github.com/gin-gonic/gin"

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

func TestAccountHandler_HandlerAccountLogin(t *testing.T) {
	logger := logrus.New()

	account := domain.Account{
		AccountNumber:  1,
		CustomerNumber: 1,
		Balance:        1000,
		Email:          "mail@email.com",
		Password:       "$2y$12$55Pvvir6aXTbi3tE5toEyuUMgPCJ1uytiVREzrSHDgXoNFva7kLOK",
	}

	mockAccountUseCase := new(account_usecase_mock.AccountMockUseCase)

	t.Run("account-not-found", func(t *testing.T) {
		mockAccountUseCase.On("Login", mock.Anything, mock.AnythingOfType("domain.AccountLoginParam")).Return(domain.LoginResponse{}, sql.ErrNoRows).Once()

		r := gin.Default()
		r = NewAccountHandler(r, mockAccountUseCase, logger)

		accountMarshal, err := json.Marshal(&account)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/account/login", bytes.NewBuffer(accountMarshal))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		mockAccountUseCase.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		mockAccountUseCase.On("Login", mock.Anything, mock.AnythingOfType("domain.AccountLoginParam")).Return(domain.LoginResponse{Token: "token"}, nil).Once()

		r := gin.Default()
		r = NewAccountHandler(r, mockAccountUseCase, logger)

		param := domain.AccountLoginParam{
			Email:    "email@mail.com",
			Password: "secret",
		}

		paramMarshal, err := json.Marshal(&param)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/account/login", bytes.NewBuffer(paramMarshal))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		resp, err := ioutil.ReadAll(rec.Body)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"token\":\"token\"}", string(resp))
		mockAccountUseCase.AssertExpectations(t)
	})
}
