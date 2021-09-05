package trades

import (
	"time"
)

func NewConfirmation() Confirmation {
	return &confirmation{}
}

type Confirmation interface{}

type confirmation struct {
	u Tradeable
	t time.Now()
}
