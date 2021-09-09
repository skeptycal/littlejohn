package trades

type Trades map[int64]confirmation

type Position struct {
	underlying Tradeable
	trades     []confirmation
}

type Positioner interface{}
