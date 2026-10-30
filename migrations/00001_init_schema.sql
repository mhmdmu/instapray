-- +goose Up
CREATE TABLE account (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    balance BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- Constraints
    CONSTRAINT chk_balance_non_negative CHECK (balance >= 0)
);

CREATE TABLE transfer (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    sender_id BIGINT NOT NULL,
    receiver_id BIGINT NOT NULL,
    amount BIGINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- Constraints
    CONSTRAINT chk_amount_positive CHECK (amount > 0),
    CONSTRAINT chk_different_accounts CHECK (sender_id <> receiver_id),
    CONSTRAINT fk_transfer_sender FOREIGN KEY (sender_id) REFERENCES account(id),
    CONSTRAINT fk_transfer_receiver FOREIGN KEY (receiver_id) REFERENCES account(id)
);

CREATE TABLE ledger (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    transfer_id BIGINT NOT NULL,
    account_id BIGINT NOT NULL,
    money_in BIGINT NOT NULL DEFAULT 0,
    money_out BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- Constraints
    CONSTRAINT chk_money_in_out
        CHECK (
            (money_in >= 0 AND money_out = 0)
            OR
            (money_out >= 0 AND money_in = 0)
        ),

    CONSTRAINT fk_ledger_transfer FOREIGN KEY (transfer_id) REFERENCES transfer(id),
    CONSTRAINT fk_ledger_account FOREIGN KEY (account_id) REFERENCES account(id)
);

-- +goose Down
DROP TABLE ledger;
DROP TABLE transfer;
DROP TABLE account;
