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

func (a *PrintWorkLogsAction) Print(workLogs []model.WorkLog, selectedRow int, needCheck bool) error {
	t, err := page_work_logs.DrawWorkLogsTable(a.screen, workLogs, selectedRow, !needCheck)
	if err != nil {
		return err
	}

	if needCheck {
		return a.checkWorkLogs(t, workLogs)
	}

	return nil
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
