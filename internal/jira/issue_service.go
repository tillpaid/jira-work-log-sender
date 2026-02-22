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

func (i *issueService) GetIssueID(issueKey string) (string, error) {
	issueCacheItem, ok := i.issuesCache.GetIssue(issueKey)
	if ok {
		return issueCacheItem.IssueID, nil
	}

	issue, err := i.getIssue(issueKey)
	if err != nil {
		return "", err
	}

	if err := i.issuesCache.SaveIssue(issue.Key, issue.ID); err != nil {
		return "", err
	}

	return issue.ID, nil
}

func (i *issueService) IsIssueExists(issueKey string) (bool, error) {
	if i.issuesCache.IsIssueExists(issueKey) {
		return true, nil
	}

	issue, err := i.getIssue(issueKey)
	if err != nil {
		return false, err
	}

	if err := i.issuesCache.SaveIssue(issue.Key, issue.ID); err != nil {
		return false, err
	}

	return true, nil
}

func (i *issueService) IsIssueExistsInCache(issueKey string) bool {
	return i.issuesCache.IsIssueExists(issueKey)
}

func (i *issueService) getIssue(issueKey string) (*jira.Issue, error) {
	issue, response, jiraError := i.jiraClient.Issue.Get(issueKey, nil)

	if jiraError != nil {
		if response != nil {
			if response.StatusCode == http.StatusNotFound {
				return nil, errors.New("issue not found in Jira")
			}

			return nil, fmt.Errorf("got %d status code", response.StatusCode)
		}

		if strings.Contains(jiraError.Error(), "no such host") {
			return nil, errors.New("cannot resolve Jira host")
		}

		return nil, jiraError
	}

	return issue, nil
}
