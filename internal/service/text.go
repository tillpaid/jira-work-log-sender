package service

import (
	"fmt"
	"strings"
)

func GetTextWithSpaces(text string, width int) string {
	width += 2

	text = strings.Replace(text, "\n", "|", -1)
	text = fmt.Sprintf(" %s ", text)

	neededSpaces := width - len(text)

	spaces := ""
	if neededSpaces > 0 {
		spaces = strings.Repeat(" ", neededSpaces)
	}

	text = text + spaces

	if len(text) > width {
		text = text[:width-1] + " "
	}

	return text
}
