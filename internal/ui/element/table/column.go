package table

import "github.com/tillpaid/paysera-log-time-golang/internal/service"

type Column struct {
	Text     string
	Width    int
	Position int
	Color    int16
}

func (c *Column) GetText() string {
	return service.GetTextWithSpaces(c.Text, c.Width)
}
