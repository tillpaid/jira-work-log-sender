package jira

import (
	"github.com/andygrunwald/go-jira"
)

type issueService struct {
	jiraClient *jira.Client
}

func newIssueService(jiraClient *jira.Client) *issueService {
	return &issueService{jiraClient: jiraClient}
}

func (i *issueService) IsIssueExists(issueID string) bool {
	_, _, err := i.jiraClient.Issue.Get(issueID, nil)

	return err == nil
}
