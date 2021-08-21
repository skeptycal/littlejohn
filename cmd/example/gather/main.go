package main

import (
	"fmt"

	"github.com/skeptycal/littlejohn/gather"
)

func main() {
	fmt.Println("Here is a random string: ", gather.RandomString(8))
}
