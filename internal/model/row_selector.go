package model

type RowSelector struct {
	Row       int
	rowsCount int
}

func NewRowSelector(rowsCount int) *RowSelector {
	return &RowSelector{Row: 1, rowsCount: rowsCount}
}

func (r *RowSelector) FirstRow() {
	r.Row = 1
}

func (r *RowSelector) LastRow() {
	r.Row = r.rowsCount
}

func (r *RowSelector) NextRow() {
	r.Row++

	if r.Row > r.rowsCount {
		r.Row = 1
	}
}

func (r *RowSelector) PrevRow() {
	r.Row--

	if r.Row < 1 {
		r.Row = r.rowsCount
	}
}
