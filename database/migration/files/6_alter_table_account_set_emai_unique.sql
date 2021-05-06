-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE account ADD CONSTRAINT email_unique UNIQUE(email);
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE account;