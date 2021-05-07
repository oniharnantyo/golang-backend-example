package delivery_http_account

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/oniharnantyo/golang-backend-example/util"

	"github.com/oniharnantyo/golang-backend-example/middleware"

	"github.com/oniharnantyo/golang-backend-example/domain"

	"github.com/gin-gonic/gin"

	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"
)

type AccountHandler struct {
	accountUseCase domain.AccountUseCase
	logger         *logrus.Logger
}

func NewAccountHandler(r *gin.Engine, ctx domain.AccountUseCase, l *logrus.Logger) *gin.Engine {
	handler := &AccountHandler{accountUseCase: ctx, logger: l}

	r.GET("/account", handler.HandlerGetAccountList)
	r.GET("/account/:account_number", handler.HandlerGetAccountByAccountNumber)
	r.POST("/account", handler.HandlerAccountStore)
	r.PUT("/account", handler.HandlerAccountUpdate)
	r.DELETE("/account", handler.HandlerAccountDelete)
	r.POST("/account/login", handler.HandlerLogin)
	r.POST("/account/:from_account_number/transfer", middleware.JWT(), handler.HandlerAccountTransfer)

	return r
}

func (a *AccountHandler) HandlerGetAccountList(ctx *gin.Context) {
	var filter domain.AccountListParam
	err := ctx.ShouldBind(&filter)
	if err != nil {
		a.logger.Errorf("%s : %v", "AccountHandler/HandlerGetAccountList/ShouldBindQuery", err)
		ctx.String(http.StatusBadRequest, "Bad request")
		return
	}

	fmt.Println(filter)

	accounts, err := a.accountUseCase.List(ctx, filter)
	if err != nil {
		a.logger.Errorf("%s : %v", "AccountHandler/HandlerGetAccountList/List", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, accounts)
	return
}

func (a *AccountHandler) HandlerGetAccountByAccountNumber(ctx *gin.Context) {
	customerNumber, err := strconv.Atoi(ctx.Param("account_number"))
	if err != nil {
		a.logger.Errorf("%s : %v", "AccountHandler/HandlerGetAccountByAccountNumber/parseAccountNumber", err)
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Account not exists"))
		return
	}

	account, err := a.accountUseCase.GetByAccountNumber(ctx, customerNumber)
	if err != nil {
		a.logger.Errorf("%s : %v", "AccountHandler/HandlerGetAccountByAccountNumber/GetByAccountNumber", err)
		if errors.Cause(err) == sql.ErrNoRows {
			ctx.AbortWithError(http.StatusNotFound, errors.New("Account not exists"))
			return
		}
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, account)
	return
}

func (a *AccountHandler) HandlerAccountStore(ctx *gin.Context) {
	var param domain.Account
	err := ctx.Bind(&param)
	if err != nil {
		a.logger.Errorf("%s : %v", "AccountHandler/HandlerAccountStore/ParseBodyData", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = a.accountUseCase.Store(ctx, &param)
	if err != nil {
		a.logger.Errorf("%s : %v", "AccountHandler/HandlerAccountStore/Store", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusCreated)
	return
}

func (a *AccountHandler) HandlerAccountUpdate(ctx *gin.Context) {
	var param domain.Account

	err := ctx.Bind(&param)
	if err != nil {
		a.logger.Errorf("%s : %v", "AccountHandler/HandlerAccountUpdate/ParseBodyData", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = a.accountUseCase.Update(ctx, &param)
	if err != nil {
		a.logger.Errorf("%s : %v", "AccountHandler/HandlerAccountUpdate/Store", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
	return
}

func (a *AccountHandler) HandlerAccountDelete(ctx *gin.Context) {
	var param domain.Account

	err := ctx.Bind(&param)
	if err != nil {
		a.logger.Errorf("%s : %v", "AccountHandler/HandlerAccountDelete/ParseBodyData", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = a.accountUseCase.Delete(ctx, &param)
	if err != nil {
		a.logger.Errorf("%s : %v", "AccountHandler/HandlerAccountDelete/Store", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
	return
}

func (a *AccountHandler) HandlerAccountTransfer(ctx *gin.Context) {
	fromAccountNumber, err := strconv.Atoi(ctx.Param("from_account_number"))
	if err != nil {
		a.logger.Errorf("%s : %v", "AccountHandler/HandlerAccountTransfer/parseFromAccountNumber", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var param domain.TransferParam
	err = ctx.Bind(&param)
	if err != nil {
		a.logger.Errorf("%s : %v", "AccountHandler/HandlerAccountTransfer/ParseBodyData", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = a.accountUseCase.Transfer(ctx, fromAccountNumber, param)
	if err != nil {
		a.logger.Errorf("%s : %v", "AccountHandler/HandlerAccountTransfer/Transfer", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (a *AccountHandler) HandlerLogin(ctx *gin.Context) {
	var param domain.AccountLoginParam
	err := ctx.Bind(&param)
	if err != nil {
		a.logger.Errorf("%s : %v", "AccountHandler/HandlerLogin/Bind", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := a.accountUseCase.Login(ctx, param)
	if err != nil {
		a.logger.Errorf("%s : %v", "AccountHandler/HandlerLogin/Login", err)
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			ctx.JSON(http.StatusBadRequest, util.Response{
				Errors: []string{"Email not found"},
			})
		} else if strings.Contains(err.Error(), "hashedPassword is not the hash of the given password") {
			ctx.JSON(http.StatusUnauthorized, util.Response{
				Errors: []string{"Invalid Password"},
			})
		}
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, response)
}
