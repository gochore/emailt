package htmlt

import (
	"fmt"
)

type Html = string

// Sprintf return a html element
func Sprintf(format string, a ...interface{}) Html {
	return Html(fmt.Sprintf(format, a...))
}

// T return a html element with specified tag
func T(tag string, format string, a ...interface{}) Html {
	return Sprintf(fmt.Sprintf("<%s>%s</%s>", tag, format, tag), a...)
}

// H return a html element <h>
func H(level int, format string, a ...interface{}) Html {
	if level < 1 {
		level = 1
	}
	if level > 6 {
		level = 6
	}
	return T(fmt.Sprintf("h%d", level), format, a...)
}

// Hr return a html element <hr>
func Hr() Html {
	return Sprintf("<hr/>")
}

// P return a html element <p>
func P(format string, a ...interface{}) Html {
	return T("p", format, a...)
}

// Pre return a html element <pre>
func Pre(format string, a ...interface{}) Html {
	return T("pre", format, a...)
}

// A return a html element <a>
func A(href, format string, a ...interface{}) Html {
	return Sprintf(fmt.Sprintf(`<a href="%s">%s</a>`, href, format), a...)
}

// B return a html element <b>
func B(format string, a ...interface{}) Html {
	return T("b", format, a...)
}

// Br return a html element <br>
func Br() Html {
	return Sprintf("<br/>")
}

// Code return a html element <code>
func Code(format string, a ...interface{}) Html {
	return T("code", format, a...)
}

// Em return a html element <em>
func Em(format string, a ...interface{}) Html {
	return T("em", format, a...)
}

// I return a html element <i>
func I(format string, a ...interface{}) Html {
	return T("i", format, a...)
}

// Small return a html element <small>
func Small(format string, a ...interface{}) Html {
	return T("small", format, a...)
}

// Strong return a html element <strong>
func Strong(format string, a ...interface{}) Html {
	return T("strong", format, a...)
}

// Img return a html element <img>
func Img(src, alt, format string, a ...interface{}) Html {
	return Sprintf(fmt.Sprintf(`<a src="%s" alt="%s">%s</a>`, src, alt, format), a...)
}

// Del return a html element <del>
func Del(format string, a ...interface{}) Html {
	return T("del", format, a...)
}

// Ins return a html element <ins>
func Ins(format string, a ...interface{}) Html {
	return T("ins", format, a...)
}
