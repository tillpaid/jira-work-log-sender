package ui

import "github.com/rthornton128/goncurses"

func InitializeWindow() (*goncurses.Window, error) {
	window, err := goncurses.Init()
	if err != nil {
		return nil, err
	}

	if err := configureParams(window); err != nil {
		return nil, err
	}

	return window, nil
}

func EndWindow() {
	goncurses.End()
}

func configureParams(window *goncurses.Window) error {
	goncurses.Raw(true)
	goncurses.Echo(false)

	if err := initColors(); err != nil {
		return err
	}

	if err := goncurses.Cursor(0); err != nil {
		return err
	}

	if err := window.Keypad(true); err != nil {
		return err
	}

	if err := window.Clear(); err != nil {
		return err
	}

	return nil
}
