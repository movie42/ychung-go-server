package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
		panic(err)
	}
	// Return the value of the variable
	return os.Getenv(key)
}
