package dumb_regexp

import (
	"strings"
)

func Fix(val string) string {
	var b strings.Builder

	// escape non-terminal $
	runes := strings.Split(val, "")
	for i, c := range runes {
		if c == "$" && i != len(runes)-1 && !(i > 0 && runes[i-1] == "\\") {
			b.WriteRune('\\')
		}
		b.WriteString(c)
	}

	// escape non-closed (,[
	// escape non-opened ),]

	return b.String()
}
