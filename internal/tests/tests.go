package tests

import (
	"github.com/jannotti-glaucio/timescale-assignment/internal/logger"

	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
)

const TestDataPath = "../../test/data/"

func NewSQLMock() (*sql.DB, sqlmock.Sqlmock) {

	db, mock, err := sqlmock.New()
	if err != nil {
		logger.Fatal("Error creating mocked database connection: %v", err)
	}

	return db, mock
}
