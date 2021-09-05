package config

var DefaultColorList = []string{
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

type (
	AnsiColorMap map[string]string
)

// NewColorMap loads the ColorMap from environment variables
// named in colorList. Each entry is checked for malformed
// ANSI strings.
//
// If colorList is empty, the DefaultColorList is used
func NewColorMap(colorList []string) (acm AnsiColorMap) {
	if colorList == nil || len(colorList) == 0 {
		colorList = DefaultColorList
	}
	for _, c := range colorList {
		if _, ok := acm[c]; !ok {
			a, err := GetEnv(c)
			if Err(err) == nil && a != "" {
				acm[c] = AnsiClean(a)
			}
		}
	}
	return
}
