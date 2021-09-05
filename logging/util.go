package logging

import (
	"os"
	"strings"
)

var ColorMap = map[string]string{}

var ColorList = []string{
	"NORMAL",
	"BOLD",
	"DIM",
	"ITALIC",
	"UNDERLINE",
	"REVERSED",
	"STRIKEOUT",
	"ATTN",
	"BLUE",
	"GOLANG",
	"CANARY",
	"CHERRY",
	"COOL",
	"CYAN",
	"DARKGREEN",
	"DBLUE",
	"GREEN",
	"LIME",
	"MAGENTA",
	"MAIN",
	"ORANGE",
	"PINK",
	"PURPLE",
	"RAIN",
	"RED",
	"TANGERINE",
	"WARN",
	"WHITE",
	"YELLOW",
	"RESTORE",
	"RESET",
}

// init loads the ColorMap from environment variables named in ColorList
func init() {
	for _, c := range ColorList {
		if _, ok := ColorMap[c]; !ok { // skip values already loaded
			a, err := GetEnv(c)
			// empty values are automatically considered errors by GetEnv()
			if Err(err) == nil {
				ColorMap[c] = AnsiClean(a)
			}
		}
	}
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

// AnsiClean checks a string specifically for malformed ANSI
// color codes.
// These codes are often loaded from environment variables
// and this function may help prevent malformed escape code
// sequences from being used.
//
// The allowed pattern is:
//  leading ESC + '['
//  any number of digits or semicolons
//  trailing 'm'
//
// All other patterns are rejected and "" is returned. This
// most often results in an empty string being passed to
// print functions, which has no effect.
func AnsiClean(s string) string {

	b := []byte(s)
	l := len(b) - 1

	if b[0] != '\x1b' || b[1] != '[' || b[l] != 'm' {
		return ""
	}

	sb := strings.Builder{}
	defer sb.Reset()

	sb.WriteByte(b[0])
	sb.WriteByte(b[1])

	for _, c := range b[2:l] {
		if c != ';' && (c < '0' || c > '9') {
			return "" // fail fast on any error
		}
		sb.WriteByte(c)
	}

	sb.WriteByte(b[l])

	return sb.String()

}
