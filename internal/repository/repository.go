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
		RunQuery(hostname string, startDate time.Time, endDate time.Time) (float64, float64, error)
	}

	repository struct {
		db *sql.DB
	}
)

func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r *repository) RunQuery(hostname string, startDate time.Time, endDate time.Time) (float64, float64, error) {

	logger.Debug("Executing query for hostname [%s], startDate: [%v], endDate: [%v]", hostname, startDate, endDate)

	var maxUsage sql.NullFloat64
	var minUsage sql.NullFloat64
	row := r.db.QueryRow(query, hostname, startDate, endDate)

	err := row.Scan(&maxUsage, &minUsage)
	if err != nil {
		return 0, 0, excepts.ThrowException(excepts.ErrorExecutingQuery, "Error scaning fields from query: %v", err)
	}

	if maxUsage.Valid && minUsage.Valid {
		return maxUsage.Float64, minUsage.Float64, nil
	} else {
		return 0, 0, nil
	}
}
