# 0006: How money enters the system

Status: accepted

## Context
Accounts start at 0 and can't go negative, so nobody can send money first.
In the real InstaPay, accounts live in banks and already hold money when a user links them.

## Decision
In this stage, tests seed accounts with a starting balance. The starting balance has no ledger entry,
so the invariant is: balance = starting balance + sum of In - sum of Out.

## Consequences
- Opening balances are not in the ledger.
- Revisit later, when the mock banks hold the balances.
