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

	height, width := screen.MaxYX()
	tableRows := buildTableRows(workLogs, height, width)

	for i, line := range tableRows {
		screen.MovePrint(i, 0, prepareRow(line, width))
	}

	screen.Refresh()
	return nil
}

func buildTableRows(workLogs []model.WorkLog, height int, width int) []string {
	delimiter := getDelimiter(width)
	var rows []string

	header := page_work_logs.GetHeader(delimiter)
	body := page_work_logs.GetBody(workLogs)
	timeRow := page_work_logs.GetTimeRow(workLogs, delimiter)
	footer := page_work_logs.GetFooter(delimiter)

	otherRowsLen := len(header) + len(timeRow) + len(footer)
	body = cutBody(body, otherRowsLen, height, width)

	rows = append(rows, header...)
	rows = append(rows, body...)
	rows = append(rows, timeRow...)
	rows = append(rows, footer...)

	return rows
}
