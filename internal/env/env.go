package env

import (
	"os"

	"github.com/jannotti-glaucio/timescale-assignment/internal/excepts"
	"github.com/jannotti-glaucio/timescale-assignment/internal/file"
	"github.com/joho/godotenv"
)

const FilePath = "FILE_PATH"
const DbUrl = "DB_URL"

func CheckVars() *excepts.Exception {
	if os.Getenv(FilePath) == "" {
		return excepts.ThrowException(excepts.MissingEnvVariable, "Missing environment variable %s", FilePath)
	}

	if os.Getenv(DbUrl) == "" {
		return excepts.ThrowException(excepts.MissingEnvVariable, "Missing environment variable %s", DbUrl)
	}

	return nil
}

func LoadFromFile() *excepts.Exception {

	if !file.FileExists("./.env") {
		return nil
	}

	err := godotenv.Load()
	if err != nil {
		return excepts.ThrowException(excepts.ErrorLoadingEnvFile, "Error loading .env file: %v", err)
	}

	return nil
}
