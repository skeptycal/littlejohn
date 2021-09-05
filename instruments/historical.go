package instruments

// Historical represents historical information
type Historical interface {
	Hi52() float64
	Low52() float64
}
