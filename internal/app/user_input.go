package app

import "github.com/rthornton128/goncurses"

func waitForAction(screen *goncurses.Window) int {
	for {
		switch screen.GetChar() {
		case 'r':
			return actionReload
		case 'l':
			return actionDump
		case 'q', ' ', goncurses.KEY_ESC, goncurses.KEY_RETURN:
			return actionQuit
		}
	}
}
