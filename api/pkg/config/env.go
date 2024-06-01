package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kr/pretty"
)

func init() {
	envFile := os.Getenv("ENV_FILE")

	if envFile != "" {
		err := godotenv.Load(envFile)
		if err != nil {
			pretty.Log("Error loading .env file", envFile, err)
		}
	} else {

		err := godotenv.Load()
		if err != nil {
			pretty.Log("Error loading default .env file", err)
		}

		return
	}
}
