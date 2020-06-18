package emailt

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type fmtWriter struct {
	writer io.Writer
	err    error
}

func newFmtWriter(writer io.Writer) *fmtWriter {
	return &fmtWriter{
		writer: writer,
		err:    nil,
	}
}

func (w fmtWriter) Write(p []byte) (n int, err error) {
	if w.err != nil {
		return 0, w.err
	}
	n, err = w.writer.Write(p)
	w.err = err
	return
}

func (w *fmtWriter) Err() error {
	return w.err
}

func (w *fmtWriter) Print(a ...interface{}) {
	if w.err != nil {
		return
	}
	_, w.err = fmt.Fprint(w.writer, a...)
}

func (w *fmtWriter) Println(a ...interface{}) {
	if w.err != nil {
		return
	}
	_, w.err = fmt.Fprintln(w.writer, a...)
}

func (w *fmtWriter) Printlnf(format string, a ...interface{}) {
	w.Printf(format+"\n", a...)
}

func (w *fmtWriter) Printf(format string, a ...interface{}) {
	if w.err != nil {
		return
	}
	_, w.err = fmt.Fprintf(w.writer, format, a...)
}

func writeStyles(node *html.Node, theme Theme) {
	if node == nil {
		return
	}
	if node.Type == html.ElementNode {
		node.Attr = theme.Attributes(node.Data).Merge(node.Attr)
	}
	writeStyles(node.FirstChild, theme)
	writeStyles(node.NextSibling, theme)
}

func mergeThemes(themes []Theme) Theme {
	theme := ChainTheme{}
	for _, m := range themes {
		theme = ChainTheme{
			upstream: theme,
			inner:    m,
		}
	}
	return theme
}

func htmlRender(reader io.Reader, writer io.Writer, theme Theme) error {
	nodes, err := html.ParseFragment(reader, &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
	})
	if err != nil {
		return fmt.Errorf("ParseFragment: %w", err)
	}

	for _, node := range nodes {
		writeStyles(node, theme)
		if err := html.Render(writer, node); err != nil {
			return fmt.Errorf("html render: %w", err)
		}
	}
	return nil
}
