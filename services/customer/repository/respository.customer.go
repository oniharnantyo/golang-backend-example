package repository_customer

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/oniharnantyo/golang-backend-example/domain"
	"github.com/oniharnantyo/golang-backend-example/util"
)

type customerRepository struct {
	dbPool *sql.DB
}

func (c customerRepository) List(ctx context.Context, param domain.CustomerListParam) ([]domain.Customer, error) {
	var filters []string

	if param.Search != "" {
		filters = append(filters,
			fmt.Sprintf(`(LOWER(customer_number) LIKE '%%%s%%' OR LOWER(name) LIKE '%%%s%%')`,
				param.Search, param.Search))
	}

	filterQuery := util.BuildFilterQuery(filters)

	stmt, err := c.dbPool.Prepare(fmt.Sprintf(`
		SELECT
			customer_number,
			name
		FROM customer
			%s
		ORDER BY name %s
		LIMIT $1 OFFSET $2
	`, filterQuery, param.Order))
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, param.Limit, param.Offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var customers []domain.Customer
	for rows.Next() {
		var customer domain.Customer
		err := rows.Scan(
			&customer.CustomerNumber,
			&customer.Name,
		)
		if err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	return customers, nil
}

func (c customerRepository) GetByCustomerNumber(ctx context.Context, customerNumber int) (domain.Customer, error) {
	stmt, err := c.dbPool.Prepare(fmt.Sprintf(`
		SELECT
			customer_number,
			name
		FROM customer
		WHERE
			customer_number = $1
	`))
	if err != nil {
		return domain.Customer{}, err
	}

	var customer domain.Customer
	err = stmt.QueryRowContext(ctx, customerNumber).Scan(
		&customer.CustomerNumber,
		&customer.Name,
	)
	if err != nil {
		return domain.Customer{}, err
	}

	return customer, nil
}

func (c customerRepository) Store(ctx context.Context, a *domain.Customer) error {
	stmt, err := c.dbPool.Prepare(fmt.Sprintf(`
		INSERT INTO customer (
			customer_number,
			name
		) VALUES (
			$1, $2
		)`))
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx,
		a.CustomerNumber,
		a.Name)
	if err != nil {
		return err
	}

	return nil
}

func (c customerRepository) Update(ctx context.Context, a *domain.Customer) error {
	stmt, err := c.dbPool.Prepare(fmt.Sprintf(`
		UPDATE customer SET
			name = $1
		WHERE
			customer_number = $2
	`))

	_, err = stmt.ExecContext(ctx,
		a.Name,
		a.CustomerNumber)
	if err != nil {
		return err
	}

	return nil
}

func (c customerRepository) Delete(ctx context.Context, a *domain.Customer) error {

	stmt, err := c.dbPool.Prepare(fmt.Sprintf(`
		DELETE FROM customer
		WHERE
			customer_number = $1
	`))

	_, err = stmt.ExecContext(ctx,
		a.CustomerNumber)
	if err != nil {
		return err
	}

	return nil
}

func NewCustomerRepository(db *sql.DB) domain.CustomerRepository {
	return &customerRepository{
		dbPool: db,
	}
}
