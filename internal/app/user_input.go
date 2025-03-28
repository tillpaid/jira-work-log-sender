package app

import "github.com/rthornton128/goncurses"

type UserInput struct {
	window        *goncurses.Window
	singleKeyMap  map[goncurses.Key]int
	doubleKeysMap map[goncurses.Key]int
}

func NewUserInput(window *goncurses.Window) *UserInput {
	singleKeyConfig := map[int][]goncurses.Key{
		actionReload:              {'r'},
		actionLastRow:             {'G'},
		actionCopy:                {'y'},
		actionCopyWithoutExit:     {'Y'},
		actionToggleModifyTime:    {'m'},
		actionToggleAllModifyTime: {'M'},
		actionNextRow:             {'j', goncurses.KEY_DOWN},
		actionPrevRow:             {'k', goncurses.KEY_UP},
		actionQuit:                {'q', ' ', goncurses.KEY_ESC, goncurses.KEY_RETURN},
	}
	doubleKeysConfig := map[int][]goncurses.Key{
		actionSend:     {'l'},
		actionFirstRow: {'g'},
	}

	return &UserInput{
		window:        window,
		singleKeyMap:  convertConfigToMap(singleKeyConfig),
		doubleKeysMap: convertConfigToMap(doubleKeysConfig),
	}
}

func (u *UserInput) WaitForAction() int {
	var pressedKey, previousKey goncurses.Key

	for {
		previousKey = pressedKey
		pressedKey = u.window.GetChar()

		if action, ok := u.singleKeyMap[pressedKey]; ok {
			return action
		}

		if action, ok := u.doubleKeysMap[pressedKey]; ok {
			if pressedKey == previousKey {
				return action
			}
		}
	}
}

func convertConfigToMap(config map[int][]goncurses.Key) map[goncurses.Key]int {
	keyMap := make(map[goncurses.Key]int)

	for action, keys := range config {
		for _, key := range keys {
			keyMap[key] = action
		}
	}

	return keyMap
}
