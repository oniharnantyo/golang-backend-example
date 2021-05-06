-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS account (
    account_number      SERIAL NOT NULL,
    customer_number     INT NOT NULL REFERENCES customer(customer_number),
    balance           INT NOT NULL,
    PRIMARY KEY(customer_number)
);

ALTER TABLE account ADD COLUMN email varchar(255) NOT NULL DEFAULT '';
ALTER TABLE account ADD COLUMN password varchar(255) NOT NULL DEFAULT '';
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE account;