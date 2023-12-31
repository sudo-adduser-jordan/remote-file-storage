package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Env(key string) string {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return os.Getenv(key)
}
