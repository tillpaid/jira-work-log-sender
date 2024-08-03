package app

import (
	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/app/action"
	"github.com/tillpaid/paysera-log-time-golang/internal/resource"
)

const (
	actionReload = iota
	actionDump   = iota
	actionQuit   = iota
)

func StartApp(config *resource.Config, screen *goncurses.Window) error {
	if err := action.PrintWorkLogs(config, screen); err != nil {
		return err
	}

	for {
		switch waitForAction(screen) {
		case actionReload:
			if err := action.PrintWorkLogs(config, screen); err != nil {
				return err
			}
		case actionDump:
			if err := action.DumpWorkLogs(config, screen); err != nil {
				return err
			}
		case actionQuit:
			return nil
		}
	}
}
