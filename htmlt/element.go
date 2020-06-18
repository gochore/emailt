package htmlt

import (
	"fmt"
	"io"
	"strings"

	"github.com/gochore/emailt/internal/rend"
	"github.com/gochore/emailt/style"
)

type Element Html

// Render implement email.Element
func (e Element) Render(writer io.Writer, themes ...style.Theme) error {
	errPrefix := "Html.Render: "
	if err := rend.RenderTheme(strings.NewReader(string(e)), writer, rend.MergeThemes(themes)); err != nil {
		return fmt.Errorf(errPrefix+"%w", err)
	}
	return nil
}
