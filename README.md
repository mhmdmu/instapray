# InstaPray 🙏

An InstaPay-inspired payment backend, built as a sandbox for learning how to move money between accounts without losing or duplicating it.

> No real money. No real banks. No real customer data. **Learning project only.**

## Requirements

- Go (version pinned in `go.mod`)
- Docker with Compose

## Run it locally

```bash
git clone https://github.com/mhmdmu/instapray.git
cd instapray
cp .env.example .env        # Windows cmd: copy .env.example .env
docker compose up -d
docker compose ps           # wait until the db shows "healthy"
go run .
```

> Stop the database: `docker compose down` (add `-v` to wipe its data).
