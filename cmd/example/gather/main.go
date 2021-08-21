package main

import (
	"fmt"
	"math/rand"

	"github.com/skeptycal/littlejohn/gather"
)

func main() {
	go gather.Looper(example)
	fmt.Println("Press the Enter Key to stop anytime")
	fmt.Scanln()
}

func example() {
	i := rand.Intn(8) + 8
	fmt.Printf("Here is a random string of length %3d: %s\n", i, gather.RandomString(i))
}
