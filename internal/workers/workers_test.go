//go:build !integration

package workers

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/jannotti-glaucio/timescale-assignment/internal/database"
	"github.com/jannotti-glaucio/timescale-assignment/internal/excepts"
	"github.com/jannotti-glaucio/timescale-assignment/internal/logger"
	"github.com/jannotti-glaucio/timescale-assignment/internal/model"
	"github.com/jannotti-glaucio/timescale-assignment/internal/repository"
	"github.com/jannotti-glaucio/timescale-assignment/internal/tests/mocks"

	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExecWorker(t *testing.T) {

	t.Run("sucessfull", func(t *testing.T) {
		startDate, _ := time.Parse(time.RFC3339, "2017-01-01 08:59:22")
		endDate, _ := time.Parse(time.RFC3339, "2017-01-01 09:59:22")
		requests := model.QueryRequests{
			{
				StartDate: startDate,
				EndDate:   endDate,
			},
		}
		maxUsage := float64(200)
		minUsage := float64(100)

		mockedRepository := mocks.NewMockedRepository()
		mockedRepository.
			On("RunQuery", "host-01", startDate, endDate).
			Return(maxUsage, minUsage, nil)

		var waitToStart sync.WaitGroup
		waitToStart.Add(1)
		var waitToFinish sync.WaitGroup
		waitToFinish.Add(1)
		resulChan := make(chan ResultChannel, 1)

		waitToStart.Done()
		err := execWorker(mockedRepository, "host-01", requests, &waitToStart, &waitToFinish, resulChan)

		waitToFinish.Wait()
		resultChannel := <-resulChan

		mockedRepository.AssertExpectations(t)

		assert.Nil(t, err)
		assert.NotNil(t, resultChannel)
		assert.Equal(t, "host-01", resultChannel.HostName)
		assert.Equal(t, 1, len(resultChannel.Results))

		result := resultChannel.Results[0]
		assert.Equal(t, float64(200), result.MaxUsage)
		assert.Equal(t, float64(100), result.MinUsage)
	})
	t.Run("error", func(t *testing.T) {
		startDate, _ := time.Parse(time.RFC3339, "2017-01-01 08:59:22")
		endDate, _ := time.Parse(time.RFC3339, "2017-01-01 09:59:22")
		requests := model.QueryRequests{
			{
				StartDate: startDate,
				EndDate:   endDate,
			},
		}
		var maxUsage, minUsage float64

		mockedRepository := mocks.NewMockedRepository()
		mockedRepository.
			On("RunQuery", "host-01", startDate, endDate).
			Return(maxUsage, minUsage, excepts.ThrowException("001", "Error"))

		var waitToStart sync.WaitGroup
		waitToStart.Add(1)
		var waitToFinish sync.WaitGroup
		waitToFinish.Add(1)
		resulChan := make(chan ResultChannel, 1)

		waitToStart.Done()

		err := execWorker(mockedRepository, "host-01", requests, &waitToStart, &waitToFinish, resulChan)
		exception := excepts.FromError(err)

		mockedRepository.AssertExpectations(t)

		assert.NotNil(t, err)
		assert.Equal(t, "001", exception.Code)
	})
}

func BenchmarkExecWorker(b *testing.B) {

	os.Setenv("DB_URL", "postgres://homework:abc123@localhost:5434/homework")
	db, err := database.OpenConnection()
	if err != nil {
		logger.FatalError(err)
	}

	repository := repository.NewRepository(db)

	firstDate, _ := time.Parse(time.RFC3339, "2016-12-31T22:00:00Z")

	// Prepare requests
	var requests []model.QueryRequest
	for i := 1; i <= 100000; i++ { // 100k queries by host
		startDate := firstDate.AddDate(0, 0, rand.Intn(30)) // randon days between 1 month
		endDate := startDate.Add(time.Duration(time.Hour))  // 1 hour of duration

		requests = append(requests, model.QueryRequest{
			StartDate: startDate,
			EndDate:   endDate,
		})
	}

	for n := 0; n < b.N; n++ {
		hostname := fmt.Sprintf("host_%06d", rand.Intn(19))

		var waitToStart sync.WaitGroup
		waitToStart.Add(1)
		var waitToFinish sync.WaitGroup
		waitToFinish.Add(1)
		resulChan := make(chan ResultChannel, 1)

		waitToStart.Done()
		err := execWorker(repository, hostname, requests, &waitToStart, &waitToFinish, resulChan)

		if !assert.Nil(b, err) {
			logger.Fatal("Error executing worker: %v", err)
		}

		resultChannel := <-resulChan
		assert.NotNil(b, resultChannel)
	}
}
