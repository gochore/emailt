package emailt

import (
	"fmt"
	"io"

	"github.com/gochore/emailt/internal/rend"
	"github.com/gochore/emailt/style"
)

type List struct {
	items   []Element
	ordered bool
}

func NewUnordered() *List {
	return &List{}
}

func NewOrdered() *List {
	return &List{
		ordered: true,
	}
}

func (l *List) Ordered() bool {
	return l.ordered
}

func (l *List) Add(item ...Element) {
	l.items = append(l.items, item...)
}

func (l *List) Render(writer io.Writer, themes ...style.Theme) error {
	errPrefix := "List.Render: "

	theme := rend.MergeThemes(themes)

	render := rend.NewFmtWriter(writer)

	tag := "ul"
	if l.ordered {
		tag = "ol"
	}

	render.Printlnf("<%s %s>", tag, theme.Attributes(tag))

	for _, item := range l.items {
		render.Printf("<li %s>", theme.Attributes("li"))
		if err := item.Render(render, theme); err != nil {
			return fmt.Errorf(errPrefix+"render li: %w", err)
		}
		render.Printlnf("</li>")
	}

	render.Printlnf("</%s>", tag)

	if err := render.Err(); err != nil {
		return fmt.Errorf(errPrefix+"%w", err)
	}
	return nil
}
