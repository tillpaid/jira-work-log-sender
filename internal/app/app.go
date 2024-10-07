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
	"github.com/tillpaid/paysera-log-time-golang/internal/ui/pages"
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

func StartApp(client *jira.Client, config *resource.Config, screen *goncurses.Window, loading *pages.Loading) error {
	var workLogsSent bool

	workLogs, err := import_data.ParseWorkLogs(loading, client, config)
	if err != nil {
		return err
	}

	actions := action.NewActions(client, screen)
	rowSelector := model.NewRowSelector(len(workLogs))

	if err = actions.PrintWorkLogs.Print(workLogs, rowSelector.Row); err != nil {
		return err
	}

	for {
		switch waitForAction(screen) {
		case actionReload:
			if err = actions.PrintWorkLogs.Print(workLogs, rowSelector.Row); err != nil {
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

			if err = actions.PrintWorkLogs.Print(workLogs, rowSelector.Row); err != nil {
				return err
			}
		case actionPrevRow:
			rowSelector.PrevRow()

			if err = actions.PrintWorkLogs.Print(workLogs, rowSelector.Row); err != nil {
				return err
			}
		case actionFirstRow:
			rowSelector.FirstRow()

			if err = actions.PrintWorkLogs.Print(workLogs, rowSelector.Row); err != nil {
				return err
			}
		case actionLastRow:
			rowSelector.LastRow()

			if err = actions.PrintWorkLogs.Print(workLogs, rowSelector.Row); err != nil {
				return err
			}
		case actionCopy:
			if err = clipboard.CopyToClipboard(workLogs[rowSelector.Row-1].HeaderText); err != nil {
				return err
			}

			ui.EndScreen()
			return nil
		case actionQuit:
			ui.EndScreen()
			return nil
		}
	}
}
