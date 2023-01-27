//go:build integration

package main

import (
	"github.com/jannotti-glaucio/timescale-assignment/internal/excepts"

	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessIntegration(t *testing.T) {
	t.Run("sucessfull", func(t *testing.T) {
		os.Setenv("FILE_PATH", "../test/data/query_params.csv")
		os.Setenv("DB_URL", "postgres://homework:abc123@localhost:5434/homework")

		totalProcessingTime, summarizeResult, err := process()

		assert.NotNil(t, totalProcessingTime)
		assert.NotNil(t, summarizeResult)
		assert.Nil(t, err)
		assert.Equal(t, 200, summarizeResult.NumberOfQueries)
	})
	t.Run("missing variables", func(t *testing.T) {
		os.Setenv("FILE_PATH", "")
		os.Setenv("DB_URL", "")

		totalProcessingTime, summarizeResult, err := process()

		assert.Nil(t, totalProcessingTime)
		assert.Nil(t, summarizeResult)
		assert.NotNil(t, err)

		exception := excepts.FromError(err)
		assert.Equal(t, excepts.MissingEnvVariable, exception.Code)
	})
	t.Run("database connection error", func(t *testing.T) {
		os.Setenv("FILE_PATH", "../test/data/query_params.csv")
		os.Setenv("DB_URL", "postgres://homework:XXX@localhost:5434/homework")

		totalProcessingTime, summarizeResult, err := process()

		assert.Nil(t, totalProcessingTime)
		assert.Nil(t, summarizeResult)
		assert.NotNil(t, err)

		exception := excepts.FromError(err)
		assert.Equal(t, excepts.ErrorConnectingToDB, exception.Code)
	})
}
