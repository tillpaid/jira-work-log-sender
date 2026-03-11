package app

import (
	"fmt"

	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/jira-work-log-sender/internal/app/action"
	"github.com/tillpaid/jira-work-log-sender/internal/clipboard"
	"github.com/tillpaid/jira-work-log-sender/internal/import_data"
	"github.com/tillpaid/jira-work-log-sender/internal/jira"
	"github.com/tillpaid/jira-work-log-sender/internal/model"
	"github.com/tillpaid/jira-work-log-sender/internal/resource"
	"github.com/tillpaid/jira-work-log-sender/internal/service"
	"github.com/tillpaid/jira-work-log-sender/internal/ui"
	"github.com/tillpaid/jira-work-log-sender/internal/ui/element/table"
)

const (
	actionReload = iota
	actionSend
	actionNextRow
	actionPrevRow
	actionFirstRow
	actionLastRow
	actionCopy
	actionCopyWithoutExit
	actionToggleModifyTime
	actionToggleAllModifyTime
	actionQuit
)

type Application struct {
	window    *goncurses.Window
	client    *jira.Client
	userInput *UserInput
	actions   *action.Actions
	cfg       *resource.Config

	table        *table.Table
	worklogs     []model.Worklog
	selector     *model.RowSelector
	worklogsSent bool
}

func NewApplication(window *goncurses.Window, client *jira.Client, input *UserInput, actions *action.Actions, cfg *resource.Config) *Application {
	selector := model.NewRowSelector(0)

	return &Application{
		window:       window,
		client:       client,
		userInput:    input,
		cfg:          cfg,
		actions:      actions,
		selector:     selector,
		worklogsSent: false,
	}
}

func (a *Application) Start() error {
	if err := a.loadWorklogs(); err != nil {
		return err
	}

	if err := a.printTable(); err != nil {
		return err
	}

	if err := a.setIssueIDs(); err != nil {
		return err
	}

	handleResize(&a.window, &a.table, a.selector, a.actions, &a.worklogs)

	for {
		switch a.userInput.WaitForAction() {
		case actionReload:
			if err := a.processActionReload(); err != nil {
				return err
			}
		case actionSend:
			if err := a.processActionSend(); err != nil {
				return err
			}
		case actionNextRow:
			a.selector.NextRow()
			a.actions.PrintWorklogs.UpdateSelectedRow(a.table, a.selector)
		case actionPrevRow:
			a.selector.PrevRow()
			a.actions.PrintWorklogs.UpdateSelectedRow(a.table, a.selector)
		case actionFirstRow:
			a.selector.FirstRow()
			a.actions.PrintWorklogs.UpdateSelectedRow(a.table, a.selector)
		case actionLastRow:
			a.selector.LastRow()
			a.actions.PrintWorklogs.UpdateSelectedRow(a.table, a.selector)
		case actionCopy:
			return a.processActionCopy(true)
		case actionCopyWithoutExit:
			if err := a.processActionCopy(false); err != nil {
				return err
			}
		case actionToggleModifyTime:
			if err := a.processActionToggleModifyTime(false); err != nil {
				return err
			}
		case actionToggleAllModifyTime:
			if err := a.processActionToggleModifyTime(true); err != nil {
				return err
			}
		case actionQuit:
			ui.EndWindow()
			return nil
		}
	}
}

func (a *Application) processActionReload() error {
	if err := a.loadWorklogs(); err != nil {
		return err
	}

	if err := a.printTable(); err != nil {
		return err
	}

	return a.setIssueIDs()
}

func (a *Application) processActionSend() error {
	if a.worklogsSent {
		return nil
	}

	a.worklogsSent = true
	return a.actions.SendWorklogs.Send(a.worklogs)
}

func (a *Application) processActionCopy(exit bool) error {
	if len(a.worklogs) == 0 {
		return nil
	}

	if err := clipboard.CopyToClipboard(a.worklogs[a.selector.Row-1].HeaderText); err != nil {
		return err
	}

	if exit {
		ui.EndWindow()
	}

	return nil
}

func (a *Application) processActionToggleModifyTime(all bool) error {
	if len(a.worklogs) == 0 || !a.cfg.TimeAdjustment.Enabled {
		return nil
	}

	for i := range a.worklogs {
		if all || i == a.selector.Row-1 {
			a.worklogs[i].ToggleModifyTime()
		}
	}

	a.worklogs = service.ModifyWorklogsTime(a.worklogs, a.cfg)
	return a.printTable()
}

func (a *Application) loadWorklogs() error {
	worklogs, err := import_data.ParseWorklogs(a.cfg, a.worklogs)
	if err != nil {
		return err
	}

	a.worklogs = worklogs
	a.selector.Update(len(worklogs))

	return nil
}

func (a *Application) setIssueIDs() error {
	for i := range a.worklogs {
		worklog := &a.worklogs[i]

		issueID, err := a.client.IssueService.GetIssueID(worklog.IssueNumber)
		if err != nil {
			return fmt.Errorf("failed to get issue ID for issue number %s: %w", worklog.IssueNumber, err)
		}

		worklog.IssueID = issueID
	}

	return nil
}

func (a *Application) printTable() error {
	t, err := a.actions.PrintWorklogs.Print(a.worklogs, a.selector)
	if err != nil {
		return err
	}

	a.table = t
	return nil
}
