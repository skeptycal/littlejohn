package examples

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

// DefaultDelay is the default sleep delay between successive
// example output statements sent to Stdout.
//
// It is purposely not setup as a constant so the default value
// of one second can be modified as needed.
var DefaultDelay time.Duration = 1 * time.Second

func PrintArch() {
	fmt.Println("compiler and architecture details: ", runtime.Compiler, runtime.GOARCH, runtime.GOOS)
}

// RunExample will loop through an example function fn and
// display the results to Stdout. There is a delay between
// each function call (default 1 second) that may be adjusted
// by setting DefaultDelay before calling RunExample or by
// passing a CLI argument that contains a valid time.Duration
// string.
func RunExample(fn func()) {
	go Looper(fn, CliDelay())
	PrintArch()
	fmt.Println("Press the <ENTER> key to stop anytime...")
	fmt.Scanln()
}

// Looper repeats a function call with a sleep delay between
// calls as defined by the time.Duration 'pause'.
// Any duration less than 10 ms is adjusted up to 10 ms.
// The function fn cannot have arguments or return values.
func Looper(fn func(), pause time.Duration) {
	if pause < 10*time.Millisecond {
		pause = time.Second * 1
	}

	for {
		go fn()
		time.Sleep(pause)
	}
}

// Args concatenates any user provided command line arguments.
// If no arguments are provided, the empty string is returned.
// The delimiter string delim is placed between elements in the
// resulting string.
func Args(delim string) string {
	if len(os.Args) == 1 {
		return ""
	}
	return strings.Join(os.Args[1:], delim)
}

// CliDelay parses the command line arguments for a correctly formatted
// time duration string to use for the sleep delay between output statements.
// If no command line arguments are provided or if no valid time.Duration
// is detected, the DefaultDelay is returned.
//
// This is intended for use with example code and does not parse
// the command line arguments in any other way
//
// e.g. If the command line input was:
//  go run ./main.go 50 ms
// CliDelay would return a time.Duration of 50 milliseconds
func CliDelay() time.Duration {

	if len(os.Args) > 1 {
		// not using 'space' delim: time.ParseDuration does not
		// seem to work properly with spaces between numbers and units
		// e.g. 5 s does not seem to work, but 5s does
		d, err := time.ParseDuration(Args(""))
		if err == nil {
			return d
		}
	}
	return DefaultDelay
}
