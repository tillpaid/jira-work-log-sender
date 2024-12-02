package ui

import "github.com/rthornton128/goncurses"

const (
	YellowOnBlack = iota + 1
	RedOnBlack    = iota + 1
	GreenOnBlack  = iota + 1
	CyanOnBlack   = iota + 1
	DefaultColor  = -1
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

		if err = goncurses.UseDefaultColors(); err != nil {
			return nil, err
		}

		if err = goncurses.InitPair(YellowOnBlack, goncurses.C_YELLOW, DefaultColor); err != nil {
			return nil, err
		}

		if err = goncurses.InitPair(RedOnBlack, goncurses.C_RED, DefaultColor); err != nil {
			return nil, err
		}

		if err = goncurses.InitPair(GreenOnBlack, goncurses.C_GREEN, DefaultColor); err != nil {
			return nil, err
		}

		if err = goncurses.InitPair(CyanOnBlack, goncurses.C_CYAN, DefaultColor); err != nil {
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
