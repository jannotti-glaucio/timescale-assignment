//go:build !integration

package repository

import (
	"github.com/jannotti-glaucio/timescale-assignment/internal/tests"

	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestRunQuery(t *testing.T) {

	db, mock := tests.NewSQLMock()
	defer db.Close()

	startDate, _ := time.Parse(time.RFC3339, "2017-01-01 09:59:22")
	endDate, _ := time.Parse(time.RFC3339, "2017-01-01 09:59:22")
	expectedMaxUsage := float32(200)
	expectedMinUsage := float32(100)

	mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs("host-01", startDate, endDate).
		WillReturnRows(sqlmock.NewRows([]string{"max", "min"}).AddRow(expectedMaxUsage, expectedMinUsage))

	repository := NewRepository(db)
	maxUsage, minUsage, err := repository.RunQuery("host-01", startDate, endDate)

	assert.Nil(t, err)
	assert.Equal(t, expectedMaxUsage, *maxUsage)
	assert.Equal(t, expectedMinUsage, *minUsage)
}
