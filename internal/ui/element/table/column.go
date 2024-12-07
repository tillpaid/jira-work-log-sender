package table

import "github.com/tillpaid/paysera-log-time-golang/internal/service"

type Column struct {
	Name     string
	Width    int
	Position int
}

func (c *Column) GetText() string {
	return service.GetTextWithSpaces(c.Name, c.Width)
}
