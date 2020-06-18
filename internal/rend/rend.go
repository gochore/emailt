package rend

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"github.com/gochore/emailt/style"
)

type FmtWriter struct {
	writer io.Writer
	err    error
}

func NewFmtWriter(writer io.Writer) *FmtWriter {
	return &FmtWriter{
		writer: writer,
		err:    nil,
	}
}

func (w FmtWriter) Write(p []byte) (n int, err error) {
	if w.err != nil {
		return 0, w.err
	}
	n, err = w.writer.Write(p)
	w.err = err
	return
}

func (w *FmtWriter) Err() error {
	return w.err
}

func (w *FmtWriter) Print(a ...interface{}) {
	if w.err != nil {
		return
	}
	_, w.err = fmt.Fprint(w.writer, a...)
}

func (w *FmtWriter) Println(a ...interface{}) {
	if w.err != nil {
		return
	}
	_, w.err = fmt.Fprintln(w.writer, a...)
}

func (w *FmtWriter) Printlnf(format string, a ...interface{}) {
	w.Printf(format+"\n", a...)
}

func (w *FmtWriter) Printf(format string, a ...interface{}) {
	if w.err != nil {
		return
	}
	_, w.err = fmt.Fprintf(w.writer, format, a...)
}

func WriteTheme(node *html.Node, theme style.Theme) {
	if node == nil {
		return
	}
	if node.Type == html.ElementNode {
		node.Attr = theme.Attributes(node.Data).Merge(node.Attr)
	}
	WriteTheme(node.FirstChild, theme)
	WriteTheme(node.NextSibling, theme)
}

func MergeThemes(themes []style.Theme) style.Theme {
	theme := ChainTheme{}
	for _, m := range themes {
		theme = ChainTheme{
			Upstream: theme,
			Inner:    m,
		}
	}
	return theme
}

func RenderTheme(reader io.Reader, writer io.Writer, theme style.Theme) error {
	nodes, err := html.ParseFragment(reader, &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
	})
	if err != nil {
		return fmt.Errorf("ParseFragment: %w", err)
	}

	for _, node := range nodes {
		WriteTheme(node, theme)
		if err := html.Render(writer, node); err != nil {
			return fmt.Errorf("html render: %w", err)
		}
	}
	return nil
}
