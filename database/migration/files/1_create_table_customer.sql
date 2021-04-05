-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS customer (
    customer_number     SERIAL NOT NULL,
    name                VARCHAR(255) NOT NULL,
    PRIMARY KEY(customer_number)
);
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE customer;