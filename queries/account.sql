-- name: CreateAccount :one
INSERT INTO account(balance) VALUES($1)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM account WHERE id = $1;

-- name: ListAccounts :many
SELECT * FROM account;
