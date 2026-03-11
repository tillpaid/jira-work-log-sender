package action

import (
	"fmt"

	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/jira-work-log-sender/internal/jira"
	"github.com/tillpaid/jira-work-log-sender/internal/model"
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
	"github.com/tillpaid/jira-work-log-sender/internal/ui/element/table"
	"github.com/tillpaid/jira-work-log-sender/internal/ui/pages/page_worklogs"
)

type PrintWorklogsAction struct {
	client *jira.Client
	window *goncurses.Window
	cfg    *resource.Config
}

func NewPrintWorklogsAction(client *jira.Client, window *goncurses.Window, cfg *resource.Config) *PrintWorklogsAction {
	return &PrintWorklogsAction{client: client, window: window, cfg: cfg}
}

func (a *PrintWorklogsAction) Print(worklogs []model.Worklog, rowSelector *model.RowSelector) (*table.Table, error) {
	t, err := page_worklogs.DrawWorklogsTable(a.window, a.cfg, worklogs, rowSelector.Row)
	if err != nil {
		return t, err
	}

	return t, a.checkWorklogs(t, worklogs)
}

func (a *PrintWorklogsAction) UpdateSelectedRow(t *table.Table, rowSelector *model.RowSelector) {
	if len(t.Rows) == 0 {
		return
	}

	t.Rows[rowSelector.PreviousRow-1].IsSelected = false
	t.ReDrawRow(t.Rows[rowSelector.PreviousRow-1])

	t.Rows[rowSelector.Row-1].IsSelected = true
	t.ReDrawRow(t.Rows[rowSelector.Row-1])

	a.window.Refresh()
}

func (a *PrintWorklogsAction) checkWorklogs(t *table.Table, worklogs []model.Worklog) error {
	for i, worklog := range worklogs {
		issueExists := a.client.IssueService.IsIssueExistsInCache(worklog.IssueNumber)
		if !issueExists {
			continue
		}

		t.ShowRow(i)
		a.window.Refresh()
	}

	for i, worklog := range worklogs {
		issueExists, err := a.client.IssueService.IsIssueExists(worklog.IssueNumber)
		if err != nil {
			return fmt.Errorf("impossible to check issue %s in jira: %s", worklog.IssueNumber, err)
		}
		if !issueExists {
			return fmt.Errorf("issue %s does not exist", worklog.IssueNumber)
		}

		t.ShowRow(i)
		a.window.Refresh()
	}

	return nil
}
