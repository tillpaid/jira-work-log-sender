package resource

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

const (
	envFileName = ".config/paysera-log-time/env"
)

type Config struct {
	PathToInputFile string
}

func InitConfig() (*Config, error) {
	err := godotenv.Load(filepath.Join(os.Getenv("HOME"), envFileName))
	if err != nil {
		return nil, errors.New("error loading .env file")
	}

	return &Config{
		PathToInputFile: os.Getenv("PATH_TO_INPUT_FILE"),
	}, nil
}
