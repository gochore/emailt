package htmlt

import (
	"fmt"
)

func H(level int, format string, a ...interface{}) Html {
	if level < 1 {
		level = 1
	}
	if level > 6 {
		level = 6
	}
	format = fmt.Sprintf("<h%d>%s</h%d>", level, format, level)
	return Html(fmt.Sprintf(format, a...))
}
