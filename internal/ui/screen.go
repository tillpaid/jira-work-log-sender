package ui

import "github.com/rthornton128/goncurses"

func InitializeScreen() (*goncurses.Window, error) {
	screen, err := goncurses.Init()
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

	if err := screen.Keypad(true); err != nil {
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

	colorPairs := map[int16][2]int16{
		RedOnBlack:     {goncurses.C_RED, DefaultColor},
		GreenOnBlack:   {goncurses.C_GREEN, DefaultColor},
		CyanOnBlack:    {goncurses.C_CYAN, DefaultColor},
		MagentaOnBlack: {goncurses.C_MAGENTA, DefaultColor},
	}

	for pairID, colors := range colorPairs {
		if err := goncurses.InitPair(pairID, colors[0], colors[1]); err != nil {
			return err
		}
	}

	return nil
}
