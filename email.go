package emailt

import (
	"fmt"
	"io"

	"github.com/gochore/emailt/internal/rend"
	"github.com/gochore/emailt/style"
)

type Email struct {
	elements []Element
	theme    style.Theme
}

type Option func(email *Email)

func WithTheme(theme style.Theme) Option {
	return func(email *Email) {
		email.theme = theme
	}
}

func NewEmail(options ...Option) *Email {
	ret := &Email{
		theme: DefaultTheme,
	}
	for _, option := range options {
		option(ret)
	}
	return ret
}

func (e *Email) AddElements(element ...Element) *Email {
	e.elements = append(e.elements, element...)
	return e
}

func (e *Email) Render(writer io.Writer) error {
	render := rend.NewFmtWriter(writer)
	render.Print(`
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <title></title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
</head>
<body>
`)

	for _, element := range e.elements {
		if err := element.Render(writer, e.theme); err != nil {
			return fmt.Errorf("render: %w", err)
		}
	}

	render.Print(`
</body>
</html>
`)

	return render.Err()
}
