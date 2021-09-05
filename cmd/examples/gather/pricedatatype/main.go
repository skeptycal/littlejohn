package main

import (
	"fmt"

	"github.com/skeptycal/littlejohn/examples"
	"github.com/skeptycal/littlejohn/gather"
)

// Usage: "go run ./main.go [time delay]"
func main() {
	examples.RunExample(func() {
		a := gather.RandInt(111, 99999)
		var p gather.Price = uint32(a)
		fmt.Printf("Here is an example of price data (type %T): %v (%#v)\n", p, p, p)
	})
}
