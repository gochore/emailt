package emailt

import (
	"io"
)

type Email struct {
	elements []Element
}

func (e *Email) AddElements(element ...Element) *Email {
	e.elements = append(e.elements, element...)
	return e
}

func (e *Email) Render(writer io.Writer) error {
	render := newFmtWriter(writer)
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
		if err := element.Render(writer); err != nil {
			return err
		}
	}

	render.Print(`
</body>
</html>
`)

	return render.Err()
}
