package emailt

import (
	"fmt"
	"io"
)

type Element interface {
	Render(writer io.Writer) error
}

type fmtWriter struct {
	writer io.Writer
	err    error
}

func (w fmtWriter) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)
}

func newFmtWriter(writer io.Writer) *fmtWriter {
	return &fmtWriter{
		writer: writer,
		err:    nil,
	}
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
	if w.err != nil {
		return
	}
	_, w.err = fmt.Fprintf(w.writer, format+"\n", a...)
}
