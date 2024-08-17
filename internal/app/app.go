package app

import (
	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/app/action"
	"github.com/tillpaid/paysera-log-time-golang/internal/import_data"
	"github.com/tillpaid/paysera-log-time-golang/internal/jira"
	"github.com/tillpaid/paysera-log-time-golang/internal/resource"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui"
)

const (
	actionReload = iota
	actionSend   = iota
	actionQuit   = iota
)

func StartApp(client *jira.Client, config *resource.Config, screen *goncurses.Window) error {
	workLogs, err := import_data.ParseWorkLogs(client, config)
	if err != nil {
		return err
	}

	if err := action.PrintWorkLogs(screen, workLogs); err != nil {
		return err
	}

	for {
		switch waitForAction(screen) {
		case actionReload:
			if err := action.PrintWorkLogs(screen, workLogs); err != nil {
				return err
			}
		case actionSend:
			if err := action.SendLogWorks(client, screen, workLogs); err != nil {
				return err
			}
		case actionQuit:
			ui.EndScreen()
			return nil
		}
	}
}
