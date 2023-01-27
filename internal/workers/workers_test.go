//go:build !integration

package workers

import (
	"github.com/jannotti-glaucio/timescale-assignment/internal/excepts"
	"github.com/jannotti-glaucio/timescale-assignment/internal/model"
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
		maxUsage := float32(200)
		minUsage := float32(100)

		mockedRepository := mocks.NewMockedRepository()
		mockedRepository.
			On("RunQuery", "host-01", startDate, endDate).
			Return(&maxUsage, &minUsage, nil)

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
		assert.Equal(t, float32(200), result.MaxUsage)
		assert.Equal(t, float32(100), result.MinUsage)
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
		var maxUsage, minUsage *float32

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
