package page_dump_work_logs

import (
	"fmt"
	"github.com/tillpaid/paysera-log-time-golang/internal/resource"
)

func GetBody(config *resource.Config) []string {
	return []string{
		"Successfully dumped work logs to file",
		"",
		fmt.Sprintf("File path: %s", config.OutputShellFile),
	}
}
