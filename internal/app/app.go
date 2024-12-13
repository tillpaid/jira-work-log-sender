package app

import (
	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/app/action"
	"github.com/tillpaid/paysera-log-time-golang/internal/clipboard"
	"github.com/tillpaid/paysera-log-time-golang/internal/import_data"
	"github.com/tillpaid/paysera-log-time-golang/internal/jira"
	"github.com/tillpaid/paysera-log-time-golang/internal/model"
	"github.com/tillpaid/paysera-log-time-golang/internal/resource"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui/element/table"
)

const (
	actionReload = iota
	actionSend
	actionNextRow
	actionPrevRow
	actionFirstRow
	actionLastRow
	actionCopy
	actionQuit
)

type Application struct {
	window  *goncurses.Window
	client  *jira.Client
	actions *action.Actions
	config  *resource.Config

	table        *table.Table
	workLogs     []model.WorkLog
	selector     *model.RowSelector
	workLogsSent bool
}

func NewApplication(window *goncurses.Window, client *jira.Client, actions *action.Actions, config *resource.Config) *Application {
	selector := model.NewRowSelector(0)

	return &Application{
		window:       window,
		client:       client,
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
		switch waitForAction(a.window) {
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
			return a.processActionCopy()
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

	if err := a.printTable(); err != nil {
		return err
	}

	return nil
}

func (a *Application) processActionSend() error {
	if a.workLogsSent {
		return nil
	}

	a.workLogsSent = true
	return a.actions.SendWorkLogs.Send(a.workLogs)
}

func (a *Application) processActionCopy() error {
	if len(a.workLogs) == 0 {
		return nil
	}

	if err := clipboard.CopyToClipboard(a.workLogs[a.selector.Row-1].HeaderText); err != nil {
		return err
	}

	ui.EndWindow()
	return nil
}

func (a *Application) loadWorkLogs() error {
	workLogs, err := import_data.ParseWorkLogs(a.config)
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
