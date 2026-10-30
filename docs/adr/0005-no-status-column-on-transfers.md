# 0005: No status column on transfers in this stage

Status: accepted

## Context
Currently a transfer happens inside one database transaction, so nobody sees half-finished states. Rejected transfers are not stored yet.
So a stored transfer is always completed, and a status column would always hold the same value.

## Decision
The transfer table has no status column.

## Alternatives
- Keep a status column: rejected, because it never varies.
- Keep an append-only history of status changes: useful when states are visible, but not needed now.

## Consequences
In future, When the debit and credit happen in separate banks, a transfer can be stuck halfway. We will then add a status column, and a history
table if we need to know when each state happened, in a new migration.
