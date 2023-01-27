build:
	docker-compose build

linter:
	golangci-lint run

tests:
	go test -short -race ./...

integration:
	go test --tags=integration ./cmd ./internal/database

coverage:
	go test -cover -p=1 -covermode=count -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

app:
	go run ./cmd/main.go

dependencies:
	docker-compose up -d timescaledb flyway

docker:
	docker-compose up --build
