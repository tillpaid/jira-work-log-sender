package import_data

import (
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
)

func resolveWorkLogExcludedFromTimeHighlight(issueNumber string, config *resource.Config) bool {
	for _, excludedNumber := range config.IssueHighlight.ExcludedNumbers {
		if issueNumber == excludedNumber {
			return true
		}
	}

	return false
}
