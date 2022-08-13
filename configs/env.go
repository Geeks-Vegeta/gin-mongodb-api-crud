package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetNameFromEnv() string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	return os.Getenv("NAME")
}

func MongoDBENV() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file")
	}
	return os.Getenv("MONGODB_URI")
}
