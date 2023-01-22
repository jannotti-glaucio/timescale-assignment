package env

import (
	"os"

	"github.com/jannotti-glaucio/timescale-assignment/internal/exceptions"
	"github.com/jannotti-glaucio/timescale-assignment/internal/file"
	"github.com/joho/godotenv"
)

const FilePath = "FILE_PATH"
const DbUrl = "DB_URL"

const errorMissingVariable = "Missing environment variable %s"
const errorLoadingFile = "Error loading .env file: %v"

func CheckVars() error {
	if os.Getenv(FilePath) == "" {
		return exceptions.ThrowError(errorMissingVariable, FilePath)
	}

	if os.Getenv(DbUrl) == "" {
		return exceptions.ThrowError(errorMissingVariable, DbUrl)
	}

	return nil
}

func LoadFromFile() error {

	if !file.FileExists("./.env") {
		return nil
	}

	err := godotenv.Load()
	if err != nil {
		return exceptions.ThrowError(errorLoadingFile, err)
	}

	return nil
}
