package emailt

import (
	"bytes"
	"fmt"
	"io"
	"text/template"

	"github.com/gochore/emailt/internal/rend"
	"github.com/gochore/emailt/style"
)

type Element interface {
	Render(writer io.Writer, themes ...style.Theme) error
}

type TemplateElement struct {
	Data     interface{}
	Template string
	Funcs    template.FuncMap
}

func (e TemplateElement) Render(writer io.Writer, themes ...style.Theme) error {
	errPrefix := "TemplateElement.Render: "

	t, err := template.New("").Funcs(e.Funcs).Parse(e.Template)
	if err != nil {
		return fmt.Errorf(errPrefix+"parse template: %w", err)
	}
	buffer := &bytes.Buffer{}
	if err := t.Execute(buffer, e.Data); err != nil {
		return fmt.Errorf(errPrefix+"template execute: %w", err)
	}
	if err := rend.RenderTheme(buffer, writer, rend.MergeThemes(themes)); err != nil {
		return fmt.Errorf(errPrefix+"%w", err)
	}
	return nil
}
