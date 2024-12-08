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

	if len(c.Text) > limit {
		return c.Text[:limit-1] + " "
	}

	text := " " + strings.Replace(c.Text, "\n", "|", -1)

	neededSpaces := limit - len(text)
	if neededSpaces > 0 {
		text += strings.Repeat(" ", neededSpaces)
	}

	return text
}
