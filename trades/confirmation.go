package trades

import (
	"time"
)

func NewConfirmation() Confirmation {
	return &confirmation{}
}

type Confirmation interface{}

type confirmation struct {
	mainUnderlying Tradeable
	legs           []leg
	timestamp      time.Time
}

func (c *confirmation) Time() time.Time {
	return c.timestamp
}

func (c *confirmation) Symbol() string {
	return c.underlying.Symbol()
}

func (c *confirmation) Underlying() Tradeable {
	return c.underlying
}

type leg struct {
	timestamp  time.Time
	side       Side
	quantity   int
	underlying Tradeable
}

type Side int

const (
	Sell Side = iota - 1
	None
	Buy
)

func (s Side) String() string {
	if s == Buy {
		return "Buy"
	}
	if s == Sell {
		return "Sell"
	}
	return "None"
}
