package emailt

import (
	"fmt"
	"html/template"
	"io"
)

type Element interface {
	Render(writer io.Writer) error
}

type StringElement string

func (e StringElement) Render(writer io.Writer) error {
	_, err := writer.Write([]byte(e))
	return err
}

type TemplateElement struct {
	Data     interface{}
	Template string
}

func (e TemplateElement) Render(writer io.Writer) error {
	t, err := template.New("").Parse(e.Template)
	if err != nil {
		return fmt.Errorf("parse template: %w", err)
	}
	return t.Execute(writer, e.Data)
}
