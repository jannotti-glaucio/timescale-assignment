package database

import (
	"github.com/jannotti-glaucio/timescale-assignment/internal/exceptions"

	"context"

	pgx "github.com/jackc/pgx/v5"
)

func OpenConnection(ctx context.Context, url string) (*pgx.Conn, error) {

	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		exceptions.ThrowError("Unable to connect to database: %v", err)
	}
	return conn, nil
}

func CloseConnection(ctx context.Context, conn *pgx.Conn) {
	conn.Close(ctx)
}

func QueryRow(ctx context.Context, conn *pgx.Conn, sql string, args ...any) pgx.Row {
	return conn.QueryRow(ctx, sql, args...)
}
