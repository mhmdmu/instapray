# InstaPray 🙏

An InstaPay-inspired payment backend, built as a sandbox for learning how to move money between accounts without losing or duplicating it.

> No real money. No real banks. No real customer data. **Learning project only.**

Built with Go and PostgreSQL. See [`docs/adr`](docs/adr) for why.

## Requirements

- Go (version pinned in `go.mod`)
- Docker with Compose
- [goose](https://github.com/pressly/goose) for database migrations

Install goose once **version 3.29.0**.

```bash
go install github.com/pressly/goose/v3/cmd/goose@v3.29.0
```

> Make sure `~/go/bin` (Windows: `%USERPROFILE%\go\bin`) is on your PATH. Check with `goose -version`.

## Run it locally

1. Clone the repo:

   ```bash
   git clone https://github.com/mhmdmu/instapray.git
   cd instapray
   ```

2. Create your `.env` file. The defaults work as they are:

   ```bash
   cp .env.example .env        # Windows cmd: copy .env.example .env
   ```

3. Start the database and wait until it shows `healthy`:

   ```bash
   docker compose up -d
   ```

4. Create the tables. Run this from the repo root, because goose reads `.env` from the current folder:

   ```bash
   goose up
   ```

5. Run the app:

   ```bash
   go run .
   ```

To check that the database and goose are connected, run `goose status`.

## Stop the database

```bash
docker compose down         # add -v to wipe its data
```
