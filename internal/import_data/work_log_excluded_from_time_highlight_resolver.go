package import_data

import (
	"strings"

	"github.com/tillpaid/jira-work-log-sender/internal/resource"
)

func resolveWorkLogExcludedFromTimeHighlight(issueNumber string, config *resource.Config) bool {
	for _, value := range config.IssueHighlight.ExcludedNumbers {
		if strings.HasPrefix(issueNumber, value) {
			return true
		}
	}

	return false
}
