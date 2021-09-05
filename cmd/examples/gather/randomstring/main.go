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
		i := rand.Intn(8) + 8
		fmt.Printf("Here is a random string of length %3d: %s\n", i, gather.RandomString(i))
	})
}
