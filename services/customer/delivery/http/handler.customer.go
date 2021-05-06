package delivery_http_customer

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oniharnantyo/golang-backend-example/domain"

	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"
)

type CustomerHandler struct {
	customerUseCase domain.CustomerUseCase
	logger          *logrus.Logger
}

func NewCustomerHandler(r *gin.Engine, c domain.CustomerUseCase, l *logrus.Logger) *gin.Engine {
	handler := &CustomerHandler{customerUseCase: c, logger: l}

	r.GET("/customer", handler.HandlerGetCustomerList)
	r.GET("/customer/:customer_number", handler.HandlerGetCustomerByCustomerNumber)
	r.POST("/customer", handler.HandlerCustomerStore)
	r.PUT("/customer", handler.HandlerCustomerUpdate)
	r.DELETE("/customer", handler.HandlerCustomerDelete)

	return r
}

func (c *CustomerHandler) HandlerGetCustomerList(ctx *gin.Context) {
	var param domain.CustomerListParam
	err := ctx.BindQuery(&param)
	if err != nil {
		c.logger.Errorf("%s : %v", "CustomerHandler/HandlerGetCustomerList/ParseQueryParams", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	customers, err := c.customerUseCase.List(ctx, param)
	if err != nil {
		c.logger.Errorf("%s : %v", "CustomerHandler/HandlerGetCustomerList/List", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, customers)
	return
}

func (c *CustomerHandler) HandlerGetCustomerByCustomerNumber(ctx *gin.Context) {
	customerNumber, err := strconv.Atoi(ctx.Param("customer_number"))
	if err != nil {
		c.logger.Errorf("%s : %v", "CustomerHandler/HandlerGetCustomerByCustomerNumber/parseCustomerNumber", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	customer, err := c.customerUseCase.GetByCustomerNumber(ctx, customerNumber)
	if err != nil {
		c.logger.Errorf("%s : %v", "CustomerHandler/HandlerGetCustomerByCustomerNumber/GetByCustomerNumber", err)
		if errors.Cause(err) == sql.ErrNoRows {
			ctx.AbortWithError(http.StatusNotFound, errors.New("Customer not exists"))
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, customer)
	return
}

func (c *CustomerHandler) HandlerCustomerStore(ctx *gin.Context) {
	var param domain.Customer
	err := ctx.Bind(&param)
	if err != nil {
		c.logger.Errorf("%s : %v", "CustomerHandler/HandlerCustomerStore/ParseBodyData", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = c.customerUseCase.Store(ctx, &param)
	if err != nil {
		c.logger.Errorf("%s : %v", "CustomerHandler/HandlerCustomerStore/Store", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusCreated)
	return
}

func (c *CustomerHandler) HandlerCustomerUpdate(ctx *gin.Context) {
	var param domain.Customer

	err := ctx.Bind(&param)
	if err != nil {
		c.logger.Errorf("%s : %v", "CustomerHandler/HandlerCustomerUpdate/ParseBodyData", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = c.customerUseCase.Update(ctx, &param)
	if err != nil {
		c.logger.Errorf("%s : %v", "CustomerHandler/HandlerCustomerUpdate/Store", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
	return
}

func (c *CustomerHandler) HandlerCustomerDelete(ctx *gin.Context) {
	var param domain.Customer

	err := ctx.Bind(&param)
	if err != nil {
		c.logger.Errorf("%s : %v", "CustomerHandler/HandlerCustomerDelete/ParseBodyData", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = c.customerUseCase.Delete(ctx, &param)
	if err != nil {
		c.logger.Errorf("%s : %v", "CustomerHandler/HandlerCustomerDelete/Store", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
	return
}
