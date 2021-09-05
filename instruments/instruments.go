package instruments

type Underlying struct{}

func (u *Underlying) Data() DataSet {
	return nil
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
}

// UnTradeable implements an Instrument interface for
// an untradeable underlying that is used for reference
// or modeling
type UnTradeable interface {
	Instrument
}
