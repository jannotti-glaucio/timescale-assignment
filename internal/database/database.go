package database

import (
	"github.com/jannotti-glaucio/timescale-assignment/internal/excepts"

	"database/sql"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func OpenConnection() (*sql.DB, error) {

	url := os.Getenv("DB_URL")

	db, openErr := sql.Open("pgx", url)
	if openErr != nil {
		return nil, excepts.ThrowException(excepts.ErrorConnectingToDB, "Unable to connect to database: %v", openErr)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, excepts.ThrowException(excepts.ErrorConnectingToDB, "Unable to ping the database connection: %v", pingErr)
	}

	return db, nil
}

func CloseConnection(db *sql.DB) error {
	if db == nil {
		return nil
	}

	err := db.Close()
	if err != nil {
		return excepts.ThrowException(excepts.ErrorDisconectingFromDB, "Unable to close database connection: %v", err)
	}

	return nil
}
