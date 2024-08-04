package import_data

import (
	"github.com/tillpaid/paysera-log-time-golang/internal/resource"
	"strings"
)

func checkWorkLogAllowedTag(config *resource.Config, description string) bool {
	firstPart := strings.ToLower(strings.TrimSpace(strings.ReplaceAll(description, "-", "")))

	for _, allowedTag := range config.AllowedTags {
		if firstPart == allowedTag {
			return true
		}
	}

	return false
}
