package parsers

import (
	"time"

	"github.com/jannotti-glaucio/timescale-assignment/internal/excepts"
	"github.com/jannotti-glaucio/timescale-assignment/internal/logger"
	"github.com/jannotti-glaucio/timescale-assignment/internal/model"

	"bufio"
	"os"
	"strings"
)

const timeFormat = "2006-01-02 15:04:05"

func ParseFile(path string) (model.QueryRequestsByHost, error) {

	file, openErr := os.Open(path)
	if openErr != nil {
		return nil, excepts.ThrowException(excepts.FailedOpeningCSVFile, "Failed opening CSV file: %s", openErr)
	}
	defer file.Close()

	logger.Debug("Reading CSV file contents")
	requests, readErr := parseLines(file)
	if readErr != nil {
		return nil, readErr
	}

	logger.Debug("CSV file contents read and grouped by hostname")
	return requests, nil
}

func parseLines(file *os.File) (model.QueryRequestsByHost, error) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	requests := make(model.QueryRequestsByHost)
	count := 0
	for scanner.Scan() {
		count++

		// Ignore head line
		if count <= 1 {
			continue
		}

		hostName, startDate, endDate, err := validateAndConvertColumns(scanner)
		if err != nil {
			return nil, err
		}

		request := model.QueryRequest{
			StartDate: *startDate,
			EndDate:   *endDate,
		}

		groupByHost(requests, *hostName, request)
	}
	return requests, nil
}

func validateAndConvertColumns(scanner *bufio.Scanner) (*string, *time.Time, *time.Time, error) {

	line := strings.Split(scanner.Text(), ",")
	if len(line) < 3 {
		return nil, nil, nil, excepts.ThrowException(excepts.InvalidNumberOfColumns, "Invalid number of columns [%d]: %s", len(line), scanner.Text())
	}

	hostName := line[0]
	if len(hostName) == 0 {
		return nil, nil, nil, excepts.ThrowException(excepts.MissingHostnameColumn, "Missing hostname column [%s]", hostName)
	}

	if len(line[1]) == 0 {
		return nil, nil, nil, excepts.ThrowException(excepts.MissingStartDateColumn, "Missing startDate column [%s]", line[1])
	}
	startDate, err := time.Parse(timeFormat, line[1])
	if err != nil {
		return nil, nil, nil, excepts.ThrowException(excepts.InvalidStartDateColumn, "Invalid startDate column [%s]", line[1])
	}

	if len(line[2]) == 0 {
		return nil, nil, nil, excepts.ThrowException(excepts.MissingEndDateColumn, "Missing endDate column [%s]", line[2])
	}
	endDate, err := time.Parse(timeFormat, line[2])
	if err != nil {
		return nil, nil, nil, excepts.ThrowException(excepts.InvalidEndDateColumn, "Invalid endDate column [%s]", line[2])
	}
	return &hostName, &startDate, &endDate, nil
}

func groupByHost(requestsByHost model.QueryRequestsByHost, hostName string, request model.QueryRequest) {

	requests, exists := requestsByHost[hostName]

	if exists {
		requests = append(requests, request)
		requestsByHost[hostName] = requests

	} else {
		var requests model.QueryRequests
		requests = append(requests, request)
		requestsByHost[hostName] = requests
	}
}
