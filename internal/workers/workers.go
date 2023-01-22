package workers

import (
	"context"
	"time"

	"github.com/jannotti-glaucio/timescale-assignment/internal/logger"
	"github.com/jannotti-glaucio/timescale-assignment/internal/parsers"
	"github.com/jannotti-glaucio/timescale-assignment/internal/repository"

	"sync"
)

type ResultChannel struct {
	HostName string
	Results  []parsers.QueryResult
}

func RunWorkers(requests parsers.QueryRequestsByHost) parsers.QueryResults {

	numJobs := len(requests)
	var waitToStart sync.WaitGroup
	var waitToFinish sync.WaitGroup
	resulChan := make(chan ResultChannel, numJobs)

	waitToStart.Add(1)
	for key, value := range requests {
		waitToFinish.Add(1)

		logger.Info("Starting worker for hostname [%v]", key)
		go worker(key, value, &waitToStart, &waitToFinish, resulChan)
	}
	waitToStart.Done()

	waitToFinish.Wait()

	defer close(resulChan)
	var allHostsResults parsers.QueryResults
	for a := 1; a <= numJobs; a++ {
		result := <-resulChan
		results := result.Results

		allHostsResults = append(allHostsResults, results...)
	}

	return allHostsResults
}

func worker(hostname string, requests parsers.QueryRequests, waitToStart *sync.WaitGroup, waitToFinish *sync.WaitGroup, resulChan chan<- ResultChannel) {
	defer waitToFinish.Done()

	ctx := context.Background()

	logger.Info("Starting executing %d queries for hostname [%s]", len(requests), hostname)
	var results parsers.QueryResults

	conn, err := repository.OpenConnection(ctx)
	if err != nil {
		logger.Fatal("Error opening database connection [%v]", err)
	}

	defer repository.CloseConnection(ctx, conn)

	waitToStart.Wait()

	for _, request := range requests {

		start := time.Now()
		maxUsage, minUsage, err := repository.RunQuery(ctx, conn, hostname, request.StartDate, request.EndDate)
		duration := time.Since(start)
		if err != nil {
			logger.Fatal("Error executing query: %v", err)
		}

		result := parsers.QueryResult{
			Duration: duration,
			MinUsage: maxUsage,
			MaxUsage: minUsage,
		}
		results = append(results, result)
	}

	logger.Info("Finished executing queries for hostname [%s]", hostname)
	resulChan <- ResultChannel{
		HostName: hostname,
		Results:  results,
	}
}
