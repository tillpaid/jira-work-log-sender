package table

import "github.com/rthornton128/goncurses"

const (
	BorderTypeHeaderTop = iota
	BorderTypeHeaderBottom
	BorderTypeTableBottom
	BorderCharFirst
	BorderCharMiddle
	BorderCharLast
)

var borderTypesMap = map[uint16]map[uint16]goncurses.Char{
	BorderTypeHeaderTop: {
		BorderCharFirst:  goncurses.ACS_ULCORNER,
		BorderCharMiddle: goncurses.ACS_TTEE,
		BorderCharLast:   goncurses.ACS_URCORNER,
	},
	BorderTypeHeaderBottom: {
		BorderCharFirst:  goncurses.ACS_LTEE,
		BorderCharMiddle: goncurses.ACS_PLUS,
		BorderCharLast:   goncurses.ACS_RTEE,
	},
	BorderTypeTableBottom: {
		BorderCharFirst:  goncurses.ACS_LLCORNER,
		BorderCharMiddle: goncurses.ACS_BTEE,
		BorderCharLast:   goncurses.ACS_LRCORNER,
	},
}

type BorderChars struct {
	Count int
	Char  goncurses.Char
}
