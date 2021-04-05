-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO account (account_number, customer_number, balance) VALUES (555001, 1001, 10000);
INSERT INTO account (account_number, customer_number, balance) VALUES (555002, 1002, 15000);
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM account WHERE account_number = 555001;
DELETE FROM account WHERE account_number = 555002;