package import_data

import (
	"strings"

	"github.com/tillpaid/paysera-log-time-golang/internal/resource"
)

func resolveWorkLogExcludedFromTimeHighlight(issueNumber string, config *resource.Config) bool {
	for _, value := range config.ExcludedFromSpentTimeHighlight {
		if strings.HasPrefix(issueNumber, value) {
			return true
		}
	}

	return false
}
