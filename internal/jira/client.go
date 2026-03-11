package jira

import (
	"fmt"

	"github.com/andygrunwald/go-jira"
	"github.com/tillpaid/jira-work-log-sender/internal/cache"
	"github.com/tillpaid/jira-work-log-sender/internal/model"
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
)

type IssueServiceInterface interface {
	GetIssueID(issueKey string) (string, error)
	IsIssueExists(issueKey string) (bool, error)
	IsIssueExistsInCache(issueKey string) bool
}

type WorklogServiceInterface interface {
	GetSpentTime(issueKey string) (*model.WorklogTime, error)
	SendWorklog(worklog model.Worklog) error
}

type Client struct {
	jiraClient     *jira.Client
	IssueService   IssueServiceInterface
	WorklogService WorklogServiceInterface
}

func NewClient(cfg *resource.Config) (*Client, error) {
	tp := jira.BasicAuthTransport{
		Username: cfg.Jira.User,
		Password: cfg.Jira.Token,
	}

	issuesCache, err := cache.NewIssuesCache(cfg.Cache.Directory)
	if err != nil {
		return nil, fmt.Errorf("failed to create issues cache: %v", err)
	}

	jiraClient, err := jira.NewClient(tp.Client(), cfg.Jira.Url)
	if err != nil {
		return nil, fmt.Errorf("failed to create Jira client: %v", err)
	}

	return &Client{
		jiraClient:     jiraClient,
		IssueService:   newIssueService(jiraClient, issuesCache),
		WorklogService: newWorklogService(cfg.Jira.User, jiraClient, cfg),
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
