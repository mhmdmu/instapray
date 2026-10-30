# 0004: Enforce rules in the database when one row is enough

Status: accepted

## Context
A bug in Go code must not be able to create an impossible state.

## Decision
If one row is enough to decide a rule, the database enforces it (constraints, foreign keys).
Rules that need other rows (totals, counts) are enforced in Go code inside the transaction and covered by invariant tests.

## Alternatives
Checking everything in Go only: simpler, but one bug can corrupt data.

## Consequences
- Database: balance >= 0, amount > 0, sender != receiver, Out/In rule, foreign keys, ledger entries can't be changed or deleted.
- Go and tests: balance matches the ledger, Out equals In, two entries per transfer.
