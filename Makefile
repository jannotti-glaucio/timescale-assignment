build:
	docker-compose build

tests:
	go test -short -race -v ./...

integration:
	go test -run Integration -v ./...

coverage:
	go test -cover -p=1 -covermode=count -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

app:
	go run ./cmd/main.go

dependencies:
	docker-compose up -d timescaledb flyway

docker:
	docker-compose up --build
