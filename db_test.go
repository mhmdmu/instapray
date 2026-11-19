package main

import (
	"context"
	"database/sql"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib" // lets goose use pgx through database/sql
	"github.com/mhmdmu/instapray/db"
	"github.com/pressly/goose/v3"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

func newTestDB(t *testing.T) *pgxpool.Pool {
	t.Helper()
	ctx := context.Background()

	// Start Postgres 18
	ctr, err := postgres.Run(ctx, "postgres:18",
		postgres.WithDatabase("test"),
		postgres.WithUsername("test"),
		postgres.WithPassword("test"),
		postgres.BasicWaitStrategies(),
	)
	testcontainers.CleanupContainer(t, ctr) // stops it when the test ends
	if err != nil {
		t.Fatalf("start postgres: %v", err)
	}

	// Connection string with the random port
	url, err := ctr.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("connection string: %v", err)
	}

	// Run goose migrations (goose needs database/sql)
	sqlDB, err := sql.Open("pgx", url)
	if err != nil {
		t.Fatalf("open sql db: %v", err)
	}
	defer sqlDB.Close()
	if err := goose.SetDialect("postgres"); err != nil {
		t.Fatalf("goose dialect: %v", err)
	}
	if err := goose.Up(sqlDB, "migrations"); err != nil {
		t.Fatalf("goose up: %v", err)
	}

	// Pool for the test
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		t.Fatalf("pgx pool: %v", err)
	}
	t.Cleanup(pool.Close)
	return pool
}

// seedAccount creates an account with a starting balance (ADR 0006)
// and returns its id. Tests use it instead of writing SQL by hand.
func seedAccount(t *testing.T, pool *pgxpool.Pool, balance int64) int64 {
	t.Helper()

	acc, err := db.New(pool).CreateAccount(context.Background(), balance)
	if err != nil {
		t.Fatalf("seed account: %v", err)
	}
	return acc.ID
}

// Proves the whole chain works: container -> migrations -> sqlc -> real table.
func TestDBRoundTrip(t *testing.T) {
	pool := newTestDB(t)
	q := db.New(pool)

	id := seedAccount(t, pool, 1000)

	got, err := q.GetAccount(context.Background(), id)
	if err != nil {
		t.Fatalf("GetAccount: %v", err)
	}
	if got.ID != id {
		t.Errorf("id = %d, want %d", got.ID, id)
	}
	if got.Balance != 1000 {
		t.Errorf("balance = %d, want 1000", got.Balance)
	}
}
