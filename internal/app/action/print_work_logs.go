package action

import (
	"fmt"

	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/jira"
	"github.com/tillpaid/paysera-log-time-golang/internal/model"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui/element/table"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui/pages/page_work_logs"
)

type PrintWorkLogsAction struct {
	client *jira.Client
	screen *goncurses.Window
}

func NewPrintWorkLogsAction(client *jira.Client, screen *goncurses.Window) *PrintWorkLogsAction {
	return &PrintWorkLogsAction{client: client, screen: screen}
}

func (a *PrintWorkLogsAction) Print(workLogs []model.WorkLog, rowSelector *model.RowSelector) (*table.Table, error) {
	t, err := page_work_logs.DrawWorkLogsTable(a.screen, workLogs, rowSelector.Row)
	if err != nil {
		return t, err
	}

	return t, a.checkWorkLogs(t, workLogs)
}

func (a *PrintWorkLogsAction) UpdateSelectedRow(t *table.Table, rowSelector *model.RowSelector) {
	t.Rows[rowSelector.PreviousRow-1].IsSelected = false
	t.ReDrawRow(t.Rows[rowSelector.PreviousRow-1])

	t.Rows[rowSelector.Row-1].IsSelected = true
	t.ReDrawRow(t.Rows[rowSelector.Row-1])

	a.screen.Refresh()
}

func (a *PrintWorkLogsAction) checkWorkLogs(t *table.Table, workLogs []model.WorkLog) error {
	for i, workLog := range workLogs {
		issueExists, err := a.client.IssueService.IsIssueExists(workLog.IssueNumber)
		if err != nil {
			return fmt.Errorf("impossible to check issue %s in jira: %s", workLog.IssueNumber, err)
		}
		if !issueExists {
			return fmt.Errorf("issue %s does not exist", workLog.IssueNumber)
		}

		row := t.Rows[i]
		row.ShowText = true

		t.ReDrawRow(row)
		a.screen.Refresh()
	}

	return nil
}
