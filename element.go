package emailt

import (
	"bytes"
	"fmt"
	"html/template"
	"io"

	"github.com/gochore/emailt/internal/rend"
	"github.com/gochore/emailt/style"
)

type Element interface {
	Render(writer io.Writer, themes ...style.Theme) error
}

type TemplateElement struct {
	Data     interface{}
	Template string
}

func (e TemplateElement) Render(writer io.Writer, themes ...style.Theme) error {
	t, err := template.New("").Parse(e.Template)
	if err != nil {
		return fmt.Errorf("parse template: %w", err)
	}
	buffer := &bytes.Buffer{}
	if err := t.Execute(buffer, e.Data); err != nil {
		return fmt.Errorf("template execute: %w", err)
	}
	return rend.RenderTheme(buffer, writer, rend.MergeThemes(themes))
}
