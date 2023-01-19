## Build
FROM golang:1.19-alpine AS build
WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY /cmd ./cmd
COPY /internal ./internal

RUN go build -o ./timescale-assignment ./cmd/main.go

## Run
FROM alpine:3
WORKDIR /app

COPY --from=build /build/timescale-assignment ./

ENTRYPOINT ["/app/timescale-assignment"]