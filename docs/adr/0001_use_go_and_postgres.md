# 0001: Use Go and PostgreSQL

**Status:** Accepted

## Context

Inspired by real-world instant payment systems like ***InstaPay***, we want to understand how to build a reliable money-transfer backend.

- What if a user submits the same transfer twice?
- How do we ensure the sender is debited and the recipient credited atomically?
- What happens if a failure occurs midway through a transfer?
- How do we prevent concurrent transfers from overspending an account?
- How do we maintain data integrity and performance under heavy load?

These questions guide our exploration of transactions, concurrency control, idempotency, and failure handling.

## Decision

- **Go** for application code: fast, modern, and lets us explore concurrency and transaction handling with minimal abstraction.
- **PostgreSQL** for persistence: an enterprise-grade database with strong transactional guarantees, row-level locking, and isolation levels.

## Alternatives

- **Java/Spring**: A solid choice, but introduces more framework concepts and abstractions than needed for this project's learning goals.

## Consequences

- Write SQL explicitly and use `sqlc` to generate Go code, keeping transactions and locking visible.
- Use Go goroutines to build realistic concurrency tests.
