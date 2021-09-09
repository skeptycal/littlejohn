package instruments

// tradeable implements an Instrument interface for
// a tradeable underlying that can also be used for
// reference or modeling
type Tradeable struct {
	Underlying
}

func (t *Tradeable) IsTradeable() bool {
	return true
}

// untradeable implements an Instrument interface for
// an untradeable underlying that is used for reference
// or modeling
type Untradeable struct {
	Underlying
}

func (t *Untradeable) IsTradeable() bool {
	return false
}

type Data interface {
	Current
	Session
	Historical
}

type DataSet interface{}

// Instrument creates acces to the underlying instrument
type Instrument interface {
	Data() DataSet
	String() string
	IsTradeable() bool
}

type Underlying struct {
	symbol string
}

func (u *Underlying) Symbol() string {
	return u.symbol
}

func (u *Underlying) Data() DataSet {
	return nil
}
