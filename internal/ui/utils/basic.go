package utils

import (
	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui"
)

func ColorOn(screen *goncurses.Window, color int16) {
	if color != ui.DefaultColor {
		screen.ColorOn(color)
	}
}

func ColorOff(screen *goncurses.Window, color int16) {
	if color != ui.DefaultColor {
		screen.ColorOff(color)
	}
}

func SelectedOn(screen *goncurses.Window, isSelected bool) {
	if isSelected {
		screen.AttrOn(goncurses.A_REVERSE)
	}
}

func SelectedOff(screen *goncurses.Window, isSelected bool) {
	if isSelected {
		screen.AttrOff(goncurses.A_REVERSE)
	}
}
