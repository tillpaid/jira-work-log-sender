package table

import (
	"strings"
)

type Column struct {
	Text     string
	Width    int
	Position int
	Color    int16
}

func (c *Column) GetText(showText bool) string {
	limit := c.Width + 2

	if !showText {
		return strings.Repeat(" ", limit)
	}

	text := " " + strings.ReplaceAll(c.Text, "\n", "|") + " "
	if len(text) > limit {
		return text[:limit-1] + " "
	}

	neededSpaces := limit - len(text)
	if neededSpaces > 0 {
		text += strings.Repeat(" ", neededSpaces)
	}

	return text
}
