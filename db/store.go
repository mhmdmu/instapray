package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrInvalidTransfer   = errors.New("invalid transfer")
	ErrInsufficientFunds = errors.New("insufficient funds")
)

type Store struct {
	*Queries
	pool *pgxpool.Pool
}

func NewStore(pool *pgxpool.Pool) *Store {
	return &Store{Queries: New(pool), pool: pool}
}

type TransferTxParams struct {
	SenderID   int64
	ReceiverID int64
	Amount     int64
}

type TransferTxResult struct {
	Transfer      Transfer
	Sender        Account
	Receiver      Account
	SenderEntry   Ledger
	ReceiverEntry Ledger
}

func (s *Store) TransferTx(ctx context.Context, p TransferTxParams) (TransferTxResult, error) {
	var r TransferTxResult
	if p.Amount <= 0 || p.SenderID == p.ReceiverID {
		return r, ErrInvalidTransfer
	}

	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return r, err
	}
	defer tx.Rollback(ctx) // no-op once committed

	q := s.Queries.WithTx(tx)

	// Lock both accounts in id order so A->B and B->A can't deadlock.
	lo, hi := p.SenderID, p.ReceiverID
	if lo > hi {
		lo, hi = hi, lo
	}
	if _, err = q.GetAccountForUpdate(ctx, lo); err != nil {
		return r, err
	}
	if _, err = q.GetAccountForUpdate(ctx, hi); err != nil {
		return r, err
	}

	if r.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
		SenderID: p.SenderID, ReceiverID: p.ReceiverID, Amount: p.Amount,
	}); err != nil {
		return r, err
	}

	if r.Sender, err = q.DebitAccount(ctx, DebitAccountParams{ID: p.SenderID, Amount: p.Amount}); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return r, ErrInsufficientFunds
		}
		return r, err
	}
	if r.Receiver, err = q.CreditAccount(ctx, CreditAccountParams{ID: p.ReceiverID, Amount: p.Amount}); err != nil {
		return r, err
	}

	if r.SenderEntry, err = q.CreateLedgerDebit(ctx, CreateLedgerDebitParams{
		TransferID: r.Transfer.ID, AccountID: p.SenderID, Amount: p.Amount,
	}); err != nil {
		return r, err
	}
	if r.ReceiverEntry, err = q.CreateLedgerCredit(ctx, CreateLedgerCreditParams{
		TransferID: r.Transfer.ID, AccountID: p.ReceiverID, Amount: p.Amount,
	}); err != nil {
		return r, err
	}

	if err = tx.Commit(ctx); err != nil {
		return TransferTxResult{}, err
	}
	return r, nil
}
