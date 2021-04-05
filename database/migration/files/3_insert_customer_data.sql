-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO customer (customer_number, name) VALUES (1001, 'Bob Martin');
INSERT INTO customer (customer_number, name) VALUES (1002, 'Linus Torvalds');
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM customer WHERE customer_number = 1001;
DELETE FROM customer WHERE customer_number = 1002;