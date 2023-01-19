build:
	docker-compose build

docker:
	docker-compose up --build

tests:
	go test -race -covermode=atomic -count=1 ./...

run:
	go run ./cmd/main.go
