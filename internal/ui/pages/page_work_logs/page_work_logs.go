package page_work_logs

import (
	"fmt"

	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/jira-work-log-sender/internal/model"
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
	"github.com/tillpaid/jira-work-log-sender/internal/ui"
	"github.com/tillpaid/jira-work-log-sender/internal/ui/element"
	"github.com/tillpaid/jira-work-log-sender/internal/ui/element/table"
	"github.com/tillpaid/jira-work-log-sender/internal/ui/utils"
)

const (
	pageName   = " Work Logs for Today "
	footerText = " Action keys: R-Reload | L-Send work logs | [Q/Space/Return/Esc]-Exit "
)

func DrawWorkLogsTable(window *goncurses.Window, config *resource.Config, workLogs []model.WorkLog, selectedRow int) (*table.Table, error) {
	if err := window.Clear(); err != nil {
		return nil, fmt.Errorf("error clearing window: %v", err)
	}

	height, width := window.MaxYX()
	workLogsTableWidth := model.NewWorkLogTableWidthWithCalculations(workLogs, width)

	t := table.NewTable(
		getHeader(workLogsTableWidth),
		getRows(workLogs, workLogsTableWidth, selectedRow),
		window,
	)

	element.DrawBox(window, height-2, width, pageName)
	t.Draw()
	drawTimeRow(window, 2, workLogs, config)
	drawFooter(window, height)

	window.Refresh()
	return t, nil
}

func drawTimeRow(window *goncurses.Window, x int, workLogs []model.WorkLog, config *resource.Config) {
	elements := getTimeRow(workLogs, config)
	window.Move(5+len(workLogs), x+1)

	for i, e := range elements {
		utils.ColorOn(window, e.Color)
		window.Printf(" %s ", e.Text)
		utils.ColorOff(window, e.Color)

		if i < len(elements)-1 {
			window.AddChar(goncurses.ACS_VLINE)
		}
	}
}

func drawFooter(window *goncurses.Window, height int) {
	window.MovePrint(height-1, 2, footerText)
}

func getHeader(workLogsTableWidth *model.WorkLogTableWidth) *table.Header {
	columns := []*table.Column{
		{"Name", workLogsTableWidth.HeaderText, 0, 0},
		{"Time", workLogsTableWidth.OriginalTime, 0, 0},
		{"MTime", workLogsTableWidth.ModifiedTime, 0, 0},
		{"Issue", workLogsTableWidth.IssueNumber, 0, 0},
		{"Description", workLogsTableWidth.Description, 0, 0},
	}

	return table.NewHeader(columns, 3)
}

func getRows(workLogs []model.WorkLog, workLogsTableWidth *model.WorkLogTableWidth, selectedRow int) []*table.Row {
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
		row := table.NewRow(columns, i+1+3, isSelected, 3, false)

		if log.ModifyTimeDisabled {
			row.Color = ui.CyanOnBlack
		}

		rows = append(rows, row)
	}

	return rows
}
