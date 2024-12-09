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
)

const (
	actionReload   = iota
	actionSend     = iota
	actionNextRow  = iota
	actionPrevRow  = iota
	actionFirstRow = iota
	actionLastRow  = iota
	actionCopy     = iota
	actionQuit     = iota
)

func StartApp(client *jira.Client, config *resource.Config, window *goncurses.Window) error {
	var workLogsSent bool

	workLogs, err := import_data.ParseWorkLogs(config)
	if err != nil {
		return err
	}

	actions := action.NewActions(client, window)
	rowSelector := model.NewRowSelector(len(workLogs))

	t, err := actions.PrintWorkLogs.Print(workLogs, rowSelector)
	if err != nil {
		return err
	}

	for {
		switch waitForAction(window) {
		case actionReload:
			workLogs, err = import_data.ParseWorkLogs(config)
			if err != nil {
				return err
			}

			rowSelector = model.NewRowSelector(len(workLogs))

			t, err = actions.PrintWorkLogs.Print(workLogs, rowSelector)
			if err != nil {
				return err
			}
		case actionSend:
			if !workLogsSent {
				workLogsSent = true

				if err = actions.SendWorkLogs.Send(workLogs); err != nil {
					return err
				}
			}
		case actionNextRow:
			rowSelector.NextRow()
			actions.PrintWorkLogs.UpdateSelectedRow(t, rowSelector)
		case actionPrevRow:
			rowSelector.PrevRow()
			actions.PrintWorkLogs.UpdateSelectedRow(t, rowSelector)
		case actionFirstRow:
			rowSelector.FirstRow()
			actions.PrintWorkLogs.UpdateSelectedRow(t, rowSelector)
		case actionLastRow:
			rowSelector.LastRow()
			actions.PrintWorkLogs.UpdateSelectedRow(t, rowSelector)
		case actionCopy:
			if err = clipboard.CopyToClipboard(workLogs[rowSelector.Row-1].HeaderText); err != nil {
				return err
			}

			ui.EndWindow()
			return nil
		case actionQuit:
			ui.EndWindow()
			return nil
		}
	}
}
