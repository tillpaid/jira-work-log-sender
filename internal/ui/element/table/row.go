package table

type Row struct {
	Columns    []*Column
	Number     int
	IsSelected bool
	ShowText   bool
}

func NewRow(columns []*Column, number int, isSelected bool, position int, showText bool) *Row {
	row := &Row{
		Columns:    []*Column{},
		Number:     number,
		IsSelected: isSelected,
		ShowText:   showText,
	}

	for _, column := range columns {
		column.Position = position
		row.Columns = append(row.Columns, column)
		position += column.Width + 3
	}

	return row
}

func (r *Row) CalculateLastPosition() int {
	return r.Columns[len(r.Columns)-1].Position + r.Columns[len(r.Columns)-1].Width + 2
}
