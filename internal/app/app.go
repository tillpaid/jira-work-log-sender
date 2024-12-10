package app

import (
	"os"
	"os/signal"
	"syscall"

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

	handleResize(&window, &t, rowSelector, actions, &workLogs)

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

func handleResize(window **goncurses.Window, t **table.Table, rowSelector *model.RowSelector, actions *action.Actions, workLogs *[]model.WorkLog) {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGWINCH)

	go func() {
		for {
			<-sigchan

			ui.EndWindow()

			newWindow, _ := ui.InitializeWindow()
			newWindow.Refresh()

			rowSelector.Reset()
			newTable, _ := actions.PrintWorkLogs.Print(*workLogs, rowSelector)

			discardResidualInput(newWindow)

			*window = newWindow
			*t = newTable
		}
	}()
}

func discardResidualInput(window *goncurses.Window) {
	window.Timeout(0)
	defer window.Timeout(-1)

	for {
		if key := window.GetChar(); key == 0 {
			break
		}
	}
}
