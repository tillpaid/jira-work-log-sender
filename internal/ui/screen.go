package ui

import "github.com/rthornton128/goncurses"

func InitializeWindow() (*goncurses.Window, error) {
	window, err := goncurses.Init()
	if err != nil {
		return nil, err
	}

	goncurses.Raw(true)
	goncurses.Echo(false)

	if err := initColors(); err != nil {
		return nil, err
	}

	if err := goncurses.Cursor(0); err != nil {
		return nil, err
	}

	if err := window.Keypad(true); err != nil {
		return nil, err
	}

	if err := window.Clear(); err != nil {
		return nil, err
	}

	return window, nil
}

func EndWindow() {
	goncurses.End()
}

func initColors() error {
	if !goncurses.HasColors() {
		return nil
	}

	if err := goncurses.StartColor(); err != nil {
		return err
	}

	if err := goncurses.UseDefaultColors(); err != nil {
		return err
	}

	if err := initColorPairs(); err != nil {
		return err
	}

	return nil
}
