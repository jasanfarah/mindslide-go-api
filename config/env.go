package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Openai_api_key string

func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

  Openai_api_key = os.Getenv("OPENAI_API_KEY")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
