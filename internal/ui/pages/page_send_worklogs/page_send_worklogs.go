package page_send_worklogs

import (
	"fmt"
	"strconv"

	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/jira-work-log-sender/internal/model"
	"github.com/tillpaid/jira-work-log-sender/internal/ui/element"
	"github.com/tillpaid/jira-work-log-sender/internal/ui/element/table"
)

const (
	pageName   = " Send Work Logs "
	footerText = " Action keys: R-Reload | [Q/Space/Return/Esc]-Exit "
)

func DrawSendWorklogsPage(window *goncurses.Window, worklogs []model.Worklog) (*table.Table, error) {
	if err := window.Clear(); err != nil {
		return nil, fmt.Errorf("error clearing window: %v", err)
	}

	height, width := window.MaxYX()
	worklogsTableWidth := model.NewWorklogTableWidthWithCalculations(worklogs, width)

	t := table.NewTable(
		getHeader(worklogsTableWidth),
		getRows(worklogs, worklogsTableWidth),
		window,
	)

	element.DrawBox(window, height-2, width, pageName)
	t.Draw()
	drawFooter(window, height)

	window.Refresh()
	return t, nil
}

func drawFooter(window *goncurses.Window, height int) {
	window.MovePrint(height-1, 2, footerText)
}

func getHeader(worklogsTableWidth *model.WorklogTableWidth) *table.Header {
	columns := []*table.Column{
		{"#", worklogsTableWidth.Number, 0, 0},
		{"Issue", worklogsTableWidth.IssueNumber, 0, 0},
		{"MTime", worklogsTableWidth.ModifiedTime, 0, 0},
		{"Send status", worklogsTableWidth.SendStatus, 0, 0},
		{"Total time", worklogsTableWidth.TotalTime, 0, 0},
	}

	return table.NewHeader(columns, 3)
}

func getRows(worklogs []model.Worklog, worklogsTableWidth *model.WorklogTableWidth) []*table.Row {
	var rows []*table.Row

	for i, log := range worklogs {
		columns := []*table.Column{
			{strconv.Itoa(log.Number), worklogsTableWidth.Number, 0, 0},
			{log.IssueNumber, worklogsTableWidth.IssueNumber, 0, 0},
			{log.ModifiedTime.String(), worklogsTableWidth.ModifiedTime, 0, 0},
			{"", worklogsTableWidth.SendStatus, 0, 0},
			{"", worklogsTableWidth.TotalTime, 0, 0},
		}

		rows = append(rows, table.NewRow(columns, i+1+3, false, 3, true))
	}

	return rows
}
