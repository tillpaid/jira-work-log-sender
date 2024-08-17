package app

import "github.com/rthornton128/goncurses"

func waitForAction(screen *goncurses.Window) int {
	var pressedKey, previousKey goncurses.Key

	for {
		previousKey = pressedKey
		pressedKey = screen.GetChar()

		switch pressedKey {
		case 'r':
			return actionReload
		case 'l':
			if previousKey == 'l' {
				return actionSend
			}
		case 'q', ' ', goncurses.KEY_ESC, goncurses.KEY_RETURN:
			return actionQuit
		}
	}
}
