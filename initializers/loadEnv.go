package initializers

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv()  {
	err := godotenv.Load()
	if err != nil {
	  fmt.Print("Error loading .env file")
	}
}