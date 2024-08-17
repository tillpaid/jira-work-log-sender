package resource

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

const (
	envFileName           = ".config/paysera-log-time/env"
	envKeyPathToInputFile = "PATH_TO_INPUT_FILE"
	envKeyJiraBaseUrl     = "JIRA_BASE_URL"
	envKeyJiraUsername    = "JIRA_USERNAME"
	envKeyJiraApiToken    = "JIRA_API_TOKEN"
)

var allowedTags = []string{
	"engineering activities", "documentation", "deployment&monitoring", "research&investigation",
	"code review", "communication", "environment issue", "operational work", "other",
}

type jiraConfig struct {
	BaseUrl  string
	Username string
	ApiToken string
}

type Config struct {
	PathToInputFile string
	Jira            jiraConfig
	AllowedTags     []string
}

func InitConfig() (*Config, error) {
	homeDir := os.Getenv("HOME")

	if err := godotenv.Load(filepath.Join(homeDir, envFileName)); err != nil {
		return nil, errors.New("error loading env file")
	}

	envValues, err := getEnvValues(
		envKeyPathToInputFile,
		envKeyJiraBaseUrl,
		envKeyJiraUsername,
		envKeyJiraApiToken,
	)
	if err != nil {
		return nil, err
	}

	return &Config{
		PathToInputFile: filepath.Join(homeDir, envValues[envKeyPathToInputFile]),
		Jira: jiraConfig{
			BaseUrl:  envValues[envKeyJiraBaseUrl],
			Username: envValues[envKeyJiraUsername],
			ApiToken: envValues[envKeyJiraApiToken],
		},
		AllowedTags: allowedTags,
	}, nil
}

func getEnvValues(keys ...string) (map[string]string, error) {
	envMap := make(map[string]string)

	for _, key := range keys {
		value := os.Getenv(key)
		if value == "" {
			return nil, fmt.Errorf("env variable %s is not set", key)
		}

		envMap[key] = value
	}

	return envMap, nil
}
