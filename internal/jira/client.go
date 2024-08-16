package jira

import (
	"fmt"

	"github.com/andygrunwald/go-jira"
	"github.com/tillpaid/paysera-log-time-golang/internal/model"
	"github.com/tillpaid/paysera-log-time-golang/internal/resource"
)

type IssueServiceInterface interface {
	IsIssueExists(issueID string) bool
}

type WorkLogServiceInterface interface {
	GetSpentTime(issueID string) string
	SendWorkLog(workLog model.WorkLog) error
}

type Client struct {
	IssueService   IssueServiceInterface
	WorkLogService WorkLogServiceInterface
}

func NewClient(config *resource.Config) (*Client, error) {
	tp := jira.BasicAuthTransport{
		Username: config.Jira.Username,
		Password: config.Jira.ApiToken,
	}

	jiraClient, err := jira.NewClient(tp.Client(), config.Jira.BaseUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create Jira client: %v", err)
	}

	// Ping Jira to check if the authentication is successful
	_, response, err := jiraClient.User.GetSelf()
	if err != nil {
		newError := fmt.Errorf("failed to authenticate with Jira: %v", err)
		if response != nil {
			newError = fmt.Errorf("failed to authenticate with Jira. Response code is %d", response.StatusCode)
		}

		return nil, newError
	}

	return &Client{
		IssueService:   newIssueService(jiraClient),
		WorkLogService: newWorkLogService(config.Jira.Username, jiraClient),
	}, nil
}
