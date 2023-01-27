//go:build !integration

package summarizer

import (
	"testing"
	"time"

	"github.com/jannotti-glaucio/timescale-assignment/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestSummarizeResults(t *testing.T) {

	result1 := model.QueryResult{
		Duration: time.Duration(250),
		MinUsage: 20,
		MaxUsage: 80,
	}
	result2 := model.QueryResult{
		Duration: time.Duration(350),
		MinUsage: 30,
		MaxUsage: 70,
	}
	result3 := model.QueryResult{
		Duration: time.Duration(150),
		MinUsage: 10,
		MaxUsage: 90,
	}
	results := model.QueryResults{result1, result2, result3}

	summarized := SummarizeResults(results)

	assert.Equal(t, 3, summarized.NumberOfQueries)
	assert.Equal(t, time.Duration(150), summarized.MinimumQueryTime)
	assert.Equal(t, time.Duration(250), summarized.MedianQueryTime)
	assert.Equal(t, time.Duration(250), summarized.AverageQueryTime)
	assert.Equal(t, time.Duration(350), summarized.MaximumQueryTime)
}

func TestGetMinAndMaxAndAverage(t *testing.T) {

	t.Run("ordered results", func(t *testing.T) {
		result1 := model.QueryResult{
			Duration: time.Duration(100),
			MinUsage: 10,
			MaxUsage: 90,
		}
		result2 := model.QueryResult{
			Duration: time.Duration(300),
			MinUsage: 30,
			MaxUsage: 70,
		}
		result3 := model.QueryResult{
			Duration: time.Duration(200),
			MinUsage: 20,
			MaxUsage: 80,
		}
		results := model.QueryResults{result1, result2, result3}

		minimumQueryTime, maximumQueryTime, averageQueryTime := getMinAndMaxAndAverage(results)

		assert.Equal(t, time.Duration(100), minimumQueryTime)
		assert.Equal(t, time.Duration(300), maximumQueryTime)
		assert.Equal(t, time.Duration(200), averageQueryTime)
	})
	t.Run("out of order results", func(t *testing.T) {
		result1 := model.QueryResult{
			Duration: time.Duration(300),
			MinUsage: 30,
			MaxUsage: 70,
		}
		result2 := model.QueryResult{
			Duration: time.Duration(100),
			MinUsage: 10,
			MaxUsage: 90,
		}
		result3 := model.QueryResult{
			Duration: time.Duration(200),
			MinUsage: 20,
			MaxUsage: 80,
		}
		results := model.QueryResults{result1, result2, result3}

		minimumQueryTime, maximumQueryTime, averageQueryTime := getMinAndMaxAndAverage(results)

		assert.Equal(t, time.Duration(100), minimumQueryTime)
		assert.Equal(t, time.Duration(300), maximumQueryTime)
		assert.Equal(t, time.Duration(200), averageQueryTime)
	})
}

func TestGetMedian(t *testing.T) {
	t.Run("one result", func(t *testing.T) {
		result := model.QueryResult{
			Duration: time.Duration(100),
			MinUsage: 1,
			MaxUsage: 99,
		}
		results := model.QueryResults{result}

		median := getMedian(results)

		assert.Equal(t, result.Duration, median)
	})
	t.Run("zero results", func(t *testing.T) {
		results := model.QueryResults{}

		median := getMedian(results)

		assert.Equal(t, time.Duration(0), median)
	})
	t.Run("multipe odd results", func(t *testing.T) {
		result1 := model.QueryResult{
			Duration: time.Duration(100),
			MinUsage: 10,
			MaxUsage: 90,
		}
		result2 := model.QueryResult{
			Duration: time.Duration(200),
			MinUsage: 20,
			MaxUsage: 80,
		}
		result3 := model.QueryResult{
			Duration: time.Duration(300),
			MinUsage: 30,
			MaxUsage: 70,
		}
		results := model.QueryResults{result1, result2, result3}

		median := getMedian(results)

		assert.Equal(t, time.Duration(200), median)
	})
	t.Run("multipe even results", func(t *testing.T) {
		result1 := model.QueryResult{
			Duration: time.Duration(100),
			MinUsage: 10,
			MaxUsage: 90,
		}
		result2 := model.QueryResult{
			Duration: time.Duration(200),
			MinUsage: 20,
			MaxUsage: 80,
		}
		result3 := model.QueryResult{
			Duration: time.Duration(300),
			MinUsage: 30,
			MaxUsage: 70,
		}
		result4 := model.QueryResult{
			Duration: time.Duration(400),
			MinUsage: 40,
			MaxUsage: 60,
		}
		results := model.QueryResults{result1, result2, result3, result4}

		median := getMedian(results)

		assert.Equal(t, time.Duration(250), median)
	})
}
