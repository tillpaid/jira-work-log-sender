package import_data

import (
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
)

func isExcludedFromTimeHighlight(issueNumber string, config *resource.Config) bool {
	for _, excludedIssue := range config.Highlighting.ExcludedIssues {
		if issueNumber == excludedIssue {
			return true
		}
	}

	return false
}
