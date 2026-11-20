-- name: CreateLedgerDebit :one
INSERT INTO ledger (transfer_id, account_id, money_in, money_out)
VALUES (@transfer_id, @account_id, 0, @amount)
    RETURNING *;

-- name: CreateLedgerCredit :one
INSERT INTO ledger (transfer_id, account_id, money_in, money_out)
VALUES (@transfer_id, @account_id, @amount, 0)
    RETURNING *;
