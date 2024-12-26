package jira

import (
	"fmt"

	"github.com/andygrunwald/go-jira"
	"github.com/tillpaid/jira-work-log-sender/internal/cache"
	"github.com/tillpaid/jira-work-log-sender/internal/model"
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
)

type IssueServiceInterface interface {
	IsIssueExists(issueID string) (bool, error)
	IsIssueExistsInCache(issueID string) bool
}

type WorkLogServiceInterface interface {
	GetSpentTime(issueID string) (*model.WorkLogTime, error)
	SendWorkLog(workLog model.WorkLog) error
}

type Client struct {
	jiraClient     *jira.Client
	IssueService   IssueServiceInterface
	WorkLogService WorkLogServiceInterface
}

func NewClient(config *resource.Config) (*Client, error) {
	tp := jira.BasicAuthTransport{
		Username: config.Jira.Username,
		Password: config.Jira.ApiToken,
	}

	issuesExistenceCache, err := cache.NewIssuesExistenceCache(config.CacheDir)
	if err != nil {
		return nil, fmt.Errorf("failed to create issues existence cache: %v", err)
	}

	jiraClient, err := jira.NewClient(tp.Client(), config.Jira.BaseUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create Jira client: %v", err)
	}

	return &Client{
		jiraClient:     jiraClient,
		IssueService:   newIssueService(jiraClient, issuesExistenceCache),
		WorkLogService: newWorkLogService(config.Jira.Username, jiraClient, config),
	}, nil
}

func (c *Client) Ping() error {
	_, response, err := c.jiraClient.User.GetSelf()
	if err != nil {
		newError := fmt.Errorf("failed to authenticate with Jira: %v", err)
		if response != nil {
			newError = fmt.Errorf("failed to authenticate with Jira. Response code is %d", response.StatusCode)
		}

		return newError
	}

	return nil
}
