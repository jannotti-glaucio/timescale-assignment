package model

import "time"

type QueryRequest struct {
	StartDate time.Time
	EndDate   time.Time
}

type QueryRequests []QueryRequest

type QueryRequestsByHost map[string]QueryRequests

type QueryResult struct {
	Duration time.Duration
	MinUsage float32
	MaxUsage float32
}

type QueryResults []QueryResult

type QueryResultsByDuration []QueryResult

func (queryResults QueryResultsByDuration) Len() int {
	return len(queryResults)
}

func (queryResults QueryResultsByDuration) Less(i, j int) bool {
	return queryResults[i].Duration < queryResults[j].Duration
}

func (queryResults QueryResultsByDuration) Swap(i, j int) {
	queryResults[i], queryResults[j] = queryResults[j], queryResults[i]
}
