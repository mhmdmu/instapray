.PHONY: db-up db-down run fmt test vet build ci

db-up:
	docker compose up -d

db-down:
	docker compose down

run:
	go run .

fmt:
	@test -z "$$(gofmt -l .)" || \
		(echo "Run 'gofmt -w .' to fix formatting"; exit 1)

test:
	go test ./...

vet:
	go vet ./...

build:
	go build ./...

ci: fmt vet test build
