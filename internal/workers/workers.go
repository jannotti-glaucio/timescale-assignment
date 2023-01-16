package workers

import (
	"github.com/jannotti-glaucio/timescale-assignment/internal/database"
	"github.com/jannotti-glaucio/timescale-assignment/internal/logger"
	"github.com/jannotti-glaucio/timescale-assignment/internal/model"

	"sync"
)

type ResultChannel struct {
	HostName string
	Results  []model.QueryResult
}

func RunWorkers(requests model.QueryRequestsByHost) model.QueryResultsByHost {

	numJobs := len(requests)
	var waitGroup sync.WaitGroup
	channel := make(chan ResultChannel, numJobs)

	for key, value := range requests {
		waitGroup.Add(1)

		logger.Info("Starting worker for hostname %s", key)
		go worker(key, value, &waitGroup, channel)
	}

	waitGroup.Wait()

	resultsByHost := make(model.QueryResultsByHost)
	for a := 1; a <= numJobs; a++ {
		result := <-channel
		results := result.Results

		resultsByHost[result.HostName] = results
	}
	close(channel)

	return resultsByHost
}

func worker(hostname string, requests model.QueryRequests, waitGroup *sync.WaitGroup, channel chan<- ResultChannel) {
	defer waitGroup.Done()

	logger.Info("Starting executing %d queries for hostname %s", len(requests), hostname)
	var results model.QueryResults

	conn := database.OpenConnection("postgres://homework:abc123@localhost:5432/homework")
	defer database.CloseConnection(conn)

	for _, request := range requests {
		logger.Debug("Executing query for hostname %s, startDate: %s, endData: %s", hostname, request.StartDate, request.EndData)

		row := database.QueryRow(conn, "select max(usage) as maxUsage, min(usage) as minUsage from cpu_usage cu where host = '$1' and ts between '$2' and '$3'", hostname, request.StartDate, request.EndData)
		var maxUsage int
		var minUsage int
		row.Scan(&maxUsage, &minUsage)

		result := model.QueryResult{
			MinUsage: maxUsage,
			MAxUsage: minUsage,
		}
		results = append(results, result)
	}

	logger.Info("Finished executing queries for hostname %s", hostname)
	channel <- ResultChannel{
		HostName: hostname,
		Results:  results,
	}
}
