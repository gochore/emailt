package emailt

import (
	"fmt"
	"io"
)

type List struct {
	items   []Element
	ordered bool
}

func NewUnorderedList() *List {
	return &List{}
}

func NewOrderedList() *List {
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

func (l *List) Render(writer io.Writer, themes ...Theme) error {
	theme := mergeThemes(themes)

	render := newFmtWriter(writer)

	tag := "ul"
	if l.ordered {
		tag = "ol"
	}

	render.Printlnf("<%s %s>", tag, theme.Attributes(tag))

	for _, item := range l.items {
		render.Printf("<li %s>", theme.Attributes("li"))
		if err := item.Render(render, theme); err != nil {
			return fmt.Errorf("render: %w", err)
		}
		render.Printlnf("</li>")
	}

	render.Printlnf("</%s>", tag)

	return render.Err()
}
