package pages

import (
	"fmt"

	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/model"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui/pages/page_send_work_logs"
)

func DrawSendWorkLogsPage(screen *goncurses.Window, workLogs []model.WorkLog, valuesWidth *model.WorkLogTableWidth) error {
	if err := screen.Clear(); err != nil {
		return fmt.Errorf("error clearing screen: %v", err)
	}

	height, width := screen.MaxYX()
	tableRows := buildSendWorkLogsTableRows(workLogs, valuesWidth, height, width)

	for i, line := range tableRows {
		screen.MovePrint(i, 0, prepareRow(line, width))
	}

	screen.Refresh()
	return nil
}

func buildSendWorkLogsTableRows(workLogs []model.WorkLog, valuesWidth *model.WorkLogTableWidth, height int, width int) []string {
	delimiter := getDelimiter(width)
	var rows []string

	header := page_send_work_logs.GetHeader(delimiter)
	body := page_send_work_logs.GetBody(workLogs, valuesWidth)
	footer := page_send_work_logs.GetFooter(delimiter)

	otherRowsLen := len(header) + len(footer)
	body = cutBody(body, otherRowsLen, height, width)

	rows = append(rows, header...)
	rows = append(rows, body...)
	rows = append(rows, footer...)

	return rows
}
