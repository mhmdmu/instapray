# 0002: Store money as integer piasters

Status: accepted

## Context
Floats give rounding errors (0.1 + 0.2 != 0.3). We must never gain or lose a piaster.

## Decision
All amounts are whole piasters (100 piasters = 1 EGP), stored as BIGINT. No floats or decimals in code or database.

## Alternatives
- float64: rejected, rounding errors.
- NUMERIC(12,2): exact, but easy to mix with floats in Go.

## Consequences
We convert to EGP only when displaying.
