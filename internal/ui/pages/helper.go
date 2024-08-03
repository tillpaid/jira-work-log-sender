package pages

import (
	"fmt"
	"strings"
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
