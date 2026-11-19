-- name: GetAccountForUpdate :one
SELECT * FROM account
WHERE id = @id
    FOR NO KEY UPDATE;

-- name: CreateTransfer :one
INSERT INTO transfer (sender_id, receiver_id, amount)
VALUES (@sender_id, @receiver_id, @amount)
    RETURNING *;

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

-- name: CreateLedgerDebit :one
INSERT INTO ledger (transfer_id, account_id, money_in, money_out)
VALUES (@transfer_id, @account_id, 0, @amount)
    RETURNING *;

-- name: CreateLedgerCredit :one
INSERT INTO ledger (transfer_id, account_id, money_in, money_out)
VALUES (@transfer_id, @account_id, @amount, 0)
    RETURNING *;