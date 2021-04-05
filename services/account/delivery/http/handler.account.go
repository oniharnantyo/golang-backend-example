package delivery_http_account

import (
	"database/sql"
	"encoding/json"
	"linkaja-test/domain"
	"linkaja-test/util"
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	accountUseCase domain.AccountUseCase
	logger         *logrus.Logger
}

func NewAccountHandler(r *mux.Router, c domain.AccountUseCase, l *logrus.Logger) *mux.Router {
	handler := &AccountHandler{accountUseCase: c, logger: l}

	r.Handle("/account", http.HandlerFunc(handler.HandlerGetAccountList)).Methods(http.MethodGet)
	r.Handle("/account/{account_number}", http.HandlerFunc(handler.HandlerGetAccountByAccountNumber)).Methods(http.MethodGet)
	r.Handle("/account", http.HandlerFunc(handler.HandlerAccountStore)).Methods(http.MethodPost)
	r.Handle("/account", http.HandlerFunc(handler.HandlerAccountUpdate)).Methods(http.MethodPut)
	r.Handle("/account", http.HandlerFunc(handler.HandlerAccountDelete)).Methods(http.MethodDelete)
	r.Handle("/account/{from_account_number}/transfer", http.HandlerFunc(handler.HandlerAccountTransfer)).Methods(http.MethodPost)

	return r
}

func (c *AccountHandler) HandlerGetAccountList(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		c.logger.Errorf("%s : %v", "AccountHandler/HandlerGetAccountList/ParseForm", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filter, err := util.ParseQueryParams(r)
	if err != nil {
		c.logger.Errorf("%s : %v", "AccountHandler/HandlerGetAccountList/ParseQueryParams", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	customers, err := c.accountUseCase.List(r.Context(), domain.AccountListParam{Filter: filter})
	if err != nil {
		c.logger.Errorf("%s : %v", "AccountHandler/HandlerGetAccountList/List", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(customers)
	return
}

func (c *AccountHandler) HandlerGetAccountByAccountNumber(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)

	customerNumber, err := strconv.Atoi(v["account_number"])
	if err != nil {
		c.logger.Errorf("%s : %v", "AccountHandler/HandlerGetAccountByAccountNumber/parseAccountNumber", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := c.accountUseCase.GetByAccountNumber(r.Context(), customerNumber)
	if err != nil {
		c.logger.Errorf("%s : %v", "AccountHandler/HandlerGetAccountByAccountNumber/GetByAccountNumber", err)
		if errors.Cause(err) == sql.ErrNoRows {
			http.Error(w, "Account not exists", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(account)
	return
}

func (c *AccountHandler) HandlerAccountStore(w http.ResponseWriter, r *http.Request) {
	var param domain.Account
	err := util.ParseBodyData(r.Context(), r, &param)
	if err != nil {
		c.logger.Errorf("%s : %v", "AccountHandler/HandlerAccountStore/ParseBodyData", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.accountUseCase.Store(r.Context(), &param)
	if err != nil {
		c.logger.Errorf("%s : %v", "AccountHandler/HandlerAccountStore/Store", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

func (c *AccountHandler) HandlerAccountUpdate(w http.ResponseWriter, r *http.Request) {
	var param domain.Account

	err := util.ParseBodyData(r.Context(), r, &param)
	if err != nil {
		c.logger.Errorf("%s : %v", "AccountHandler/HandlerAccountUpdate/ParseBodyData", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.accountUseCase.Update(r.Context(), &param)
	if err != nil {
		c.logger.Errorf("%s : %v", "AccountHandler/HandlerAccountUpdate/Store", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}

func (c *AccountHandler) HandlerAccountDelete(w http.ResponseWriter, r *http.Request) {
	var param domain.Account

	err := util.ParseBodyData(r.Context(), r, &param)
	if err != nil {
		c.logger.Errorf("%s : %v", "AccountHandler/HandlerAccountDelete/ParseBodyData", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.accountUseCase.Delete(r.Context(), &param)
	if err != nil {
		c.logger.Errorf("%s : %v", "AccountHandler/HandlerAccountDelete/Store", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}

func (c *AccountHandler) HandlerAccountTransfer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)

	fromAccountNumber, err := strconv.Atoi(vars["from_account_number"])
	if err != nil {
		c.logger.Errorf("%s : %v", "AccountHandler/HandlerAccountTransfer/parseFromAccountNumber", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var param domain.TransferParam
	err = util.ParseBodyData(ctx, r, &param)
	if err != nil {
		c.logger.Errorf("%s : %v", "AccountHandler/HandlerAccountTransfer/ParseBodyData", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.accountUseCase.Transfer(ctx, fromAccountNumber, param)
	if err != nil {
		c.logger.Errorf("%s : %v", "AccountHandler/HandlerAccountTransfer/Transfer", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
