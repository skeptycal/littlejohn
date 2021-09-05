package trades

type refType uint8

const (
	Last refType = iota
	Atm
	Itm1
	Itm2
	Otm1
	Otm2
)
