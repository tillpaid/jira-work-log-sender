package page_send_work_logs

import (
	"fmt"
	"strconv"

	"github.com/tillpaid/paysera-log-time-golang/internal/model"
	"github.com/tillpaid/paysera-log-time-golang/internal/service"
)

func GetBody(workLogs []model.WorkLog, width *model.WorkLogTableWidth) []string {
	var table []string

	for _, workLog := range workLogs {
		table = append(table, buildRow(workLog, width))
	}

	return table
}

func buildRow(workLog model.WorkLog, width *model.WorkLogTableWidth) string {
	return fmt.Sprintf(
		"%s %s | %s |",
		service.GetTextWithSpaces(strconv.Itoa(workLog.Number)+".", width.Number),
		service.GetTextWithSpaces(workLog.IssueNumber, width.IssueNumber),
		service.GetTextWithSpaces(workLog.ModifiedTime.String(), width.ModifiedTime),
	)
}
