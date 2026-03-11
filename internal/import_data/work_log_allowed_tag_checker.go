package import_data

import (
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
)

func containAllowedTag(cfg *resource.Config, tag string) bool {
	if len(cfg.Tags.Allowed) == 0 {
		return true
	}

	for _, allowedTag := range cfg.Tags.Allowed {
		if tag == allowedTag {
			return true
		}
	}

	return false
}
