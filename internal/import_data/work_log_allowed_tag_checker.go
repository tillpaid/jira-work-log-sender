package import_data

import (
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
)

func containAllowedTag(config *resource.Config, tag string) bool {
	if len(config.AllowedTags) == 0 {
		return true
	}

	for _, allowedTag := range config.AllowedTags {
		if tag == allowedTag {
			return true
		}
	}

	return false
}
