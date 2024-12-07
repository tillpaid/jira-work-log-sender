package table

type Header struct {
	Row *Row
}

func NewHeader(columns []*Column, position int) *Header {
	return &Header{
		Row: NewRow(columns, false, position),
	}
}