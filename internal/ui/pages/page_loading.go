package pages

import (
	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui"
)

type lastPrinted struct {
	row  int
	x    int
	text string
}

type Loading struct {
	screen     *goncurses.Window
	currentRow int
	lastPrint  lastPrinted
}

func NewLoading(screen *goncurses.Window) *Loading {
	return &Loading{
		screen:     screen,
		currentRow: 1,
		lastPrint:  lastPrinted{},
	}
}

func (l *Loading) PrintRow(text string, x int) {
	l.screen.MovePrint(l.currentRow, x+2, text)
	l.screen.Refresh()

	_ = PrintColored(l.screen, ui.GreenOnBlack, l.lastPrint.row, l.lastPrint.x+2, l.lastPrint.text)

	l.lastPrint.row = l.currentRow
	l.lastPrint.x = x
	l.lastPrint.text = text

	l.currentRow++
}

func (l *Loading) PrintBorder() {
	height, width := l.screen.MaxYX()
	delimiter := getDelimiter(width)

	l.screen.MovePrint(0, 0, delimiter)
	l.screen.MovePrint(height-1, 0, delimiter)

	for i := 1; i < height-1; i++ {
		l.screen.MovePrint(i, 0, "|")
		l.screen.MovePrint(i, width-1, "|")
	}
}
