package import_data

import (
	"github.com/tillpaid/paysera-log-time-golang/internal/resource"
	"strings"
)

func checkWorkLogAllowedTag(config *resource.Config, description string) bool {
	parts := strings.Split(description, "\n")
	if len(parts) == 0 {
		return false
	}

	firstPart := strings.ToUpper(strings.TrimSpace(strings.ReplaceAll(parts[0], "-", "")))

	for _, allowedTag := range config.AllowedTags {
		if firstPart == allowedTag {
			return true
		}
	}

	return false
}
