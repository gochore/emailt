package emailt

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"strings"
)

type Element interface {
	Render(writer io.Writer, themes ...Theme) error
}

type StringElement string

func (e StringElement) Render(writer io.Writer, themes ...Theme) error {
	return htmlRender(strings.NewReader(string(e)), writer, mergeThemes(themes))
}

type TemplateElement struct {
	Data     interface{}
	Template string
}

func (e TemplateElement) Render(writer io.Writer, themes ...Theme) error {
	t, err := template.New("").Parse(e.Template)
	if err != nil {
		return fmt.Errorf("parse template: %w", err)
	}
	buffer := &bytes.Buffer{}
	if err := t.Execute(buffer, e.Data); err != nil {
		return err
	}
	return htmlRender(buffer, writer, mergeThemes(themes))
}
