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
		Url   string `yaml:"url" validate:"required"`
		User  string `yaml:"user" validate:"required"`
		Token string `yaml:"token" validate:"required"`
	} `yaml:"jira"`

	Highlighting struct {
		DefaultThresholdHours int            `yaml:"defaultThresholdHours" validate:"required"`
		TagSpecificThresholds map[string]int `yaml:"tagSpecificThresholds"`
		ExcludedIssues        []string       `yaml:"excludedIssues" validate:"required"`
	} `yaml:"highlighting"`

	TimeAdjustment struct {
		Enabled                bool     `yaml:"enabled"`
		ExcludedIssues         []string `yaml:"excludedIssues"`
		TargetDailyMinutes     int      `yaml:"targetDailyMinutes" validate:"required,min=0"`
		RemainingTimeThreshold int      `yaml:"remainingTimeThreshold" validate:"required,min=0"`
	} `yaml:"timeAdjustment"`

	Input struct {
		WorkLogFile string `yaml:"workLogFile" validate:"required"`
	} `yaml:"input"`

	Cache struct {
		Directory string `yaml:"directory" validate:"required"`
	} `yaml:"cache"`

	Tags struct {
		Allowed []string `yaml:"allowed"`
	} `yaml:"tags"`

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

	config.Input.WorkLogFile = filepath.Join(homeDir, config.Input.WorkLogFile)
	config.Cache.Directory = filepath.Join(homeDir, config.Cache.Directory)
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
