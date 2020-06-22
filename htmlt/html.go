package htmlt

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gochore/emailt/internal/rend"
	"github.com/gochore/emailt/style"
)

type Html string

// Sprintf return a html element
func Sprintf(format string, a ...interface{}) Html {
	return Html(fmt.Sprintf(format, a...))
}

// Render implement email.Element
func (e Html) Render(writer io.Writer, themes ...style.Theme) error {
	errPrefix := "Html.Render: "
	if err := rend.RenderTheme(strings.NewReader(string(e)), writer, rend.MergeThemes(themes)); err != nil {
		return fmt.Errorf(errPrefix+"%w", err)
	}
	return nil
}

// T return a html element with specified tag
func T(tag string, format Html, a ...interface{}) Html {
	return Sprintf(fmt.Sprintf("<%s>%s</%s>", tag, format, tag), a...)
}

// H return a html element <h>
func H(level int, format Html, a ...interface{}) Html {
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
func P(format Html, a ...interface{}) Html {
	return T("p", format, a...)
}

// Pre return a html element <pre>
func Pre(format Html, a ...interface{}) Html {
	return T("pre", format, a...)
}

// A return a html element <a>
func A(href string, format Html, a ...interface{}) Html {
	return Sprintf(fmt.Sprintf(`<a href="%s">%s</a>`, href, format), a...)
}

// B return a html element <b>
func B(format Html, a ...interface{}) Html {
	return T("b", format, a...)
}

// Br return a html element <br>
func Br() Html {
	return Sprintf("<br/>")
}

// Code return a html element <code>
func Code(format Html, a ...interface{}) Html {
	return T("code", format, a...)
}

// Em return a html element <em>
func Em(format Html, a ...interface{}) Html {
	return T("em", format, a...)
}

// I return a html element <i>
func I(format Html, a ...interface{}) Html {
	return T("i", format, a...)
}

// Small return a html element <small>
func Small(format Html, a ...interface{}) Html {
	return T("small", format, a...)
}

// Strong return a html element <strong>
func Strong(format Html, a ...interface{}) Html {
	return T("strong", format, a...)
}

// Img return a html element <img>
func Img(src, alt string) Html {
	return Sprintf(`<img src="%s" alt="%s"/>`, src, alt)
}

// Img return a html element <img> with embedded data
func ImgEmbedded(image []byte, alt string) Html {
	src := &strings.Builder{}
	src.WriteString(fmt.Sprintf("data:%s;base64,", http.DetectContentType(image)))
	encoder := base64.NewEncoder(base64.StdEncoding, src)
	_, _ = encoder.Write(image)
	_ = encoder.Close()
	return Img(src.String(), alt)
}

// Del return a html element <del>
func Del(format Html, a ...interface{}) Html {
	return T("del", format, a...)
}

// Ins return a html element <ins>
func Ins(format Html, a ...interface{}) Html {
	return T("ins", format, a...)
}
