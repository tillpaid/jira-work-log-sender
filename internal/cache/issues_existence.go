package cache

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	issuesCacheFile = "issues_existence.json"
	issuesLimit     = 100
)

type IssuesExistenceCache struct {
	cacheDir         string
	checkedExistence map[string]bool
}

func NewIssuesExistenceCache(cacheDir string) (*IssuesExistenceCache, error) {
	checkedExistence, err := loadCheckedExistenceData(cacheDir)
	if err != nil {
		return nil, err
	}

	return &IssuesExistenceCache{cacheDir: cacheDir, checkedExistence: checkedExistence}, nil
}

func (c *IssuesExistenceCache) IsExists(issueID string) bool {
	return c.checkedExistence[issueID]
}

func (c *IssuesExistenceCache) SaveExists(issueID string) error {
	c.checkedExistence[issueID] = true

	if len(c.checkedExistence) > issuesLimit {
		c.removeHalf()
	}

	return writeToCacheFile(c.cacheDir, c.checkedExistence)
}

func (c *IssuesExistenceCache) removeHalf() {
	count := len(c.checkedExistence) / 2
	keysToRemove := make([]string, 0, count)

	for k := range c.checkedExistence {
		keysToRemove = append(keysToRemove, k)
		if len(keysToRemove) >= count {
			break
		}
	}

	for _, k := range keysToRemove {
		delete(c.checkedExistence, k)
	}
}

func loadCheckedExistenceData(cacheDir string) (map[string]bool, error) {
	filePath := filepath.Join(cacheDir, issuesCacheFile)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		initialData := make(map[string]bool)

		if err = writeToCacheFile(cacheDir, initialData); err != nil {
			return nil, err
		}
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var data map[string]bool
	if err = json.Unmarshal(content, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return data, nil
}

func writeToCacheFile(cacheDir string, data map[string]bool) error {
	filePath := filepath.Join(cacheDir, issuesCacheFile)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	if _, err = file.Write(dataBytes); err != nil {
		return fmt.Errorf("failed to write data: %w", err)
	}

	return nil
}
