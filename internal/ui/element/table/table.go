package table

import (
	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/paysera-log-time-golang/internal/ui"
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
	if len(row.Columns) > 0 {
		t.screen.MoveAddChar(row.Number, row.Columns[0].Position-1, goncurses.ACS_VLINE)
	}

	if row.IsSelected {
		t.screen.AttrOn(goncurses.A_REVERSE)
	}

	for i, column := range row.Columns {
		if i > 0 {
			t.screen.MoveAddChar(row.Number, column.Position-1, goncurses.ACS_VLINE)
		}

		if column.Color != 0 && column.Color != ui.DefaultColor {
			t.screen.ColorOn(column.Color)
		}

		t.screen.MovePrint(row.Number, column.Position, column.GetText(row.ShowText))

		if column.Color != 0 && column.Color != ui.DefaultColor {
			t.screen.ColorOff(column.Color)
		}
	}

	if row.IsSelected {
		t.screen.AttrOff(goncurses.A_REVERSE)
	}

	t.screen.MoveAddChar(row.Number, row.CalculateLastPosition(), goncurses.ACS_VLINE)
}

func (t *Table) printBorderChars(y int, x int, borderChars []TableBorderChars) {
	t.screen.Move(y, x)

	for _, borderChar := range borderChars {
		for i := 0; i < borderChar.Count; i++ {
			t.screen.AddChar(borderChar.Char)
		}
	}
}
