package import_data

import (
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
)

func isExcludedFromTimeHighlight(issueNumber string, cfg *resource.Config) bool {
	for _, excludedIssue := range cfg.Highlighting.ExcludedIssues {
		if issueNumber == excludedIssue {
			return true
		}
	}

	return false
}
