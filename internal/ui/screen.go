package ui

import "github.com/rthornton128/goncurses"

func InitializeScreen() (*goncurses.Window, error) {
	screen, err := goncurses.Init()
	if err != nil {
		return nil, err
	}

	goncurses.Raw(true)
	goncurses.Echo(false)

	if err := goncurses.Cursor(0); err != nil {
		return nil, err
	}

	if err := screen.Clear(); err != nil {
		return nil, err
	}

	return screen, nil
}

func EndScreen() {
	goncurses.End()
}
