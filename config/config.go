package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	var err error
	if env := os.Getenv("ENV"); env != "production" {
		err = godotenv.Load(".env")
	}

	if err != nil {
		panic(err)
	}
}

func LoadConfig() {
	LoadEnv()
}
