package import_data

import (
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
)

func containAllowedTag(config *resource.Config, tag string) bool {
	if len(config.Tags.Allowed) == 0 {
		return true
	}

	for _, allowedTag := range config.Tags.Allowed {
		if tag == allowedTag {
			return true
		}
	}

	return false
}
