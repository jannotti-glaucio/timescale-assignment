package parsers

import (
	"time"

	"github.com/jannotti-glaucio/timescale-assignment/internal/logger"

	"bufio"
	"os"
	"strings"
)

const timeFormat = "2006-01-02 15:04:05"

func ParseFile(path string) QueryRequestsByHost {

	file, err := os.Open(path)
	if err != nil {
		logger.Fatal("Failed opening file: %s", err)
	}
	defer file.Close()

	logger.Debug("Reading csv file contents")
	requests := parseLines(file)

	logger.Debug("Csv file contents read and grouped by hostname")
	return requests
}

func parseLines(file *os.File) QueryRequestsByHost {
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

		hostName, startDate, endDate := validateAndConvertColumns(scanner, count)

		request := QueryRequest{
			StartDate: startDate,
			EndDate:   endDate,
		}

		groupByHost(requests, hostName, request)
	}
	return requests
}

func validateAndConvertColumns(scanner *bufio.Scanner, count int) (string, time.Time, time.Time) {

	line := strings.Split(scanner.Text(), ",")
	if len(line) < 3 {
		logger.Fatal("Error reading line [%d], it contains only [%d] columns: %s", count, len(line), scanner.Text())
	}

	hostName := line[0]
	if len(hostName) == 0 {
		logger.Fatal("Error reading line [%d], invalid hostname column [%s]", count, hostName)
	}

	if len(line[1]) == 0 {
		logger.Fatal("Error reading line [%d], invalid startDate column [%s]", count, line[1])
	}
	startDate, err := time.Parse(timeFormat, line[1])
	if err != nil {
		logger.Fatal("Error reading line [%d], invalid startDate column [%s]", count, line[1])
	}

	if len(line[2]) == 0 {
		logger.Fatal("Error reading line [%d], invalid endDate column [%s]", count, line[2])
	}
	endDate, err := time.Parse(timeFormat, line[2])
	if err != nil {
		logger.Fatal("Error reading line [%d], invalid endDate column [%s]", count, line[2])
	}
	return hostName, startDate, endDate
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
