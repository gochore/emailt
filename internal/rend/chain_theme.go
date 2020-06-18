package rend

import "github.com/gochore/emailt/style"

type ChainTheme struct {
	Upstream style.Theme
	Inner    style.Theme
}

func (t ChainTheme) Attributes(tag string) style.Attributes {
	if theme := t.Inner; theme != nil && theme.Exists(tag) {
		return theme.Attributes(tag)
	}
	if theme := t.Upstream; theme != nil {
		return theme.Attributes(tag)
	}
	return nil
}

func (t ChainTheme) Exists(tag string) bool {
	if theme := t.Inner; theme != nil && theme.Exists(tag) {
		return true
	}
	theme := t.Upstream
	return theme != nil && theme.Exists(tag)
}
