package main

import (
	"fmt"
	"strconv"

	"github.com/skeptycal/errorlogger"

	"github.com/skeptycal/littlejohn/examples"
	"github.com/skeptycal/littlejohn/logging"
)

// aliases for common functions
var (
	log    = errorlogger.Log
	Err    = errorlogger.Err
	Color  = logging.ColorMap
	GetEnv = logging.GetEnv
)

// Usage: "go run ./main.go [time delay]"
func main() {
	examples.RunExample(func() {
		log.Error("sample log error message for invalid integer conversion")
		_, err := strconv.Atoi("fake integer string")
		_ = Err(err)

		colorString, _ := GetEnv("CANARY")
		effectString, _ := GetEnv("REVERSED")

		fmt.Printf("%s%s%s\n", colorString, effectString, "Sample yellow message (using os environment ANSI strings CANARY and REVERSED...)")

		name := "FAKE_VAR"
		result, err := GetEnv(name)
		fmt.Printf("sample getEnv result and error: %q, %v\n", result, err)

		fmt.Printf("%ssample blue text.\n", Color["BLUE"])
	})
}
