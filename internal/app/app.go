package app

import (
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
	config    *resource.Config

	table        *table.Table
	workLogs     []model.WorkLog
	selector     *model.RowSelector
	workLogsSent bool
}

func NewApplication(window *goncurses.Window, client *jira.Client, input *UserInput, actions *action.Actions, config *resource.Config) *Application {
	selector := model.NewRowSelector(0)

	return &Application{
		window:       window,
		client:       client,
		userInput:    input,
		config:       config,
		actions:      actions,
		selector:     selector,
		workLogsSent: false,
	}
}

func (a *Application) Start() error {
	if err := a.loadWorkLogs(); err != nil {
		return err
	}

	if err := a.printTable(); err != nil {
		return err
	}

	handleResize(&a.window, &a.table, a.selector, a.actions, &a.workLogs)

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
			a.actions.PrintWorkLogs.UpdateSelectedRow(a.table, a.selector)
		case actionPrevRow:
			a.selector.PrevRow()
			a.actions.PrintWorkLogs.UpdateSelectedRow(a.table, a.selector)
		case actionFirstRow:
			a.selector.FirstRow()
			a.actions.PrintWorkLogs.UpdateSelectedRow(a.table, a.selector)
		case actionLastRow:
			a.selector.LastRow()
			a.actions.PrintWorkLogs.UpdateSelectedRow(a.table, a.selector)
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
	if err := a.loadWorkLogs(); err != nil {
		return err
	}

	return a.printTable()
}

func (a *Application) processActionSend() error {
	if a.workLogsSent {
		return nil
	}

	a.workLogsSent = true
	return a.actions.SendWorkLogs.Send(a.workLogs)
}

func (a *Application) processActionCopy(exit bool) error {
	if len(a.workLogs) == 0 {
		return nil
	}

	if err := clipboard.CopyToClipboard(a.workLogs[a.selector.Row-1].HeaderText); err != nil {
		return err
	}

	if exit {
		ui.EndWindow()
	}

	return nil
}

func (a *Application) processActionToggleModifyTime(all bool) error {
	if len(a.workLogs) == 0 || !a.config.TimeAdjustment.Enabled {
		return nil
	}

	for i := range a.workLogs {
		if all || i == a.selector.Row-1 {
			a.workLogs[i].ToggleModifyTime()
		}
	}

	a.workLogs = service.ModifyWorkLogsTime(a.workLogs, a.config)
	return a.printTable()
}

func (a *Application) loadWorkLogs() error {
	workLogs, err := import_data.ParseWorkLogs(a.config, a.workLogs)
	if err != nil {
		return err
	}

	a.workLogs = workLogs
	a.selector.Update(len(workLogs))

	return nil
}

func (a *Application) printTable() error {
	t, err := a.actions.PrintWorkLogs.Print(a.workLogs, a.selector)
	if err != nil {
		return err
	}

	a.table = t
	return nil
}
