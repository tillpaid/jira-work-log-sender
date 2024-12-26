package table

import (
	"github.com/rthornton128/goncurses"
	"github.com/tillpaid/jira-work-log-sender/internal/ui/utils"
)

type Table struct {
	Header *Header
	Rows   []*Row
	window *goncurses.Window
}

func NewTable(header *Header, rows []*Row, window *goncurses.Window) *Table {
	return &Table{Header: header, Rows: rows, window: window}
}

func (t *Table) GetBorderChars(borderType uint16) []BorderChars {
	chars := []BorderChars{
		{1, borderTypesMap[borderType][BorderCharFirst]},
	}

	for _, column := range t.Header.Row.Columns {
		chars = append(chars, BorderChars{column.Width, goncurses.ACS_HLINE})
		chars = append(chars, BorderChars{1, borderTypesMap[borderType][BorderCharMiddle]})
	}

	chars[len(chars)-1] = BorderChars{1, borderTypesMap[borderType][BorderCharLast]}

	return chars
}

func (t *Table) ShowRow(i int) {
	t.Rows[i].ShowText = true
	t.ReDrawRow(t.Rows[i])
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
	defer t.window.MoveAddChar(row.Number, row.Columns[0].Position-1, goncurses.ACS_VLINE)
	utils.SelectedOn(t.window, row.IsSelected)

	for _, column := range row.Columns {
		t.printColumnText(row, column)
	}

	utils.SelectedOff(t.window, row.IsSelected)
	t.window.MoveAddChar(row.Number, row.CalculateLastPosition(), goncurses.ACS_VLINE)
}

func (t *Table) printColumnText(row *Row, column *Column) {
	t.window.MoveAddChar(row.Number, column.Position-1, goncurses.ACS_VLINE)
	utils.ColorOn(t.window, column.ResolveColor(row))
	t.window.MovePrint(row.Number, column.Position, column.GetText(row.ShowText))
	utils.ColorOff(t.window, column.ResolveColor(row))
}

func (t *Table) printBorderChars(y int, x int, borderChars []BorderChars) {
	t.window.Move(y, x)

	for _, borderChar := range borderChars {
		for i := 0; i < borderChar.Count; i++ {
			t.window.AddChar(borderChar.Char)
		}
	}
}
