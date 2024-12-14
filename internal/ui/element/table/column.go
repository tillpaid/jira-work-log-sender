package table

import (
	"strings"

	"github.com/tillpaid/paysera-log-time-golang/internal/ui"
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

func (c *Column) ResolveColor(row *Row) int16 {
	if c.Color != ui.DefaultColor {
		return c.Color
	}

	return row.Color
}
