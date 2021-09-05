package main

import (
	"fmt"
	"math/rand"

	"github.com/skeptycal/littlejohn/examples"
	"github.com/skeptycal/littlejohn/gather"
)

// Usage: "go run ./main.go [time delay]"
func main() {
	examples.RunExample(func() {
		i := rand.Intn(8) + 1
		j := rand.Intn(32) + 1 + i
		fmt.Printf("Here is a random string of between %3d and %3d: %d\n", i, j, gather.RandInt(i, j))
	})
}
