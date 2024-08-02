package resource

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	PathToInputFile string
}

func InitConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("error loading .env file")
	}

	return &Config{
		PathToInputFile: os.Getenv("PATH_TO_INPUT_FILE"),
	}, nil
}
