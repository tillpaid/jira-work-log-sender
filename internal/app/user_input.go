package app

import "github.com/rthornton128/goncurses"

func waitForAction(window *goncurses.Window) int {
	var pressedKey, previousKey goncurses.Key

	for {
		previousKey = pressedKey
		pressedKey = window.GetChar()

		switch pressedKey {
		case 'r':
			return actionReload
		case 'l':
			if previousKey == 'l' {
				return actionSend
			}
		case 'j', goncurses.KEY_DOWN:
			return actionNextRow
		case 'k', goncurses.KEY_UP:
			return actionPrevRow
		case 'g':
			if previousKey == 'g' {
				return actionFirstRow
			}
		case 'G':
			return actionLastRow
		case 'y':
			return actionCopy
		case 'm':
			return actionToggleModifyTime
		case 'M':
			return actionToggleAllModifyTime
		case 'q', ' ', goncurses.KEY_ESC, goncurses.KEY_RETURN:
			return actionQuit
		}
	}
}
