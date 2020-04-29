package emailt

import (
	"fmt"
)

type Headline = StringElement

func NewHeadline(level int, format string, a ...interface{}) Headline {
	if level < 1 {
		level = 1
	}
	if level > 6 {
		level = 6
	}
	format = fmt.Sprintf("<h%d>%s</h%d>", level, format, level)
	return Headline(fmt.Sprintf(format, a...))
}
