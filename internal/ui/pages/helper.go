package pages

import (
	"fmt"
	"strings"

	"github.com/rthornton128/goncurses"
)

func prepareRow(text string, width int) string {
	if len(text) > 0 && text[0] == '+' {
		return text
	}

	textLen := len(text)
	spacesLen := width - 4 - textLen

	spaces := ""
	if spacesLen > 0 {
		spaces = strings.Repeat(" ", spacesLen)
	}

	if textLen > width-4 {
		text = text[:width-7] + "..."
	}

	return fmt.Sprintf("| %s%s |", text, spaces)
}

func getDelimiter(width int) string {
	return fmt.Sprintf("+%s+", strings.Repeat("-", width-2))
}

func cutBody(body []string, otherRowsLen int, height int, width int) []string {
	if len(body) <= height-otherRowsLen {
		return body
	}

	moreText := fmt.Sprintf(" %d more rows ", len(body)-height+otherRowsLen+1)
	body = body[:height-otherRowsLen-1]

	neededDotsLen := width - 4 - len(moreText)

	dotsBefore := strings.Repeat(".", neededDotsLen/2)
	dotsAfter := strings.Repeat(".", neededDotsLen-len(dotsBefore))

	return append(body, dotsBefore+moreText+dotsAfter)
}

func PrintColored(screen *goncurses.Window, color int16, y int, x int, text string) error {
	if err := screen.ColorOn(color); err != nil {
		return err
	}

	screen.MovePrint(y, x, text)

	if err := screen.ColorOff(color); err != nil {
		return err
	}

	return nil
}
