package config

import (
	"errors"
	"os"
	"strings"

	"github.com/skeptycal/errorlogger"
	"github.com/skeptycal/gofile"
)

var (
	Log                    = errorlogger.Log
	Err                    = errorlogger.Err
	ColorMap AnsiColorMap  = NewColorMap(DefaultColorList)
	Config   Configuration = NewConfig()
)

func init() {
	Config.SetColorMap(ColorMap)
}

type (
	Configuration interface {
		SetColorMap(c AnsiColorMap)
	}
	configuration struct {
		configPath string
		colormap   AnsiColorMap
	}
)

func NewConfig() *configuration {
	return &configuration{
		colormap: NewColorMap(DefaultColorList),
	}
}

func (c *configuration) SetColorMap(a AnsiColorMap) error {
	if a == nil {
		if c.colormap == nil {
			c.colormap = NewColorMap(DefaultColorList)
			return errors.New("colormap not supplied, using default")
		}
		return errors.New("colormap not supplied, using existing")
	}
	c.colormap = a
	return nil
}

func (c *configuration) SetPath(path string) error {
	if path == "" {
		return errors.New("path not supplied")
	}

	gofile.IsDir()
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}

	c.configPath = fi.Name()

	c.configPath = "."
	return nil
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
