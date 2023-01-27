package repository

import (
	"database/sql"
	"time"

	"github.com/jannotti-glaucio/timescale-assignment/internal/excepts"
	"github.com/jannotti-glaucio/timescale-assignment/internal/logger"
)

const query = "select max(usage), min(usage) from cpu_usage cu where host = $1 and ts between $2 and $3"

type (
	Repository interface {
		RunQuery(hostname string, startDate time.Time, endDate time.Time) (*float32, *float32, error)
	}

	repository struct {
		db *sql.DB
	}
)

func NewRepository(db *sql.DB) Repository {
	return repository{db: db}
}

func (r repository) RunQuery(hostname string, startDate time.Time, endDate time.Time) (*float32, *float32, error) {

	logger.Debug("Executing query for hostname [%s], startDate: [%v], endDate: [%v]", hostname, startDate, endDate)

	var maxUsage float32
	var minUsage float32
	row := r.db.QueryRow(query, hostname, startDate, endDate)

	err := row.Scan(&maxUsage, &minUsage)
	if err != nil {
		return nil, nil, excepts.ThrowException(excepts.ErrorExecutingQuery, "Error scaning fields from query: %v", err)
	}

	return &maxUsage, &minUsage, nil
}
