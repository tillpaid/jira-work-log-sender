package resource

import (
	"os"
	"path/filepath"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"
)

const configFilePath = ".config/jira-work-log-sender/config.yml"

type JiraConfig struct {
	Url   string `yaml:"url" validate:"required"`
	User  string `yaml:"user" validate:"required"`
	Token string `yaml:"token" validate:"required"`
}

type TempoConfig struct {
	UseTempoApiToSendWorklogs          bool   `yaml:"useTempoApiToSendWorklogs"`
	WorkerID                           string `yaml:"workerID" validate:"required_if=UseTempoApiToSendWorklogs true"`
	EngineeringActivityName            string `yaml:"engineeringActivityName" validate:"required_if=UseTempoApiToSendWorklogs true"`
	EngineeringActivityWorkAttributeID int    `yaml:"engineeringActivityWorkAttributeID" validate:"required_if=UseTempoApiToSendWorklogs true"`
}

type HighlightingConfig struct {
	DefaultThresholdHours int            `yaml:"defaultThresholdHours" validate:"required"`
	TagSpecificThresholds map[string]int `yaml:"tagSpecificThresholds"`
	ExcludedIssues        []string       `yaml:"excludedIssues" validate:"required"`
}

type TimeAdjustmentConfig struct {
	Enabled                bool     `yaml:"enabled"`
	ExcludedIssues         []string `yaml:"excludedIssues"`
	TargetDailyMinutes     int      `yaml:"targetDailyMinutes" validate:"required,min=0"`
	RemainingTimeThreshold int      `yaml:"remainingTimeThreshold" validate:"required,min=0"`
}

type InputConfig struct {
	WorkLogFile string `yaml:"workLogFile" validate:"required"`
}

type CacheConfig struct {
	Directory string `yaml:"directory" validate:"required"`
}

type TagsConfig struct {
	Allowed []string `yaml:"allowed"`
}

type ForbiddenProjects []string

type Config struct {
	Jira              JiraConfig           `yaml:"jira"`
	Tempo             TempoConfig          `yaml:"tempo"`
	Highlighting      HighlightingConfig   `yaml:"highlighting"`
	TimeAdjustment    TimeAdjustmentConfig `yaml:"timeAdjustment"`
	ForbiddenProjects ForbiddenProjects    `yaml:"forbiddenProjects" validate:"required"`
	Input             InputConfig          `yaml:"input"`
	Cache             CacheConfig          `yaml:"cache"`
	Tags              TagsConfig           `yaml:"tags"`
	IsDevRun          bool
}

func InitConfig() (*Config, error) {
	cfg, err := loadConfig()
	if err != nil {
		return nil, err
	}

	if err := validateConfig(cfg); err != nil {
		return nil, err
	}

	updateConfigValues(cfg)

	return cfg, nil
}

func updateConfigValues(cfg *Config) {
	homeDir := os.Getenv("HOME")

	cfg.Input.WorkLogFile = filepath.Join(homeDir, cfg.Input.WorkLogFile)
	cfg.Cache.Directory = filepath.Join(homeDir, cfg.Cache.Directory)
	cfg.IsDevRun = isDevRun()
}

func loadConfig() (*Config, error) {
	homeDir := os.Getenv("HOME")

	file, err := os.Open(filepath.Join(homeDir, configFilePath))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cfg := &Config{}
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func validateConfig(cfg *Config) error {
	validate := validator.New()
	return validate.Struct(cfg)
}

func isDevRun() bool {
	for _, arg := range os.Args {
		if arg == "--dev" {
			return true
		}
	}

	return false
}
