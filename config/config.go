package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}
}

func LoadConfig() {
	LoadEnv()
}
