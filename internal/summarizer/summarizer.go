package summarizer

import (
	"sort"
	"time"

	"github.com/jannotti-glaucio/timescale-assignment/internal/model"
)

type SummarizeResult struct {
	NumberOfQueries  int
	MinimumQueryTime time.Duration
	MedianQueryTime  time.Duration
	AverageQueryTime time.Duration
	MaximumQueryTime time.Duration
}

func SummarizeResults(results model.QueryResults) SummarizeResult {

	sort.Sort(model.QueryResultsByDuration(results))

	numberOfQueries := len(results)
	medianQueryTime := getMedian(results)

	minimumQueryTime, maximumQueryTime, averageQueryTime := getMinAndMaxAndAverage(results)

	return SummarizeResult{
		NumberOfQueries:  numberOfQueries,
		MinimumQueryTime: minimumQueryTime,
		MedianQueryTime:  medianQueryTime,
		AverageQueryTime: averageQueryTime,
		MaximumQueryTime: maximumQueryTime,
	}
}

func getMinAndMaxAndAverage(results model.QueryResults) (time.Duration, time.Duration, time.Duration) {

	var durationsSum time.Duration
	var minimumQueryTime, maximumQueryTime time.Duration
	for i, result := range results {
		durationsSum += result.Duration

		if i == 0 {
			// Uses the first result as base to min and max
			minimumQueryTime = result.Duration
			maximumQueryTime = result.Duration

		} else {
			if result.Duration < minimumQueryTime {
				minimumQueryTime = result.Duration
			}

			if result.Duration > maximumQueryTime {
				maximumQueryTime = result.Duration
			}
		}
	}
	average := durationsSum.Nanoseconds() / int64(len(results))
	averageQueryTime := time.Duration(average)

	return minimumQueryTime, maximumQueryTime, averageQueryTime
}

func getMedian(results model.QueryResults) time.Duration {

	len := len(results)

	if len == 0 {
		return time.Duration(0)
	}
	if len == 1 {
		return results[0].Duration
	}

	if len%2 == 0 {
		// If even length
		middle := len / 2
		return (results[middle-1].Duration + results[middle].Duration) / 2

	} else {
		// If odd length
		middle := len / 2
		return results[middle].Duration
	}
}
