package workers

import (
	"github.com/jannotti-glaucio/timescale-assignment/internal/logger"
	"github.com/jannotti-glaucio/timescale-assignment/internal/model"
	"github.com/jannotti-glaucio/timescale-assignment/internal/repository"

	"sync"
	"time"
)

type (
	Workers struct {
		repo repository.Repository
	}
	ResultChannel struct {
		HostName string
		Results  []model.QueryResult
	}
)

func NewWorkers(repo repository.Repository) Workers {
	return Workers{repo: repo}
}

func (w Workers) RunWorkers(requests model.QueryRequestsByHost) model.QueryResults {

	numJobs := len(requests)
	var waitToStart sync.WaitGroup
	var waitToFinish sync.WaitGroup
	resulChan := make(chan ResultChannel, numJobs)

	waitToStart.Add(1)
	for key, value := range requests {
		waitToFinish.Add(1)

		logger.Info("Starting worker for hostname [%v]", key)
		go runWorker(w.repo, key, value, &waitToStart, &waitToFinish, resulChan)
	}
	waitToStart.Done()

	waitToFinish.Wait()

	defer close(resulChan)
	var allHostsResults model.QueryResults
	for a := 1; a <= numJobs; a++ {
		result := <-resulChan
		results := result.Results

		allHostsResults = append(allHostsResults, results...)
	}

	return allHostsResults
}

func runWorker(repository repository.Repository, hostname string, requests model.QueryRequests, waitToStart *sync.WaitGroup,
	waitToFinish *sync.WaitGroup, resulChan chan<- ResultChannel) {

	err := execWorker(repository, hostname, requests, waitToStart, waitToFinish, resulChan)
	if err != nil {
		logger.Fatal("Error executing worker: %v", err)
	}
}

func execWorker(repository repository.Repository, hostname string, requests model.QueryRequests, waitToStart *sync.WaitGroup,
	waitToFinish *sync.WaitGroup, resulChan chan<- ResultChannel) error {

	defer waitToFinish.Done()

	logger.Info("Starting executing %d queries for hostname [%s]", len(requests), hostname)
	var results model.QueryResults

	waitToStart.Wait()

	for _, request := range requests {

		start := time.Now()
		maxUsage, minUsage, err := repository.RunQuery(hostname, request.StartDate, request.EndDate)
		duration := time.Since(start)
		if err != nil {
			return err
		}

		result := model.QueryResult{
			Duration: duration,
			MinUsage: *minUsage,
			MaxUsage: *maxUsage,
		}
		results = append(results, result)
	}

	logger.Info("Finished executing queries for hostname [%s]", hostname)
	resulChan <- ResultChannel{
		HostName: hostname,
		Results:  results,
	}

	return nil
}
