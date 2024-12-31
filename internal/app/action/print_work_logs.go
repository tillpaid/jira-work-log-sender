package action

import (
	"fmt"

	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/jira-work-log-sender/internal/jira"
	"github.com/tillpaid/jira-work-log-sender/internal/model"
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
	"github.com/tillpaid/jira-work-log-sender/internal/ui/element/table"
	"github.com/tillpaid/jira-work-log-sender/internal/ui/pages/page_work_logs"
)

type PrintWorkLogsAction struct {
	client *jira.Client
	window *goncurses.Window
	config *resource.Config
}

func NewPrintWorkLogsAction(client *jira.Client, window *goncurses.Window, config *resource.Config) *PrintWorkLogsAction {
	return &PrintWorkLogsAction{client: client, window: window, config: config}
}

func (a *PrintWorkLogsAction) Print(workLogs []model.WorkLog, rowSelector *model.RowSelector) (*table.Table, error) {
	t, err := page_work_logs.DrawWorkLogsTable(a.window, a.config, workLogs, rowSelector.Row)
	if err != nil {
		return t, err
	}

	return t, a.checkWorkLogs(t, workLogs)
}

func (a *PrintWorkLogsAction) UpdateSelectedRow(t *table.Table, rowSelector *model.RowSelector) {
	if len(t.Rows) == 0 {
		return
	}

	t.Rows[rowSelector.PreviousRow-1].IsSelected = false
	t.ReDrawRow(t.Rows[rowSelector.PreviousRow-1])

	t.Rows[rowSelector.Row-1].IsSelected = true
	t.ReDrawRow(t.Rows[rowSelector.Row-1])

	a.window.Refresh()
}

func (a *PrintWorkLogsAction) checkWorkLogs(t *table.Table, workLogs []model.WorkLog) error {
	for i, workLog := range workLogs {
		issueExists := a.client.IssueService.IsIssueExistsInCache(workLog.IssueNumber)
		if !issueExists {
			continue
		}

		t.ShowRow(i)
		a.window.Refresh()
	}

	for i, workLog := range workLogs {
		issueExists, err := a.client.IssueService.IsIssueExists(workLog.IssueNumber)
		if err != nil {
			return fmt.Errorf("impossible to check issue %s in jira: %s", workLog.IssueNumber, err)
		}
		if !issueExists {
			return fmt.Errorf("issue %s does not exist", workLog.IssueNumber)
		}

		t.ShowRow(i)
		a.window.Refresh()
	}

	return nil
}
