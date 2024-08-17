package page_work_logs

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tillpaid/paysera-log-time-golang/internal/model"
	"github.com/tillpaid/paysera-log-time-golang/internal/service"
)

func GetBody(workLogs []model.WorkLog) []string {
	var rows []string
	workLogsTableWidth := model.NewWorkLogTableWidthWithCalculations(workLogs)

	for _, workLog := range workLogs {
		rows = append(rows, buildRow(workLog, workLogsTableWidth))
	}

	return rows
}

func buildRow(workLog model.WorkLog, width *model.WorkLogTableWidth) string {
	return fmt.Sprintf(
		"%s %s | %s | %s | %s",
		service.GetTextWithSpaces(strconv.Itoa(workLog.Number)+".", width.Number),
		service.GetTextWithSpaces(workLog.OriginalTime.String(), width.OriginalTime),
		service.GetTextWithSpaces(workLog.ModifiedTime.String(), width.ModifiedTime),
		service.GetTextWithSpaces(workLog.IssueNumber, width.IssueNumber),
		strings.ReplaceAll(workLog.Description, "\n", " "),
	)
}
