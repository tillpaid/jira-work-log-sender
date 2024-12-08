package model

type RowSelector struct {
	Row         int
	PreviousRow int
	rowsCount   int
}

func NewRowSelector(rowsCount int) *RowSelector {
	return &RowSelector{
		Row:         1,
		PreviousRow: 0,
		rowsCount:   rowsCount,
	}
}

func (r *RowSelector) FirstRow() {
	r.PreviousRow = r.Row
	r.Row = 1
}

func (r *RowSelector) LastRow() {
	r.PreviousRow = r.Row
	r.Row = r.rowsCount
}

func (r *RowSelector) NextRow() {
	r.PreviousRow = r.Row
	r.Row++

	if r.Row > r.rowsCount {
		r.Row = 1
	}
}

func (r *RowSelector) PrevRow() {
	r.PreviousRow = r.Row
	r.Row--

	if r.Row < 1 {
		r.Row = r.rowsCount
	}
}
