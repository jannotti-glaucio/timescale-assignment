package parsers

import (
	"time"

	"github.com/jannotti-glaucio/timescale-assignment/internal/exceptions"
	"github.com/jannotti-glaucio/timescale-assignment/internal/logger"

	"bufio"
	"os"
	"strings"
)

const timeFormat = "2006-01-02 15:04:05"

func ParseFile(path string) (QueryRequestsByHost, error) {

	file, err := os.Open(path)
	if err != nil {
		exceptions.ThrowError("Failed opening file: %s", err)
	}
	defer file.Close()

	logger.Debug("Reading csv file contents")
	requests, err := parseLines(file)
	if err != nil {
		exceptions.ThrowError("Error parsing file: %s", err)
	}

	logger.Debug("Csv file contents read and grouped by hostname")
	return requests, nil
}

func parseLines(file *os.File) (QueryRequestsByHost, error) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	requests := make(QueryRequestsByHost)
	count := 0
	for scanner.Scan() {
		count++

		// Ignore head line
		if count <= 1 {
			continue
		}

		hostName, startDate, endDate, err := validateAndConvertColumns(scanner)
		if err != nil {
			exceptions.ThrowError("Error reading line [%d]: [%v]", count, err)
		}

		request := QueryRequest{
			StartDate: startDate,
			EndDate:   endDate,
		}

		groupByHost(requests, hostName, request)
	}
	return requests, nil
}

func validateAndConvertColumns(scanner *bufio.Scanner) (string, time.Time, time.Time, error) {

	line := strings.Split(scanner.Text(), ",")
	if len(line) < 3 {
		exceptions.ThrowError("Invalid number of columns [%d] columns: %s", len(line), scanner.Text())
	}

	hostName := line[0]
	if len(hostName) == 0 {
		exceptions.ThrowError("Invalid hostname column [%s]", hostName)
	}

	if len(line[1]) == 0 {
		exceptions.ThrowError("Invalid startDate column [%s]", line[1])
	}
	startDate, err := time.Parse(timeFormat, line[1])
	if err != nil {
		exceptions.ThrowError("Invalid startDate column [%s]", line[1])
	}

	if len(line[2]) == 0 {
		exceptions.ThrowError("Invalid endDate column [%s]", line[2])
	}
	endDate, err := time.Parse(timeFormat, line[2])
	if err != nil {
		exceptions.ThrowError("Invalid endDate column [%s]", line[2])
	}
	return hostName, startDate, endDate, nil
}

func groupByHost(requestsByHost QueryRequestsByHost, hostName string, request QueryRequest) {

	requests, exists := requestsByHost[hostName]

	if exists {
		requests = append(requests, request)
		requestsByHost[hostName] = requests

	} else {
		var requests QueryRequests
		requests = append(requests, request)
		requestsByHost[hostName] = requests
	}
}
