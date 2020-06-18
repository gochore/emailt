package htmlt

import (
	"fmt"
	"io"
	"strings"

	"github.com/gochore/emailt/internal/rend"
	"github.com/gochore/emailt/style"
)

type Html string

func New(format string, a ...interface{}) Html {
	return Html(fmt.Sprintf(format, a...))
}

func (e Html) Render(writer io.Writer, themes ...style.Theme) error {
	return rend.RenderTheme(strings.NewReader(string(e)), writer, rend.MergeThemes(themes))
}
