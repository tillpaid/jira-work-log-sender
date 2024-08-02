package app

import (
	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/app/action"
	"github.com/tillpaid/paysera-log-time-golang/internal/resource"
)

const (
	actionQuit   = iota
	actionReload = iota
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
		case actionQuit:
			return nil
		}
	}
}
