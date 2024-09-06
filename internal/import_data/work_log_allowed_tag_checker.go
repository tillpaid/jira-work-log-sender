package import_data

import (
	"github.com/tillpaid/paysera-log-time-golang/internal/resource"
)

func checkWorkLogAllowedTag(config *resource.Config, tag string) bool {
	for _, allowedTag := range config.AllowedTags {
		if tag == allowedTag {
			return true
		}
	}

	return false
}
