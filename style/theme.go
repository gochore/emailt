package style

type Theme interface {
	Attributes(tag string) Attributes
	Exists(tag string) bool
}

type MapTheme map[string]Attributes

func (t MapTheme) Attributes(tag string) Attributes {
	if len(t) == 0 {
		return nil
	}
	return t[tag]
}

func (t MapTheme) Exists(tag string) bool {
	if len(t) == 0 {
		return false
	}
	_, ok := t[tag]
	return ok
}

type Attribute struct {
	Key string
	Val string
}
