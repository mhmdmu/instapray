-- name: CreateTransfer :one
INSERT INTO transfer (sender_id, receiver_id, amount)
VALUES (@sender_id, @receiver_id, @amount)
    RETURNING *;
