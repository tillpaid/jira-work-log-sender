package page_work_logs

import (
	"fmt"

	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/model"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui/element"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui/element/table"
)

const (
	pageName   = " Work Logs for Today "
	footerText = " Action keys: R-Reload | L-Send work logs | [Q/Space/Return/Esc]-Exit "
)

func DrawWorkLogsTable(screen *goncurses.Window, workLogs []model.WorkLog, selectedRow int, showRows bool) (*table.Table, error) {
	if err := screen.Clear(); err != nil {
		return nil, fmt.Errorf("error clearing screen: %v", err)
	}

	height, width := screen.MaxYX()
	workLogsTableWidth := model.NewWorkLogTableWidthWithCalculations(workLogs, width)

	t := table.NewTable(
		getHeader(workLogsTableWidth),
		getRows(workLogs, workLogsTableWidth, selectedRow, showRows),
		screen,
	)

	element.DrawBox(screen, height-2, width, pageName)
	t.Draw()
	drawTimeRow(screen, 2, workLogs)
	drawFooter(screen, height)

	screen.Refresh()
	return t, nil
}

func drawTimeRow(screen *goncurses.Window, x int, workLogs []model.WorkLog) {
	elements := getTimeRow(workLogs)
	screen.Move(5+len(workLogs), x+1)

	for i, e := range elements {
		screen.Printf(" %s ", e)

		if i < len(elements)-1 {
			screen.AddChar(goncurses.ACS_VLINE)
		}
	}
}

func drawFooter(screen *goncurses.Window, height int) {
	screen.MovePrint(height-1, 2, footerText)
}

func getHeader(workLogsTableWidth *model.WorkLogTableWidth) *table.Header {
	columns := []*table.Column{
		{"Name", workLogsTableWidth.HeaderText, 0, 0},
		{"T", workLogsTableWidth.OriginalTime, 0, 0},
		{"MT", workLogsTableWidth.ModifiedTime, 0, 0},
		{"Issue", workLogsTableWidth.IssueNumber, 0, 0},
		{"Description", workLogsTableWidth.Description, 0, 0},
	}

	return table.NewHeader(columns, 3)
}

func getRows(workLogs []model.WorkLog, workLogsTableWidth *model.WorkLogTableWidth, selectedRow int, showRows bool) []*table.Row {
	var rows []*table.Row

	for i, log := range workLogs {
		columns := []*table.Column{
			{log.GetHeader(), workLogsTableWidth.HeaderText, 0, 0},
			{log.OriginalTime.String(), workLogsTableWidth.OriginalTime, 0, 0},
			{log.ModifiedTime.String(), workLogsTableWidth.ModifiedTime, 0, 0},
			{log.IssueNumber, workLogsTableWidth.IssueNumber, 0, 0},
			{log.Description, workLogsTableWidth.Description, 0, 0},
		}

		isSelected := i+1 == selectedRow
		rows = append(rows, table.NewRow(columns, i+1+3, isSelected, 3, showRows))
	}

	return rows
}
