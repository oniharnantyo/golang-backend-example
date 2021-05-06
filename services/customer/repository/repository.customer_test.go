package repository_customer

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/oniharnantyo/golang-backend-example/domain"
	"github.com/oniharnantyo/golang-backend-example/util"

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

func TestCustomerRepository_List(t *testing.T) {
	db, mock := initMock()

	defer db.Close()

	rows := sqlmock.NewRows([]string{"customer_number", "name"}).
		AddRow(1001, "Bob Martin").
		AddRow(1002, "Linus Torvalds")

	search := "bob"
	order := "ASC"
	limit := 10
	offset := 0

	query := fmt.Sprintf(`
		SELECT 
			customer_number, 
			name 
		FROM customer 
		WHERE 
			(LOWER(customer_number) LIKE '%%bob%%' OR LOWER(name) LIKE '%%bob%%')
		ORDER BY name ASC 
		LIMIT $1 OFFSET $2`)

	prep := mock.ExpectPrepare(query)

	prep.ExpectQuery().WithArgs(limit, offset).WillReturnRows(rows)

	c := NewCustomerRepository(db)

	customers, err := c.List(context.Background(), domain.CustomerListParam{util.Filter{
		Limit:  limit,
		Offset: offset,
		Search: search,
		Order:  order,
	}})
	assert.NoError(t, err)
	assert.NotNil(t, customers)
	assert.Len(t, customers, 2)
}

func TestCustomerRepository_GetByCustomerNumber(t *testing.T) {
	db, mock := initMock()

	defer db.Close()

	rows := sqlmock.NewRows([]string{"customer_number", "name"}).
		AddRow(1001, "Bob Martin")

	query := fmt.Sprintf(`
		SELECT 
			customer_number, 
			name 
		FROM customer 
		WHERE customer_number = $1
	`)

	prep := mock.ExpectPrepare(query)

	customerNumber := 1001
	prep.ExpectQuery().WithArgs(customerNumber).WillReturnRows(rows)

	c := NewCustomerRepository(db)

	customers, err := c.GetByCustomerNumber(context.Background(), customerNumber)
	assert.NoError(t, err)
	assert.NotNil(t, customers)
}

func TestCustomerRepository_Store(t *testing.T) {
	db, mock := initMock()

	defer db.Close()

	query := fmt.Sprintf(`
		INSERT INTO customer (
			customer_number,
			name
		) VALUES (
			$1, $2
		)`)

	prep := mock.ExpectPrepare(query)

	customerNumber := 1001
	name := "Bob Martin"
	prep.ExpectExec().WithArgs(customerNumber, name).
		WillReturnResult(sqlmock.NewResult(1, 1))

	c := NewCustomerRepository(db)

	err := c.Store(context.Background(), &domain.Customer{
		CustomerNumber: customerNumber,
		Name:           name,
	})

	assert.NoError(t, err)
}

func TestCustomerRepository_Update(t *testing.T) {
	db, mock := initMock()

	defer db.Close()

	query := fmt.Sprintf(`
		UPDATE customer SET
			name = $1
		WHERE
			customer_number = $2`)

	prep := mock.ExpectPrepare(query)

	customerNumber := 1001
	name := "Bob Martin"
	prep.ExpectExec().WithArgs(name, customerNumber).
		WillReturnResult(sqlmock.NewResult(1, 1))

	c := NewCustomerRepository(db)

	err := c.Update(context.Background(), &domain.Customer{
		CustomerNumber: customerNumber,
		Name:           name,
	})

	assert.NoError(t, err)
}

func TestCustomerRepository_Delete(t *testing.T) {
	db, mock := initMock()

	defer db.Close()

	query := fmt.Sprintf(`
		DELETE FROM customer
		WHERE
			customer_number = $1`)

	prep := mock.ExpectPrepare(query)

	customerNumber := 1001
	prep.ExpectExec().WithArgs(customerNumber).
		WillReturnResult(sqlmock.NewResult(1, 1))

	c := NewCustomerRepository(db)

	err := c.Delete(context.Background(), &domain.Customer{
		CustomerNumber: customerNumber,
	})

	assert.NoError(t, err)
}
