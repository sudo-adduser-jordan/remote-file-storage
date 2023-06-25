package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Env func to get env value
func Env(key string) string {

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return os.Getenv(key)
}
