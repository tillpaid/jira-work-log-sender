package table

import (
	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui/utils"
)

type Table struct {
	Header *Header
	Rows   []*Row
	screen *goncurses.Window
}

func NewTable(header *Header, rows []*Row, screen *goncurses.Window) *Table {
	return &Table{Header: header, Rows: rows, screen: screen}
}

func (t *Table) GetBorderChars(borderType uint16) []TableBorderChars {
	chars := []TableBorderChars{
		{1, borderTypesMap[borderType][BorderCharFirst]},
	}

	for _, column := range t.Header.Row.Columns {
		chars = append(chars, TableBorderChars{column.Width + 2, goncurses.ACS_HLINE})
		chars = append(chars, TableBorderChars{1, borderTypesMap[borderType][BorderCharMiddle]})
	}

	chars[len(chars)-1] = TableBorderChars{1, borderTypesMap[borderType][BorderCharLast]}

	return chars
}

func (t *Table) Draw() {
	t.drawHeader()
	t.drawRows()
}

func (t *Table) ReDrawRow(row *Row) {
	t.printRowText(row)
}

func (t *Table) drawHeader() {
	baseY := 1

	t.printBorderChars(baseY, 2, t.GetBorderChars(BorderTypeHeaderTop))
	t.printRowText(t.Header.Row)
	t.printBorderChars(baseY+2, 2, t.GetBorderChars(BorderTypeHeaderBottom))
	t.printBorderChars(baseY+3+len(t.Rows), 2, t.GetBorderChars(BorderTypeTableBottom))
}

func (t *Table) drawRows() {
	for _, row := range t.Rows {
		t.printRowText(row)
	}
}

func (t *Table) printRowText(row *Row) {
	defer t.screen.MoveAddChar(row.Number, row.Columns[0].Position-1, goncurses.ACS_VLINE)
	utils.SelectedOn(t.screen, row.IsSelected)

	for _, column := range row.Columns {
		t.printColumnText(row, column)
	}

	utils.SelectedOff(t.screen, row.IsSelected)
	t.screen.MoveAddChar(row.Number, row.CalculateLastPosition(), goncurses.ACS_VLINE)
}

func (t *Table) printColumnText(row *Row, column *Column) {
	t.screen.MoveAddChar(row.Number, column.Position-1, goncurses.ACS_VLINE)
	utils.ColorOn(t.screen, column.Color)
	t.screen.MovePrint(row.Number, column.Position, column.GetText(row.ShowText))
	utils.ColorOff(t.screen, column.Color)
}

func (t *Table) printBorderChars(y int, x int, borderChars []TableBorderChars) {
	t.screen.Move(y, x)

	for _, borderChar := range borderChars {
		for i := 0; i < borderChar.Count; i++ {
			t.screen.AddChar(borderChar.Char)
		}
	}
}
