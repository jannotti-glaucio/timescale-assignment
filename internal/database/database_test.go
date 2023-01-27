//go:build integration

package database

import (
	"github.com/jannotti-glaucio/timescale-assignment/internal/excepts"

	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenConnection(t *testing.T) {
	t.Run("sucessfull", func(t *testing.T) {
		os.Setenv("DB_URL", "postgres://homework:abc123@localhost:5434/homework")

		db, err := OpenConnection()

		assert.Nil(t, err)
		assert.NotNil(t, db)
	})

	t.Run("error", func(t *testing.T) {
		os.Setenv("DB_URL", "")

		db, err := OpenConnection()

		assert.NotNil(t, err)
		assert.Nil(t, db)

		exception := excepts.FromError(err)
		assert.Equal(t, excepts.ErrorConnectingToDB, exception.Code)
	})
}

func TestCloseConnection(t *testing.T) {
	os.Setenv("DB_URL", "postgres://homework:abc123@localhost:5434/homework")

	db, err := OpenConnection()

	assert.Nil(t, err)
	assert.NotNil(t, db)

	err = CloseConnection(db)

	assert.Nil(t, err)
}
