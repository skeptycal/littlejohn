package trades

import (
	"fmt"

	"github.com/skeptycal/littlejohn/instruments"
)

type Order interface {
	Check() error
	Execute() (c Confirmation, err error)
}

// Tradeable implements an Instrument interface for
// a tradeable underlying
type Tradeable interface {
	instruments.Instrument
	isTradeable() bool
}

type tradeable struct {
	*instruments.Underlying
}

func (t *tradeable) String() string {
	return fmt.Sprintf("%v", *t)
}

func (t *tradeable) isTradeable() bool {
	return true
}

func NewTradeable(u *instruments.Underlying) Tradeable {
	return &tradeable{u}
}

func NewOrder() Order {
	return &order{}
}

type order struct {
	id         int64
	Underlying Tradeable
	actions    []OrderType
}

func (o *order) Check() error {
	//TODO implement checks for confirmation, authentication, permissions, scopes, etc.
	return nil
}

func (o *order) Execute() (c Confirmation, err error) {
	err = o.Check()
	if err != nil {
		return nil, err
	}

	//TODO create order execute "paper trading" interface

	return NewConfirmation(), nil

}

type leg struct {
	id         int64
	Underlying Tradeable
	oType      OrderType
}

type OrderTypeDetails struct {
	reference refType
	legs      []leg
}
