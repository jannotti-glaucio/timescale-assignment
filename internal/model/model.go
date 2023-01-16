package model

type QueryRequest struct {
	StartDate string
	EndData   string
}

type QueryRequests []QueryRequest

type QueryRequestsByHost map[string]QueryRequests

type QueryResult struct {
	MinUsage int
	MAxUsage int
}

type QueryResults []QueryResult

type QueryResultsByHost map[string]QueryResults
