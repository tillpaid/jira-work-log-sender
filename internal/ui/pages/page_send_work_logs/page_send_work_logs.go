package page_send_work_logs

import (
	"fmt"
	"strconv"

	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/model"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui/element"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui/element/table"
)

const (
	pageName   = " Send Work Logs "
	footerText = " Action keys: R-Reload | [Q/Space/Return/Esc]-Exit "
)

func DrawSendWorkLogsPage(screen *goncurses.Window, workLogs []model.WorkLog) (*table.Table, error) {
	if err := screen.Clear(); err != nil {
		return nil, fmt.Errorf("error clearing screen: %v", err)
	}

	height, width := screen.MaxYX()
	workLogsTableWidth := model.NewWorkLogTableWidthWithCalculations(workLogs, width)

	t := table.NewTable(
		getHeader(workLogsTableWidth),
		getRows(workLogs, workLogsTableWidth),
		screen,
	)

	element.DrawBox(screen, height-2, width, pageName)
	t.Draw()
	drawFooter(screen, height)

	screen.Refresh()
	return t, nil
}

func drawFooter(screen *goncurses.Window, height int) {
	screen.MovePrint(height-1, 2, footerText)
}

func getHeader(workLogsTableWidth *model.WorkLogTableWidth) *table.Header {
	columns := []*table.Column{
		{"#", workLogsTableWidth.Number, 0, 0},
		{"Issue", workLogsTableWidth.IssueNumber, 0, 0},
		{"MT", workLogsTableWidth.ModifiedTime, 0, 0},
		{"Send status", workLogsTableWidth.SendStatus, 0, 0},
		{"Total time", workLogsTableWidth.TotalTime, 0, 0},
	}

	return table.NewHeader(columns, 3)
}

func getRows(workLogs []model.WorkLog, workLogsTableWidth *model.WorkLogTableWidth) []*table.Row {
	var rows []*table.Row

	for i, log := range workLogs {
		columns := []*table.Column{
			{strconv.Itoa(log.Number), workLogsTableWidth.Number, 0, 0},
			{log.IssueNumber, workLogsTableWidth.IssueNumber, 0, 0},
			{log.ModifiedTime.String(), workLogsTableWidth.ModifiedTime, 0, 0},
			{"", workLogsTableWidth.SendStatus, 0, 0},
			{"", workLogsTableWidth.TotalTime, 0, 0},
		}

		rows = append(rows, table.NewRow(columns, i+1+3, false, 3, true))
	}

	return rows
}
