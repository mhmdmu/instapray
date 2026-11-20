-- name: CreateAccount :one
INSERT INTO account(balance) VALUES($1)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM account WHERE id = $1;

-- name: ListAccounts :many
SELECT * FROM account;

-- name: GetAccountForUpdate :one
SELECT * FROM account
WHERE id = @id
    FOR NO KEY UPDATE;

-- name: DebitAccount :one
UPDATE account
SET balance = balance - @amount
WHERE id = @id AND balance >= @amount
    RETURNING *;

-- name: CreditAccount :one
UPDATE account
SET balance = balance + @amount
WHERE id = @id
    RETURNING *;
