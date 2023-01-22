package database

import (
	"context"

	pgx "github.com/jackc/pgx/v5"
	"github.com/jannotti-glaucio/timescale-assignment/internal/excepts"
)

func OpenConnection(ctx context.Context, url string) (*pgx.Conn, *excepts.Exception) {

	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		return nil, excepts.ThrowException(excepts.ErrorConnectingToDB, "Unable to connect to database: %v", err)
	}
	return conn, nil
}

func CloseConnection(ctx context.Context, conn *pgx.Conn) {
	conn.Close(ctx)
}

func QueryRow(ctx context.Context, conn *pgx.Conn, sql string, args ...any) pgx.Row {
	return conn.QueryRow(ctx, sql, args...)
}
