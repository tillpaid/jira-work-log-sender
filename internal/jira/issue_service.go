package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/tillpaid/paysera-log-time-golang/internal/cache"
)

type issueService struct {
	jiraClient           *jira.Client
	issuesExistenceCache *cache.IssuesExistenceCache
}

func newIssueService(jiraClient *jira.Client, issuesExistenceCache *cache.IssuesExistenceCache) *issueService {
	return &issueService{jiraClient: jiraClient, issuesExistenceCache: issuesExistenceCache}
}

func (i *issueService) IsIssueExists(issueID string) (bool, error) {
	if i.issuesExistenceCache.IsExists(issueID) {
		return true, nil
	}

	if _, _, jiraError := i.jiraClient.Issue.Get(issueID, nil); jiraError != nil {
		return false, nil
	}

	if err := i.issuesExistenceCache.SaveExists(issueID); err != nil {
		return false, err
	}

	return true, nil
}
