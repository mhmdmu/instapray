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
