package pages

import (
	"fmt"
	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/model"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui/pages/page_work_logs"
)

func DrawWorkLogsTable(screen *goncurses.Window, workLogs []model.WorkLog) error {
	if err := screen.Clear(); err != nil {
		return fmt.Errorf("error clearing screen: %v", err)
	}

	_, width := screen.MaxYX()
	tableRows := buildTableRows(workLogs, width)

	for i, line := range tableRows {
		screen.MovePrint(i, 0, prepareRow(line, width))
	}

	screen.Refresh()
	return nil
}

func buildTableRows(workLogs []model.WorkLog, width int) []string {
	delimiter := getDelimiter(width)
	var rows []string

	rows = append(rows, page_work_logs.GetHeader(delimiter)...)
	rows = append(rows, page_work_logs.GetBody(workLogs)...)
	rows = append(rows, page_work_logs.GetTimeRow(workLogs, delimiter)...)
	rows = append(rows, page_work_logs.GetFooter(delimiter)...)

	return rows
}
