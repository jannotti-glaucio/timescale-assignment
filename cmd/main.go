package main

import (
	"os"
	"time"

	"github.com/jannotti-glaucio/timescale-assignment/internal/env"
	"github.com/jannotti-glaucio/timescale-assignment/internal/logger"
	"github.com/jannotti-glaucio/timescale-assignment/internal/parsers"
	"github.com/jannotti-glaucio/timescale-assignment/internal/summarizer"
	"github.com/jannotti-glaucio/timescale-assignment/internal/workers"
)

func main() {

	// Logger
	logger.Init()
	defer logger.Clean()

	// Environment variables
	env.LoadFromFile()
	err := env.CheckVars()
	if err != nil {
		logger.Fatal("Error loading envrionment variables: %v", err)
	}

	filePath := os.Getenv(env.FilePath)
	requests, err := parsers.ParseFile(filePath)
	if err != nil {
		logger.Fatal("Error parsing file: %v", err)
	}

	processingStart := time.Now()
	resultsByHost := workers.RunWorkers(requests)
	totalProcessingTime := time.Since(processingStart)

	summarizeResult := summarizer.SummarizeResults(resultsByHost)

	logger.Info("##### Processing Results #####")
	logger.Info("Number of Queries:     [%d]", summarizeResult.NumberOfQueries)
	logger.Info("Total Processing Time: [%v] milliseconds", totalProcessingTime.Milliseconds())
	logger.Info("Minimum Query Time:    [%v] nanoseconds", summarizeResult.MinimumQueryTime.Nanoseconds())
	logger.Info("Median Query Time:     [%v] nanoseconds", summarizeResult.MedianQueryTime.Nanoseconds())
	logger.Info("Average Query Time:    [%v] nanoseconds", summarizeResult.AverageQueryTime.Nanoseconds())
	logger.Info("Maximum Query Time:    [%v] nanoseconds", summarizeResult.MaximumQueryTime.Nanoseconds())
}
