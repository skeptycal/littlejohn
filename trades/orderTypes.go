package trades

type OrderType uint8

const (
	FakeOrderType OrderType = iota
	Buy
	Sell
	CreditSpread
	CreditIronCondor
	DebitSpread
	DebitCalendar
	Straddle
	Strangled
	Butterfly
	ButterflyBrokenWing

	Reserved OrderType = 1<<8 - 1 // max order type
)
