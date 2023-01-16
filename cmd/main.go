package main

import (
	"github.com/jannotti-glaucio/timescale-assignment/internal/logger"
	"github.com/jannotti-glaucio/timescale-assignment/internal/parsers"
	"github.com/jannotti-glaucio/timescale-assignment/internal/workers"

	"context"

	_ "go.uber.org/automaxprocs"
)

func main() {

	// Init Logger
	logger.Init()
	defer logger.Clean()

	// Init Context
	context.Background()

	requests := parsers.ParseFile("./query_params.csv")

	workers.RunWorkers(requests)
}
