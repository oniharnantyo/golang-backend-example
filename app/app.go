package app

import (
	"context"
	"database/sql"
	"fmt"
	"linkaja-test/database"
	"linkaja-test/database/migration"
	"linkaja-test/domain"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	delivery_http_account "linkaja-test/services/account/delivery/http"
	repository_account "linkaja-test/services/account/repository"
	usecase_account "linkaja-test/services/account/usecase"
	delivery_http_customer "linkaja-test/services/customer/delivery/http"
	repository_customer "linkaja-test/services/customer/repository"
	usecase_customer "linkaja-test/services/customer/usecase"
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

	r := mux.NewRouter()

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

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		if err := srv.Shutdown(ctx); err != nil {
			logger.Printf("Server shutdown error : %v", err)
		}

		logger.Println("Server shutdown gracefully")
		close(idleConnsClosed)
	}()

	logger.Println(fmt.Sprintf("Server Listen And Serve on Port : %d", viper.GetInt("app.port")))
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		logger.Println(fmt.Sprintf("err on listen and serve : %v", err))
		os.Exit(0)
	}

	<-idleConnsClosed
}
