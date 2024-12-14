package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/app/action"
	"github.com/tillpaid/paysera-log-time-golang/internal/model"
	"github.com/tillpaid/paysera-log-time-golang/internal/service"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui/element/table"
)

func handleResize(window **goncurses.Window, t **table.Table, rowSelector *model.RowSelector, actions *action.Actions, workLogs *[]model.WorkLog) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGWINCH)

	go func() {
		defer service.HandlePanic()

		for range c {
			ui.EndWindow()

			newWindow, _ := ui.InitializeWindow()
			newWindow.Refresh()

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

	for window.GetChar() != 0 {
	}
}
