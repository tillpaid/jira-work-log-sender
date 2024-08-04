package app

import "github.com/rthornton128/goncurses"

func waitForAction(screen *goncurses.Window) int {
	var previousKey goncurses.Key

	for {
		pressedKey := screen.GetChar()

		switch pressedKey {
		case 'r':
			return actionReload
		case 'l':
			if previousKey == 'l' {
				return actionDump
			}
		case 'q', ' ', goncurses.KEY_ESC, goncurses.KEY_RETURN:
			return actionQuit
		}

		previousKey = pressedKey
	}
}
