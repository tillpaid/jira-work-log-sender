package ui

import "github.com/rthornton128/goncurses"

const (
	DefaultColor = iota
	YellowOnBlack
	RedOnBlack
	GreenOnBlack
	CyanOnBlack
	MagentaOnBlack
)

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

func initColorPairs() error {
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
