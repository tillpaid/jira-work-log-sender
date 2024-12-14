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

	colorPairs := map[int16][2]int16{
		RedOnBlack:     {goncurses.C_RED, -1},
		YellowOnBlack:  {goncurses.C_YELLOW, -1},
		GreenOnBlack:   {goncurses.C_GREEN, -1},
		CyanOnBlack:    {goncurses.C_CYAN, -1},
		MagentaOnBlack: {goncurses.C_MAGENTA, -1},
	}

	for pairID, colors := range colorPairs {
		if err := goncurses.InitPair(pairID, colors[0], colors[1]); err != nil {
			return err
		}
	}

	return nil
}
