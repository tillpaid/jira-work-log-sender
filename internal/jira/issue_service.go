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
	jiraClient  *jira.Client
	issuesCache *cache.IssuesCache
}

func newIssueService(jiraClient *jira.Client, issuesCache *cache.IssuesCache) *issueService {
	return &issueService{
		jiraClient:  jiraClient,
		issuesCache: issuesCache,
	}
}

func (i *issueService) IsIssueExists(issueKey string) (bool, error) {
	if i.issuesCache.IsIssueExists(issueKey) {
		return true, nil
	}

	issue, response, jiraError := i.jiraClient.Issue.Get(issueKey, nil)
	if jiraError != nil {
		if response != nil {
			if response.StatusCode == http.StatusNotFound {
				return false, errors.New("issue not found in Jira")
			}

			return false, fmt.Errorf("got %d status code", response.StatusCode)
		}

		if strings.Contains(jiraError.Error(), "no such host") {
			return false, errors.New("cannot resolve Jira host")
		}

		return false, jiraError
	}

	if err := i.issuesCache.SaveIssue(issue.Key, issue.ID); err != nil {
		return false, err
	}

	return true, nil
}

func (i *issueService) IsIssueExistsInCache(issueKey string) bool {
	return i.issuesCache.IsIssueExists(issueKey)
}
