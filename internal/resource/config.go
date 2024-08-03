package resource

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

const (
	envFileName           = ".config/paysera-log-time/env"
	envKeyPathToInputFile = "PATH_TO_INPUT_FILE"
	envKeyOutputShellFile = "OUTPUT_SHELL_FILE"
)

type Config struct {
	PathToInputFile string
	OutputShellFile string
	AllowedTags     []string
}

func InitConfig() (*Config, error) {
	allowedTags := []string{"CODING", "INVESTIGATION", "REVIEW", "DEPLOYMENT", "DOC", "RESEARCH", "MEETING", "OTHER"}
	homeDir := os.Getenv("HOME")

	err := godotenv.Load(filepath.Join(homeDir, envFileName))
	if err != nil {
		return nil, errors.New("error loading env file")
	}

	pathToInputFile, err := getEnvValue(envKeyPathToInputFile)
	if err != nil {
		return nil, err
	}

	outputShellFile, err := getEnvValue(envKeyOutputShellFile)
	if err != nil {
		return nil, err
	}

	return &Config{
		PathToInputFile: filepath.Join(homeDir, pathToInputFile),
		OutputShellFile: filepath.Join(homeDir, outputShellFile),
		AllowedTags:     allowedTags,
	}, nil
}

func getEnvValue(key string) (string, error) {
	value := os.Getenv(key)

	if value == "" {
		return "", fmt.Errorf("env variable %s is not set", key)
	}

	return value, nil
}
