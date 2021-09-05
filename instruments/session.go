package instruments

// Session represents session information
type Session interface {
	Volume() float64
	Hi() float64
	Low() float64
}
