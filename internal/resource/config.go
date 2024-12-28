package resource

import (
	"os"
	"path/filepath"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"
)

const configFilePath = ".config/jira-work-log-sender/config.yml"

type Config struct {
	Jira struct {
		BaseUrl  string `yaml:"baseUrl" validate:"required"`
		Username string `yaml:"username" validate:"required"`
		ApiToken string `yaml:"apiToken" validate:"required"`
	} `yaml:"jira"`

	IssueHighlight struct {
		HighlightAfterHours int      `yaml:"highlightAfterHours" validate:"required"`
		ExcludedNumbers     []string `yaml:"excludedNumbers" validate:"required"`
	} `yaml:"issueHighlight"`

	TimeModification struct {
		Enabled         bool     `yaml:"enabled"`
		ExcludedNumbers []string `yaml:"excludedNumbers" validate:"required"`
	} `yaml:"timeModification"`

	PathToInputFile string   `yaml:"pathToInputFile" validate:"required"`
	CacheDir        string   `yaml:"cacheDir" validate:"required"`
	AllowedTags     []string `yaml:"allowedTags"`

	IsDevRun bool
}

func InitConfig() (*Config, error) {
	config, err := loadConfig()
	if err != nil {
		return nil, err
	}

	if err := validateConfig(config); err != nil {
		return nil, err
	}

	updateConfigValues(config)

	return config, nil
}

func updateConfigValues(config *Config) {
	homeDir := os.Getenv("HOME")

	config.PathToInputFile = filepath.Join(homeDir, config.PathToInputFile)
	config.CacheDir = filepath.Join(homeDir, config.CacheDir)
	config.IsDevRun = isDevRun()
}

func loadConfig() (*Config, error) {
	homeDir := os.Getenv("HOME")

	file, err := os.Open(filepath.Join(homeDir, configFilePath))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func validateConfig(config *Config) error {
	validate := validator.New()
	return validate.Struct(config)
}

func isDevRun() bool {
	for _, arg := range os.Args {
		if arg == "--dev" {
			return true
		}
	}

	return false
}
