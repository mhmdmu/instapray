# 0003: The ledger is the source of truth

Status: accepted

## Context
A balance is a summary: if it is wrong, nothing tells us why.
We need a record that proves where money went.

## Decision
- Every transfer writes two ledger entries: one with an Out amount
  for the sender and one with an In amount for the receiver.
- An entry has exactly one positive amount; the other is zero.
- Entries are never changed or deleted.
- Accounts also store a balance. If it disagrees with the ledger,
  the ledger is right.

## Alternatives
- One signed amount per entry: simpler, but we preferred two columns because they read like a real ledger.

## Consequences
We need tests that check balance = sum of entries, and that Out equals In for every transfer.
