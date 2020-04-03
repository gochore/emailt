package emailt

import (
	"io"
)

type Element interface {
	Render(writer io.Writer) error
}

type HtmlElement string

func (e HtmlElement) Render(writer io.Writer) error {
	_, err := writer.Write([]byte(e))
	return err
}
