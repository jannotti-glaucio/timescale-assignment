package database

import (
	"github.com/jannotti-glaucio/timescale-assignment/internal/logger"

	"context"

	pgx "github.com/jackc/pgx/v5"
)

func OpenConnection(url string) *pgx.Conn {

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		logger.Fatal("Unable to connect to database: %v", err)
	}
	return conn
}

func CloseConnection(conn *pgx.Conn) {
	conn.Close(context.Background())
}

func QueryRow(conn *pgx.Conn, sql string, args ...any) pgx.Row {
	return conn.QueryRow(context.Background(), sql, args...)
}
