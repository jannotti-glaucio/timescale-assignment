//go:build !integration

package parsers

import (
	"bufio"
	"os"
	"testing"
	"time"

	"github.com/jannotti-glaucio/timescale-assignment/internal/excepts"
	"github.com/jannotti-glaucio/timescale-assignment/internal/logger"
	"github.com/jannotti-glaucio/timescale-assignment/internal/model"
	"github.com/jannotti-glaucio/timescale-assignment/internal/tests"
	"github.com/stretchr/testify/assert"
)

const testFilesPath = tests.TestDataPath + "parsers/"

func TestParseFile(t *testing.T) {
	t.Run("valid file", func(t *testing.T) {
		requestsByHost, err := ParseFile(testFilesPath + "/valid_file.csv")

		assert.Nil(t, err)
		assert.Equal(t, 2, len(requestsByHost))

		requestHost1 := requestsByHost["host_000001"]
		assert.Equal(t, 2, len(requestHost1))

		expectedLine1StartDate, _ := time.Parse(timeFormat, "2017-01-01 08:59:22")
		expectedLine1EndDate, _ := time.Parse(timeFormat, "2017-01-01 09:59:22")
		expectedLine2StartDate, _ := time.Parse(timeFormat, "2017-01-02 13:02:02")
		expectedLine2EndDate, _ := time.Parse(timeFormat, "2017-01-02 14:02:02")
		assert.Equal(t, expectedLine1StartDate, requestHost1[0].StartDate)
		assert.Equal(t, expectedLine1EndDate, requestHost1[0].EndDate)
		assert.Equal(t, expectedLine2StartDate, requestHost1[1].StartDate)
		assert.Equal(t, expectedLine2EndDate, requestHost1[1].EndDate)

		requestHost2 := requestsByHost["host_000002"]
		assert.Equal(t, 1, len(requestHost2))

		expectedLine3StartDate, _ := time.Parse(timeFormat, "2017-01-02 18:50:28")
		expectedLine3EndDate, _ := time.Parse(timeFormat, "2017-01-02 19:50:28")
		assert.Equal(t, expectedLine3StartDate, requestHost2[0].StartDate)
		assert.Equal(t, expectedLine3EndDate, requestHost2[0].EndDate)
	})
	t.Run("file not found", func(t *testing.T) {
		_, err := ParseFile(testFilesPath + "/file_xxx.csv")
		exception := excepts.FromError(err)

		assert.NotNil(t, err)
		assert.Equal(t, excepts.FailedOpeningCSVFile, exception.Code)
	})
	t.Run("invalid file", func(t *testing.T) {
		_, err := ParseFile(testFilesPath + "/invalid_file.csv")
		exception := excepts.FromError(err)

		assert.NotNil(t, err)
		assert.Equal(t, excepts.InvalidNumberOfColumns, exception.Code)
	})
}

func TestParseLines(t *testing.T) {
	t.Run("valid file", func(t *testing.T) {
		file, _ := os.Open(testFilesPath + "/valid_file.csv")

		requestsByHost, err := parseLines(file)

		assert.Nil(t, err)
		assert.Equal(t, 2, len(requestsByHost))

		requestHost1 := requestsByHost["host_000001"]
		assert.Equal(t, 2, len(requestHost1))

		expectedLine1StartDate, _ := time.Parse(timeFormat, "2017-01-01 08:59:22")
		expectedLine1EndDate, _ := time.Parse(timeFormat, "2017-01-01 09:59:22")
		expectedLine2StartDate, _ := time.Parse(timeFormat, "2017-01-02 13:02:02")
		expectedLine2EndDate, _ := time.Parse(timeFormat, "2017-01-02 14:02:02")
		assert.Equal(t, expectedLine1StartDate, requestHost1[0].StartDate)
		assert.Equal(t, expectedLine1EndDate, requestHost1[0].EndDate)
		assert.Equal(t, expectedLine2StartDate, requestHost1[1].StartDate)
		assert.Equal(t, expectedLine2EndDate, requestHost1[1].EndDate)

		requestHost2 := requestsByHost["host_000002"]
		assert.Equal(t, 1, len(requestHost2))

		expectedLine3StartDate, _ := time.Parse(timeFormat, "2017-01-02 18:50:28")
		expectedLine3EndDate, _ := time.Parse(timeFormat, "2017-01-02 19:50:28")
		assert.Equal(t, expectedLine3StartDate, requestHost2[0].StartDate)
		assert.Equal(t, expectedLine3EndDate, requestHost2[0].EndDate)
	})
	t.Run("invalid file", func(t *testing.T) {
		file, _ := os.Open(testFilesPath + "/invalid_file.csv")

		_, err := parseLines(file)
		exception := excepts.FromError(err)

		assert.NotNil(t, err)
		assert.Equal(t, excepts.InvalidNumberOfColumns, exception.Code)
	})
}

func TestValidateAndConvertColumns(t *testing.T) {

	t.Run("sucessfull", func(t *testing.T) {
		scanner := openScanner(testFilesPath + "/valid_line.csv")

		hostName, startDate, endDate, err := validateAndConvertColumns(scanner)

		expectedStartDate, _ := time.Parse(timeFormat, "2017-01-02 13:02:02")
		expectedEndDate, _ := time.Parse(timeFormat, "2017-01-02 14:02:02")

		assert.Nil(t, err)
		assert.Equal(t, "host_000001", *hostName)
		assert.Equal(t, expectedStartDate, *startDate)
		assert.Equal(t, expectedEndDate, *endDate)
	})

	type test struct {
		name       string
		file       string
		excepected string
	}
	tests := []test{
		{name: "invalid number of columns", file: "/invalid_number_of_columns.csv", excepected: excepts.InvalidNumberOfColumns},
		{name: "missing hostname", file: "/missing_hostname.csv", excepected: excepts.MissingHostnameColumn},
		{name: "missing start date", file: "/missing_start_date.csv", excepected: excepts.MissingStartDateColumn},
		{name: "missing end date", file: "/missing_end_date.csv", excepected: excepts.MissingEndDateColumn},
		{name: "invalid start date", file: "/invalid_start_date.csv", excepected: excepts.InvalidStartDateColumn},
		{name: "invalid end date", file: "/invalid_end_date.csv", excepected: excepts.InvalidEndDateColumn},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			scanner := openScanner(testFilesPath + testCase.file)

			_, _, _, err := validateAndConvertColumns(scanner)
			exception := excepts.FromError(err)

			assert.NotNil(t, err)
			assert.Equal(t, exception.Code, testCase.excepected)
		})
	}
}

func openScanner(path string) *bufio.Scanner {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	return scanner
}

func TestGroupByHost(t *testing.T) {

	requestsByHost := make(model.QueryRequestsByHost)

	host1Request1 := model.QueryRequest{
		StartDate: time.Now(),
		EndDate:   time.Now().Add(time.Duration(100)),
	}
	groupByHost(requestsByHost, "host1", host1Request1)

	host1Request2 := model.QueryRequest{
		StartDate: time.Now(),
		EndDate:   time.Now().Add(time.Duration(200)),
	}
	groupByHost(requestsByHost, "host1", host1Request2)

	host2Request1 := model.QueryRequest{
		StartDate: time.Now(),
		EndDate:   time.Now().Add(time.Duration(300)),
	}
	groupByHost(requestsByHost, "host2", host2Request1)

	requestsHost1 := requestsByHost["host1"]
	assert.Equal(t, 2, len(requestsHost1))
	assert.Equal(t, host1Request1.StartDate, requestsHost1[0].StartDate)
	assert.Equal(t, host1Request1.EndDate, requestsHost1[0].EndDate)
	assert.Equal(t, host1Request2.StartDate, requestsHost1[1].StartDate)
	assert.Equal(t, host1Request2.EndDate, requestsHost1[1].EndDate)

	requestsHost2 := requestsByHost["host2"]
	assert.Equal(t, 1, len(requestsHost2))
	assert.Equal(t, host2Request1.StartDate, requestsHost2[0].StartDate)
	assert.Equal(t, host2Request1.EndDate, requestsHost2[0].EndDate)
}

func BenchmarkParseLines(b *testing.B) {

	file, err := os.Open("../../test/data/query_params_bench.csv")
	if err != nil {
		logger.FatalError(err)
	}

	for n := 0; n < b.N; n++ {
		requestsByHost, err := parseLines(file)

		assert.Nil(b, err)
		assert.NotNil(b, requestsByHost)
	}
}
