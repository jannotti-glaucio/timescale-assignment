package parsers

import (
	"github.com/jannotti-glaucio/timescale-assignment/internal/logger"
	"github.com/jannotti-glaucio/timescale-assignment/internal/model"

	"bufio"
	"os"
	"strings"
)

func ParseFile(path string) model.QueryRequestsByHost {

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

func parseLines(file *os.File) model.QueryRequestsByHost {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	requests := make(model.QueryRequestsByHost)
	line := 0
	for scanner.Scan() {
		line++

		// Ignore head line
		if line <= 1 {
			continue
		}

		line := strings.Split(scanner.Text(), ",")

		hostName := line[0]

		request := model.QueryRequest{
			StartDate: line[1],
			EndData:   line[2],
		}

		groupByHost(requests, hostName, request)
	}
	return requests
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
