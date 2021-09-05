package logging

import (
	"github.com/skeptycal/errorlogger"
)

var Err = errorlogger.Err
var Log = errorlogger.Log

func init() {
	// TODO initialize logger
	// check for settings in environment variables
	// or settings file.
}
