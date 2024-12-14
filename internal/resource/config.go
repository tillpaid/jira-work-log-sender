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
	envKeyCacheDir        = "CACHE_DIR"
)

var allowedTags = []string{
	"[Engineering activities]",
	"[Documentation]",
	"[Deployment&Monitoring]",
	"[Research&Investigation]",
	"[Code review]",
	"[Communication]",
	"[Environment Issue]",
	"[Operational work]",
	"[Other]",
}

var excludedFromSpentTimeHighlight = []string{
	"TIME",
}

type jiraConfig struct {
	BaseUrl  string
	Username string
	ApiToken string
}

type Config struct {
	PathToInputFile                string
	CacheDir                       string
	Jira                           jiraConfig
	AllowedTags                    []string
	ExcludedFromSpentTimeHighlight []string
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
		envKeyCacheDir,
	)
	if err != nil {
		return nil, err
	}

	return &Config{
		PathToInputFile: filepath.Join(homeDir, envValues[envKeyPathToInputFile]),
		CacheDir:        filepath.Join(homeDir, envValues[envKeyCacheDir]),
		Jira: jiraConfig{
			BaseUrl:  envValues[envKeyJiraBaseUrl],
			Username: envValues[envKeyJiraUsername],
			ApiToken: envValues[envKeyJiraApiToken],
		},
		AllowedTags:                    allowedTags,
		ExcludedFromSpentTimeHighlight: excludedFromSpentTimeHighlight,
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
