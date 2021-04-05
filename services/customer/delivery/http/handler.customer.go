package delivery_http_customer

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

type CustomerHandler struct {
	customerUseCase domain.CustomerUseCase
	logger          *logrus.Logger
}

func NewCustomerHandler(r *mux.Router, c domain.CustomerUseCase, l *logrus.Logger) *mux.Router {
	handler := &CustomerHandler{customerUseCase: c, logger: l}

	r.Handle("/customer", http.HandlerFunc(handler.HandlerGetCustomerList)).Methods(http.MethodGet)
	r.Handle("/customer/{customer_number}", http.HandlerFunc(handler.HandlerGetCustomerByCustomerNumber)).Methods(http.MethodGet)
	r.Handle("/customer", http.HandlerFunc(handler.HandlerCustomerStore)).Methods(http.MethodPost)
	r.Handle("/customer", http.HandlerFunc(handler.HandlerCustomerUpdate)).Methods(http.MethodPut)
	r.Handle("/customer", http.HandlerFunc(handler.HandlerCustomerDelete)).Methods(http.MethodDelete)

	return r
}

func (c *CustomerHandler) HandlerGetCustomerList(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		c.logger.Errorf("%s : %v", "CustomerHandler/HandlerGetCustomerList/ParseForm", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filter, err := util.ParseQueryParams(r)
	if err != nil {
		c.logger.Errorf("%s : %v", "CustomerHandler/HandlerGetCustomerList/ParseQueryParams", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	customers, err := c.customerUseCase.List(r.Context(), domain.CustomerListParam{Filter: filter})
	if err != nil {
		c.logger.Errorf("%s : %v", "CustomerHandler/HandlerGetCustomerList/List", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(customers)
	return
}

func (c *CustomerHandler) HandlerGetCustomerByCustomerNumber(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)

	customerNumber, err := strconv.Atoi(v["customer_number"])
	if err != nil {
		c.logger.Errorf("%s : %v", "CustomerHandler/HandlerGetCustomerByCustomerNumber/parseCustomerNumber", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	customer, err := c.customerUseCase.GetByCustomerNumber(r.Context(), customerNumber)
	if err != nil {
		c.logger.Errorf("%s : %v", "CustomerHandler/HandlerGetCustomerByCustomerNumber/GetByCustomerNumber", err)
		if errors.Cause(err) == sql.ErrNoRows {
			http.Error(w, "Customer not exists", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(customer)
	return
}

func (c *CustomerHandler) HandlerCustomerStore(w http.ResponseWriter, r *http.Request) {
	var param domain.Customer
	err := util.ParseBodyData(r.Context(), r, &param)
	if err != nil {
		c.logger.Errorf("%s : %v", "CustomerHandler/HandlerCustomerStore/ParseBodyData", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.customerUseCase.Store(r.Context(), &param)
	if err != nil {
		c.logger.Errorf("%s : %v", "CustomerHandler/HandlerCustomerStore/Store", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

func (c *CustomerHandler) HandlerCustomerUpdate(w http.ResponseWriter, r *http.Request) {
	var param domain.Customer

	err := util.ParseBodyData(r.Context(), r, &param)
	if err != nil {
		c.logger.Errorf("%s : %v", "CustomerHandler/HandlerCustomerUpdate/ParseBodyData", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.customerUseCase.Update(r.Context(), &param)
	if err != nil {
		c.logger.Errorf("%s : %v", "CustomerHandler/HandlerCustomerUpdate/Store", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}

func (c *CustomerHandler) HandlerCustomerDelete(w http.ResponseWriter, r *http.Request) {
	var param domain.Customer

	err := util.ParseBodyData(r.Context(), r, &param)
	if err != nil {
		c.logger.Errorf("%s : %v", "CustomerHandler/HandlerCustomerDelete/ParseBodyData", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.customerUseCase.Delete(r.Context(), &param)
	if err != nil {
		c.logger.Errorf("%s : %v", "CustomerHandler/HandlerCustomerDelete/Store", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
