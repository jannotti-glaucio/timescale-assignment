build:
	docker-compose build

linter:
	golangci-lint run

tests:
	go test -short -race ./...

coverage:
	go test -cover -p=1 -covermode=count -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

timescaledb:
	docker-compose up -d timescaledb

migrations:
	docker-compose up flyway

docker-app:
	docker-compose build
	docker-compose up app

app:
	go run ./cmd/main.go

integration:
	go test --tags=integration ./cmd ./internal/database

benchmark:
	go test -bench=. ./...
