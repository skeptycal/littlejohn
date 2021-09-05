package config

import (
	"os"
	"strings"

	"github.com/skeptycal/errorlogger"
)

var (
	Log                    = errorlogger.Log
	Err                    = errorlogger.Err
	ColorMap AnsiColorMap  = NewColorMap(DefaultColorList)
	Config   Configuration = NewConfig()
)

func init() {
	Config.ColorMap = ColorMap
}

type (
	Configuration interface {
		SetColorMap(c AnsiColorMap)
	}
	configuration struct{}
)

func NewConfig() *configuration {
	return &configuration{}
}

// GetEnv returns the current value of the environment variable 'name'
//
// If the value is not set or is "", an error is returned. The
// error is of type *EnvironmentError and wraps os.ErrNotExist
// for use with errors.Is()
func GetEnv(name string) (string, error) {
	env := os.Getenv(name)
	if env == "" {
		nameUp := strings.ToUpper(name)
		env = os.Getenv(nameUp)
		if env == "" {
			return "", NewEnvironmentError(name)
		}
	}
	return env, nil
}
