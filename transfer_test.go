package main

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mhmdmu/instapray/db"
)

func countRows(t *testing.T, pool *pgxpool.Pool, query string, arg int64) int {
	t.Helper()
	var n int
	if err := pool.QueryRow(context.Background(), query, arg).Scan(&n); err != nil {
		t.Fatalf("count: %v", err)
	}
	return n
}

func TestTransferTx_HappyPath(t *testing.T) {
	ctx := context.Background()
	pool := newTestDB(t)
	store := db.NewStore(pool)

	from := seedAccount(t, pool, 1000)
	to := seedAccount(t, pool, 500)

	res, err := store.TransferTx(ctx, db.TransferTxParams{
		SenderID: from, ReceiverID: to, Amount: 200,
	})
	if err != nil {
		t.Fatalf("TransferTx: %v", err)
	}

	// Re-read from the DB, don't trust only the returned structs
	gotFrom, err := store.GetAccount(ctx, from)
	if err != nil {
		t.Fatalf("GetAccount sender: %v", err)
	}
	gotTo, err := store.GetAccount(ctx, to)
	if err != nil {
		t.Fatalf("GetAccount receiver: %v", err)
	}
	if gotFrom.Balance != 800 {
		t.Errorf("sender balance = %d, want 800", gotFrom.Balance)
	}
	if gotTo.Balance != 700 {
		t.Errorf("receiver balance = %d, want 700", gotTo.Balance)
	}

	// 1 transfer row, 2 ledger entries
	if n := countRows(t, pool, `SELECT count(*) FROM transfer WHERE id = $1`, res.Transfer.ID); n != 1 {
		t.Errorf("transfer rows = %d, want 1", n)
	}
	if n := countRows(t, pool, `SELECT count(*) FROM ledger WHERE transfer_id = $1`, res.Transfer.ID); n != 2 {
		t.Errorf("ledger rows = %d, want 2", n)
	}
}

func TestTransferTx_RollsBackOnInsufficientFunds(t *testing.T) {
	ctx := context.Background()
	pool := newTestDB(t)
	store := db.NewStore(pool)

	from := seedAccount(t, pool, 100)
	to := seedAccount(t, pool, 500)

	_, err := store.TransferTx(ctx, db.TransferTxParams{
		SenderID: from, ReceiverID: to, Amount: 200, // more than the sender has
	})
	if !errors.Is(err, db.ErrInsufficientFunds) {
		t.Fatalf("err = %v, want ErrInsufficientFunds", err)
	}

	// Nothing should have changed
	gotFrom, _ := store.GetAccount(ctx, from)
	gotTo, _ := store.GetAccount(ctx, to)
	if gotFrom.Balance != 100 || gotTo.Balance != 500 {
		t.Errorf("balances = %d/%d, want 100/500", gotFrom.Balance, gotTo.Balance)
	}
	if n := countRows(t, pool, `SELECT count(*) FROM transfer WHERE sender_id = $1`, from); n != 0 {
		t.Errorf("transfer rows = %d, want 0 (rollback failed)", n)
	}
	if n := countRows(t, pool, `SELECT count(*) FROM ledger WHERE account_id = $1`, from); n != 0 {
		t.Errorf("ledger rows = %d, want 0 (rollback failed)", n)
	}
}
