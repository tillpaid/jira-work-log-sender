package page_work_logs

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tillpaid/paysera-log-time-golang/internal/model"
	"github.com/tillpaid/paysera-log-time-golang/internal/service"
)

func GetBody(workLogs []model.WorkLog, selectedRow int) []string {
	var rows []string
	workLogsTableWidth := model.NewWorkLogTableWidthWithCalculations(workLogs)

	for i, workLog := range workLogs {
		rows = append(rows, buildRow(workLog, workLogsTableWidth, selectedRow == i+1))
	}

	return rows
}

func buildRow(workLog model.WorkLog, width *model.WorkLogTableWidth, isSelected bool) string {
	arrow := "   "
	if isSelected {
		arrow = "-->"
	}

	return fmt.Sprintf(
		"%s | %s %s | %s | %s | %s",
		arrow,
		service.GetTextWithSpaces(strconv.Itoa(workLog.Number)+".", width.Number),
		service.GetTextWithSpaces(workLog.OriginalTime.String(), width.OriginalTime),
		service.GetTextWithSpaces(workLog.ModifiedTime.String(), width.ModifiedTime),
		service.GetTextWithSpaces(workLog.IssueNumber, width.IssueNumber),
		strings.ReplaceAll(workLog.Description, "\n", " "),
	)
}
