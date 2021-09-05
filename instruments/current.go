package instruments

// Current represents instantaneous information
type Current interface {
	Last() float64
	Bid() float64
	Ask() float64
	Spread() float64
	OpenInterest() float64
}

// func (c Current) Spread() float64 {
// 	return c.Bid() - c.Ask()
// }
