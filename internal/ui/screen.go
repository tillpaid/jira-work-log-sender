package ui

import "github.com/rthornton128/goncurses"

const (
	YellowOnBlack = iota + 1
	RedOnBlack    = iota + 1
	GreenOnBlack  = iota + 1
)

func InitializeScreen() (*goncurses.Window, error) {
	screen, err := goncurses.Init()
	if err != nil {
		return nil, err
	}

	goncurses.Raw(true)
	goncurses.Echo(false)

	if goncurses.HasColors() {
		if err = goncurses.StartColor(); err != nil {
			return nil, err
		}

		if err = goncurses.InitPair(YellowOnBlack, goncurses.C_YELLOW, goncurses.C_BLACK); err != nil {
			return nil, err
		}

		if err = goncurses.InitPair(RedOnBlack, goncurses.C_RED, goncurses.C_BLACK); err != nil {
			return nil, err
		}

		if err = goncurses.InitPair(GreenOnBlack, goncurses.C_GREEN, goncurses.C_BLACK); err != nil {
			return nil, err
		}
	}

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
