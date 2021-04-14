package app

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"golang-backend-example/database"
	"golang-backend-example/database/migration"
	"golang-backend-example/domain"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	delivery_http_account "golang-backend-example/services/account/delivery/http"
	repository_account "golang-backend-example/services/account/repository"
	usecase_account "golang-backend-example/services/account/usecase"
	delivery_http_customer "golang-backend-example/services/customer/delivery/http"
	repository_customer "golang-backend-example/services/customer/repository"
	usecase_customer "golang-backend-example/services/customer/usecase"
)

func Run() {
	initConfig()

	logger := initLogger()

	dbPool, err := initDatabase()
	if err != nil {
		logger.Fatalf("%s: %v", "Error on connect to database", err)
	}

	accountUseCase, customerUseCase := initService(dbPool, logger)

	initHandler(accountUseCase, customerUseCase, logger)
}

func initConfig() {
	viper.SetConfigType("toml")

	viper.AddConfigPath(".")
	viper.SetConfigName(".config")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file: ", viper.ConfigFileUsed())
	} else {
		log.Fatal(err)
	}
}

func initLogger() *logrus.Logger {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}

	return logger
}

func initDatabase() (*sql.DB, error) {
	dbConn := database.DatabaseConnector{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		DBName:   viper.GetString("database.name"),
		SSLMode:  viper.GetString("database.sslmode"),
	}

	dbPool, err := dbConn.Connect()
	if err != nil {
		return nil, err
	}

	// Migrate database schema
	err = migration.Up(dbPool)
	if err != nil {
		log.Fatalf("Error on migrate schema : %v", err)
	}

	return dbPool, nil
}

func initService(dbPool *sql.DB, logger *logrus.Logger) (domain.AccountUseCase, domain.CustomerUseCase) {
	accountRepository := repository_account.NewAccountRepository(dbPool)
	customerRepository := repository_customer.NewCustomerRepository(dbPool)

	accountUseCase := usecase_account.NewAccountUseCase(accountRepository, customerRepository, logger)
	customerUseCase := usecase_customer.NewCustomerUseCase(customerRepository, logger)

	return accountUseCase, customerUseCase
}

func initHandler(accountUseCase domain.AccountUseCase, customerUseCase domain.CustomerUseCase, logger *logrus.Logger) {
	ctx := context.Background()

	r := gin.Default()

	http.Handle("/", r)

	delivery_http_account.NewAccountHandler(r, accountUseCase, logger)
	delivery_http_customer.NewCustomerHandler(r, customerUseCase, logger)

	srv := &http.Server{
		Addr:         fmt.Sprintf(`:%d`, viper.GetInt("app.port")),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
