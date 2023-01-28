package main

import (
	"database/sql"
	"os"
	"time"

	"github.com/jannotti-glaucio/timescale-assignment/internal/database"
	"github.com/jannotti-glaucio/timescale-assignment/internal/env"
	"github.com/jannotti-glaucio/timescale-assignment/internal/logger"
	"github.com/jannotti-glaucio/timescale-assignment/internal/parsers"
	"github.com/jannotti-glaucio/timescale-assignment/internal/repository"
	"github.com/jannotti-glaucio/timescale-assignment/internal/summarizer"
	"github.com/jannotti-glaucio/timescale-assignment/internal/workers"
)

func main() {

	totalProcessingTime, summarizeResult, err := process()

	if err != nil {
		logger.Fatal("Error on processing: %v", err)
	}

	logger.Info("##### Processing Results #####")
	logger.Info("Number of Queries:     [%d]", summarizeResult.NumberOfQueries)
	logger.Info("Total Processing Time: [%v] milliseconds", totalProcessingTime.Milliseconds())
	logger.Info("Minimum Query Time:    [%v] nanoseconds", summarizeResult.MinimumQueryTime.Nanoseconds())
	logger.Info("Median Query Time:     [%v] nanoseconds", summarizeResult.MedianQueryTime.Nanoseconds())
	logger.Info("Average Query Time:    [%v] nanoseconds", summarizeResult.AverageQueryTime.Nanoseconds())
	logger.Info("Maximum Query Time:    [%v] nanoseconds", summarizeResult.MaximumQueryTime.Nanoseconds())
}

func process() (*time.Duration, *summarizer.SummarizeResult, error) {

	// Logger
	err := logger.Init()
	if err != nil {
		return nil, nil, err
	}

	// Environment variables
	err = env.LoadFromFile()
	if err != nil {
		return nil, nil, err
	}

	// Validate env variables
	err = env.CheckVars()
	if err != nil {
		return nil, nil, err
	}

	// Env files
	filePath := os.Getenv(env.FilePath)
	requests, err := parsers.ParseFile(filePath)
	if err != nil {
		return nil, nil, err
	}

	// Database
	db, err := database.OpenConnection()
	if err != nil {
		return nil, nil, err
	}

	defer cleanUp(db)

	repository := repository.NewRepository(db)
	workers := workers.NewWorkers(repository)

	processingStart := time.Now()
	resultsByHost := workers.RunWorkers(requests)
	totalProcessingTime := time.Since(processingStart)

	summarizeResult := summarizer.SummarizeResults(resultsByHost)

	return &totalProcessingTime, &summarizeResult, nil
}

func cleanUp(db *sql.DB) {

	dbErr := database.CloseConnection(db)
	if dbErr != nil {
		logger.FatalError(dbErr)
	}
}
