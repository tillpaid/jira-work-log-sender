package export_data

import (
	"fmt"
	"github.com/tillpaid/paysera-log-time-golang/internal/model"
	"strings"
)

func convertWorkLogsToFileContent(workLogs []model.WorkLog) string {
	builder := strings.Builder{}

	builder.WriteString("#!/bin/bash\n\n")

	for _, workLog := range workLogs {
		firstLine := fmt.Sprintf("echo \"Command number: %d\" && ", workLog.Number)
		secondLine := fmt.Sprintf("please log-time %s \"%s\" \"%s\"\n\n",
			workLog.IssueNumber, workLog.ModifiedTime.String(), workLog.Description)

		builder.WriteString(firstLine)
		builder.WriteString(secondLine)
	}

	return builder.String()
}
