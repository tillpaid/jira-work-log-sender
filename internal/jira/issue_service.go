package jira

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/andygrunwald/go-jira"
	"github.com/tillpaid/jira-work-log-sender/internal/cache"
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

	_, response, jiraError := i.jiraClient.Issue.Get(issueID, nil)
	if jiraError != nil {
		if response != nil {
			if response.StatusCode == http.StatusNotFound {
				return false, errors.New("issue not found in Jira")
			}

			return false, fmt.Errorf("got %d status code", response.StatusCode)
		}

		if strings.Contains(jiraError.Error(), "no such host") {
			return false, errors.New("cannot resolve Jura host")
		}

		return false, jiraError
	}

	if err := i.issuesExistenceCache.SaveExists(issueID); err != nil {
		return false, err
	}

	return true, nil
}

func (i *issueService) IsIssueExistsInCache(issueID string) bool {
	return i.issuesExistenceCache.IsExists(issueID)
}
