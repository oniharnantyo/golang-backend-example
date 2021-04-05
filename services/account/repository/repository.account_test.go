package repository_account

import (
	"context"
	"database/sql"
	"fmt"
	"linkaja-test/domain"
	"linkaja-test/util"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
)

func initMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestAccountRepository_List(t *testing.T) {
	db, mock := initMock()

	defer db.Close()

	rows := sqlmock.NewRows([]string{"account_number", "customer_number", "balance"}).
		AddRow(555001, 1001, 10000).
		AddRow(555002, 1002, 15000)

	search := "1"
	order := "ASC"
	limit := 10
	offset := 0

	query := fmt.Sprintf(`
		SELECT
			account_number, 
			customer_number,
			balance 
		FROM account 
		WHERE 
			(LOWER(account_number) LIKE '%%1%%' OR LOWER(customer_number) LIKE '%%1%%') 
		ORDER BY account_number ASC 
		LIMIT $1 OFFSET $2`)

	prep := mock.ExpectPrepare(query)

	prep.ExpectQuery().WithArgs(limit, offset).WillReturnRows(rows)

	c := NewAccountRepository(db)

	customers, err := c.List(context.Background(), domain.AccountListParam{
		util.Filter{
			Limit:  limit,
			Offset: offset,
			Search: search,
			Order:  order,
		}})
	assert.NoError(t, err)
	assert.NotNil(t, customers)
	assert.Len(t, customers, 2)
}

func TestAccountRepository_GetByAccountNumber(t *testing.T) {
	db, mock := initMock()

	defer db.Close()

	rows := sqlmock.NewRows([]string{"account_number", "customer_number", "balance"}).
		AddRow(555001, 1001, 10000)

	query := fmt.Sprintf(`
		SELECT
			account_number,
			customer_number,
			balance
		FROM account
		WHERE
			account_number = $1
	`)

	prep := mock.ExpectPrepare(query)

	accountNumber := 555001
	prep.ExpectQuery().WithArgs(accountNumber).WillReturnRows(rows)

	c := NewAccountRepository(db)

	customers, err := c.GetByAccountNumber(context.Background(), accountNumber)
	assert.NoError(t, err)
	assert.NotNil(t, customers)
}

func TestAccountRepository_Store(t *testing.T) {
	db, mock := initMock()

	defer db.Close()

	query := fmt.Sprintf(`
		INSERT INTO account (
			account_number,
			customer_number,
			balance
		) VALUES (
			$1, $2, $3
		)`)

	prep := mock.ExpectPrepare(query)

	accountNumber := 555001
	customerNumber := 1001
	balance := 10000
	prep.ExpectExec().WithArgs(accountNumber, customerNumber, balance).
		WillReturnResult(sqlmock.NewResult(1, 1))

	c := NewAccountRepository(db)

	err := c.Store(context.Background(), &domain.Account{
		AccountNumber:  accountNumber,
		CustomerNumber: customerNumber,
		Balance:        balance,
	})

	assert.NoError(t, err)
}

func TestAccountRepository_Update(t *testing.T) {
	db, mock := initMock()

	defer db.Close()

	query := fmt.Sprintf(`
		UPDATE account SET
			customer_number = $1,
			balance = $2
		WHERE
			account_number = $3`)

	prep := mock.ExpectPrepare(query)

	accountNumber := 555001
	customerNumber := 1001
	balance := 10000
	prep.ExpectExec().WithArgs(customerNumber, balance, accountNumber).
		WillReturnResult(sqlmock.NewResult(1, 1))

	c := NewAccountRepository(db)

	err := c.Update(context.Background(), &domain.Account{
		AccountNumber:  accountNumber,
		CustomerNumber: customerNumber,
		Balance:        balance,
	})

	assert.NoError(t, err)
}

func TestAccountRepository_Delete(t *testing.T) {
	db, mock := initMock()

	defer db.Close()

	query := fmt.Sprintf(`
		DELETE FROM account
		WHERE
			account_number = $1`)

	prep := mock.ExpectPrepare(query)

	customerNumber := 555001
	prep.ExpectExec().WithArgs(customerNumber).
		WillReturnResult(sqlmock.NewResult(1, 1))

	c := NewAccountRepository(db)

	err := c.Delete(context.Background(), &domain.Account{
		AccountNumber: customerNumber,
	})

	assert.NoError(t, err)
}
