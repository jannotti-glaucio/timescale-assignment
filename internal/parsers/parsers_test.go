package parsers

import (
	"bufio"
	"os"
	"testing"
	"time"

	"github.com/jannotti-glaucio/timescale-assignment/internal/excepts"
	"github.com/jannotti-glaucio/timescale-assignment/internal/tests"
	"github.com/stretchr/testify/assert"
)

const testFilesPath = tests.TestDataPath + "parsers/"

func TestParseFile(t *testing.T) {
	// TODO: Implement TestParseFile
}

func TestParseLines(t *testing.T) {
	// TODO: Implement TestParseLines
}

func TestValidateAndConvertColumns(t *testing.T) {

	t.Run("sucessfull", func(t *testing.T) {
		scanner := openScanner(testFilesPath + "/valid_line.csv")

		hostName, startDate, endDate, err := validateAndConvertColumns(scanner)

		expectedStartDate, _ := time.Parse(timeFormat, "2017-01-02 13:02:02")
		expectedEndData, _ := time.Parse(timeFormat, "2017-01-02 14:02:02")

		assert.Nil(t, err)
		assert.Equal(t, "host_000001", *hostName)
		assert.Equal(t, expectedStartDate, *startDate)
		assert.Equal(t, expectedEndData, *endDate)
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

			assert.NotNil(t, err)
			assert.Equal(t, err.Code, testCase.excepected)
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

	requestsByHost := make(QueryRequestsByHost)

	host1Request1 := QueryRequest{
		StartDate: time.Now(),
		EndDate:   time.Now().Add(time.Duration(100)),
	}
	groupByHost(requestsByHost, "host1", host1Request1)

	host1Request2 := QueryRequest{
		StartDate: time.Now(),
		EndDate:   time.Now().Add(time.Duration(200)),
	}
	groupByHost(requestsByHost, "host1", host1Request2)

	host2Request1 := QueryRequest{
		StartDate: time.Now(),
		EndDate:   time.Now().Add(time.Duration(300)),
	}
	groupByHost(requestsByHost, "host2", host2Request1)

	requestHost1 := requestsByHost["host1"]
	assert.Equal(t, 2, len(requestHost1))
	assert.Equal(t, host1Request1.StartDate, requestHost1[0].StartDate)
	assert.Equal(t, host1Request1.EndDate, requestHost1[0].EndDate)
	assert.Equal(t, host1Request2.StartDate, requestHost1[1].StartDate)
	assert.Equal(t, host1Request2.EndDate, requestHost1[1].EndDate)

	requestHost2 := requestsByHost["host2"]
	assert.Equal(t, 1, len(requestHost2))
	assert.Equal(t, host2Request1.StartDate, requestHost2[0].StartDate)
	assert.Equal(t, host2Request1.EndDate, requestHost2[0].EndDate)
}
