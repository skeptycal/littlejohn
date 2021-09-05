package config

import "strings"

// AnsiClean is a utility that checks a string specifically
// for malformed ANSI color codes.
// These codes are often loaded from environment variables
// and this function may help prevent malformed escape code
// sequences from being imported or processed.
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
