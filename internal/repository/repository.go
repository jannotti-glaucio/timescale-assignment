package repository

import (
	"context"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jannotti-glaucio/timescale-assignment/internal/database"
	"github.com/jannotti-glaucio/timescale-assignment/internal/excepts"
	"github.com/jannotti-glaucio/timescale-assignment/internal/logger"
)

const query = "select max(usage), min(usage) from cpu_usage cu where host = $1 and ts between $2 and $3"

func OpenConnection(ctx context.Context) (*pgx.Conn, *excepts.Exception) {
	url := os.Getenv("DB_URL")
	return database.OpenConnection(ctx, url)
}

func CloseConnection(ctx context.Context, conn *pgx.Conn) {
	database.CloseConnection(ctx, conn)
}

func RunQuery(ctx context.Context, conn *pgx.Conn, hostname string, startDate time.Time, endDate time.Time) (*float32, *float32, *excepts.Exception) {

	logger.Debug("Executing query for hostname [%s], startDate: [%v], endData: [%v]", hostname, startDate, endDate)

	var maxUsage float32
	var minUsage float32
	row := database.QueryRow(ctx, conn, query, hostname, startDate, endDate)

	err := row.Scan(&maxUsage, &minUsage)
	if err != nil {
		return nil, nil, excepts.ThrowException(excepts.ErrorExecutingQuery, "Error scaning fields from query: %v", err)
	}

	return &maxUsage, &minUsage, nil
}
