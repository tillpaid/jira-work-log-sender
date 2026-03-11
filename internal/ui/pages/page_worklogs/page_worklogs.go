package page_worklogs

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

func DrawWorklogsTable(window *goncurses.Window, cfg *resource.Config, worklogs []model.Worklog, selectedRow int) (*table.Table, error) {
	if err := window.Clear(); err != nil {
		return nil, fmt.Errorf("error clearing window: %v", err)
	}

	height, width := window.MaxYX()
	worklogsTableWidth := model.NewWorklogTableWidthWithCalculations(worklogs, width)

	t := table.NewTable(
		getHeader(worklogsTableWidth),
		getRows(worklogs, worklogsTableWidth, selectedRow),
		window,
	)

	element.DrawBox(window, height-2, width, pageName)
	t.Draw()
	drawTimeRow(window, 2, width, worklogs, cfg)
	drawFooter(window, height)

	window.Refresh()
	return t, nil
}

func drawTimeRow(window *goncurses.Window, x int, width int, worklogs []model.Worklog, cfg *resource.Config) {
	timeRow := getTimeRow(worklogs, cfg)
	useSeparateLines := timeRow.GetTotalTextLen(3) >= width-6

	y := 5 + len(worklogs)
	window.Move(y, x+1)

	for i, e := range timeRow.Elements {
		utils.ColorOn(window, e.Color)
		window.Printf(" %s ", e.Text)
		utils.ColorOff(window, e.Color)

		if !useSeparateLines && i < len(timeRow.Elements)-1 {
			window.AddChar(goncurses.ACS_VLINE)
		}

		if useSeparateLines {
			window.Move(y+i+1, x+1)
		}
	}
}

func drawFooter(window *goncurses.Window, height int) {
	window.MovePrint(height-1, 2, footerText)
}

func getHeader(worklogsTableWidth *model.WorklogTableWidth) *table.Header {
	columns := []*table.Column{
		{"Name", worklogsTableWidth.HeaderText, 0, 0},
		{"Time", worklogsTableWidth.OriginalTime, 0, 0},
		{"MTime", worklogsTableWidth.ModifiedTime, 0, 0},
		{"Issue", worklogsTableWidth.IssueNumber, 0, 0},
		{"Description", worklogsTableWidth.Description, 0, 0},
	}

	return table.NewHeader(columns, 3)
}

func getRows(worklogs []model.Worklog, worklogsTableWidth *model.WorklogTableWidth, selectedRow int) []*table.Row {
	var rows []*table.Row

	for i, log := range worklogs {
		columns := []*table.Column{
			{log.GetHeader(), worklogsTableWidth.HeaderText, 0, 0},
			{log.OriginalTime.String(), worklogsTableWidth.OriginalTime, 0, 0},
			{log.ModifiedTime.String(), worklogsTableWidth.ModifiedTime, 0, 0},
			{log.IssueNumber, worklogsTableWidth.IssueNumber, 0, 0},
			{log.Description, worklogsTableWidth.Description, 0, 0},
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
