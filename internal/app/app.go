package app

import (
	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/app/action"
	"github.com/tillpaid/paysera-log-time-golang/internal/jira"
	"github.com/tillpaid/paysera-log-time-golang/internal/resource"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui"
)

const (
	actionReload = iota
	actionDump   = iota
	actionQuit   = iota
)

func StartApp(client *jira.Client, config *resource.Config, screen *goncurses.Window) error {
	if err := action.PrintWorkLogs(client, config, screen); err != nil {
		return err
	}

	for {
		switch waitForAction(screen) {
		case actionReload:
			if err := action.PrintWorkLogs(client, config, screen); err != nil {
				return err
			}
		case actionDump:
			ui.EndScreen()
			return action.DumpWorkLogs(client, config)
		case actionQuit:
			ui.EndScreen()
			return nil
		}
	}
}
