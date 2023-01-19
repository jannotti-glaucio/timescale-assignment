package env

import (
	"os"

	"github.com/jannotti-glaucio/timescale-assignment/internal/file"
	"github.com/jannotti-glaucio/timescale-assignment/internal/logger"
	"github.com/joho/godotenv"
)

const FilePath string = "FILE_PATH"
const DbUrl string = "DB_URL"

func CheckVars() {
	if os.Getenv(FilePath) == "" {
		logger.Fatal("Missing environment variable %s", FilePath)
	}

	if os.Getenv(DbUrl) == "" {
		logger.Fatal("Missing environment variable %s", DbUrl)
	}
}

func LoadFromFile() {

	if !file.FileExists("./.env") {
		return
	}

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file: %v", err)
	}
}
