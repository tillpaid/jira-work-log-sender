package page_work_logs

import "github.com/tillpaid/paysera-log-time-golang/internal/model"

func GetBody(workLogs []model.WorkLog) []string {
	var rows []string
	workLogsTableWidth := model.NewWorkLogTableWidthWithCalculations(workLogs)

	for _, workLog := range workLogs {
		rows = append(rows, workLog.ToStringWithSpaces(workLogsTableWidth))
	}

	return rows
}
