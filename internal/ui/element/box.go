package element

import "github.com/rthornton128/goncurses"

func DrawBox(window *goncurses.Window, height, width int, title string) {
	startY, startX := 0, 0

	window.MoveAddChar(startY, startX, goncurses.ACS_ULCORNER)
	for i := 1; i < width-1; i++ {
		window.MoveAddChar(startY, startX+i, goncurses.ACS_HLINE)
	}
	window.MoveAddChar(startY, startX+width-1, goncurses.ACS_URCORNER)

	for i := 1; i < height-1; i++ {
		window.MoveAddChar(startY+i, startX, goncurses.ACS_VLINE)
		window.MoveAddChar(startY+i, startX+width-1, goncurses.ACS_VLINE)
	}

	window.MoveAddChar(startY+height-1, startX, goncurses.ACS_LLCORNER)
	for i := 1; i < width-1; i++ {
		window.MoveAddChar(startY+height-1, startX+i, goncurses.ACS_HLINE)
	}
	window.MoveAddChar(startY+height-1, startX+width-1, goncurses.ACS_LRCORNER)

	if title != "" {
		window.MovePrint(startY, startX+2, " "+title+" ")
	}
}
