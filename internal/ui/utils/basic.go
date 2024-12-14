package utils

import (
	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui"
)

func ColorOn(window *goncurses.Window, color int16) {
	if color != ui.DefaultColor {
		_ = window.ColorOn(color)
	}
}

func ColorOff(window *goncurses.Window, color int16) {
	if color != ui.DefaultColor {
		_ = window.ColorOff(color)
	}
}

func SelectedOn(window *goncurses.Window, isSelected bool) {
	if isSelected {
		_ = window.AttrOn(goncurses.A_REVERSE)
	}
}

func SelectedOff(window *goncurses.Window, isSelected bool) {
	if isSelected {
		_ = window.AttrOff(goncurses.A_REVERSE)
	}
}
