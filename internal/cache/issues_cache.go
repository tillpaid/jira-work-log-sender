package cache

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	issuesCacheFile = "issues.json"
	issuesLimit     = 100
)

type IssuesCacheItem struct {
	IssueID  string
	IssueKey string
}

type IssuesCacheData map[string]IssuesCacheItem

type IssuesCache struct {
	cacheDir string
	items    IssuesCacheData
}

func NewIssuesCache(cacheDir string) (*IssuesCache, error) {
	items, err := loadIssuesCacheData(cacheDir)
	if err != nil {
		return nil, err
	}

	return &IssuesCache{cacheDir: cacheDir, items: items}, nil
}

func (c *IssuesCache) IsIssueExists(issueKey string) bool {
	_, exists := c.items[issueKey]
	return exists
}

func (c *IssuesCache) GetIssue(issueKey string) (IssuesCacheItem, bool) {
	item, exists := c.items[issueKey]
	return item, exists
}

func (c *IssuesCache) SaveIssue(issueKey, issueID string) error {
	c.items[issueKey] = IssuesCacheItem{IssueID: issueID, IssueKey: issueKey}

	if len(c.items) > issuesLimit {
		c.removeHalf()
	}

	return writeIssuesCacheData(c.cacheDir, c.items)
}

func (c *IssuesCache) removeHalf() {
	var count int

	for key := range c.items {
		if count%2 == 0 {
			delete(c.items, key)
		}
		count++
	}
}

func loadIssuesCacheData(cacheDir string) (IssuesCacheData, error) {
	filePath := filepath.Join(cacheDir, issuesCacheFile)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		initialData := make(IssuesCacheData)

		if err = writeIssuesCacheData(cacheDir, initialData); err != nil {
			return nil, err
		}
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var data IssuesCacheData
	if err = json.Unmarshal(content, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return data, nil
}

func writeIssuesCacheData(cacheDir string, data IssuesCacheData) error {
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
