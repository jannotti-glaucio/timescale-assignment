build:
	docker-compose build

linter:
	golangci-lint run

tests:
	go test -short -race ./...

integration:
	go test --tags=integration ./cmd ./internal/database

benchmark:
	go test -bench=. ./...

coverage:
	go test -cover -p=1 -covermode=count -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

app-run:
	go run ./cmd/main.go

docker-run:
	docker-compose up --build

docker-dependencies:
	docker-compose up -d timescaledb flyway

docker-app:
	docker-compose build
	docker-compose up app
