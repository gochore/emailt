package emailt

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

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

type ChainTheme struct {
	upstream Theme
	inner    Theme
}

func (t ChainTheme) Attributes(tag string) Attributes {
	if theme := t.inner; theme != nil && theme.Exists(tag) {
		return theme.Attributes(tag)
	}
	if theme := t.upstream; theme != nil {
		return theme.Attributes(tag)
	}
	return nil
}

func (t ChainTheme) Exists(tag string) bool {
	if theme := t.inner; theme != nil && theme.Exists(tag) {
		return true
	}
	theme := t.upstream
	return theme != nil && theme.Exists(tag)
}

var (
	DefaultTheme Theme = MapTheme{
		"table": Attributes{
			{Key: "style", Val: "border:1px black solid; padding:3px 3px 3px 3px; border-collapse:collapse;"},
		},
		"th": Attributes{
			{Key: "style", Val: "border:1px black solid; padding:3px 3px 3px 3px; border-collapse:collapse; background-color:#dedede;"},
		},
		"td": Attributes{
			{Key: "style", Val: "border:1px black solid; padding:3px 3px 3px 3px; border-collapse:collapse;"},
		},
	}
)

type Attribute struct {
	Key string
	Val string
}

type Attributes []html.Attribute

func (as Attributes) String() string {
	var strs []string
	for _, a := range as {
		strs = append(strs, fmt.Sprintf(`%s="%s"`, a.Key, a.Val))
	}
	return strings.Join(strs, " ")
}

func (as Attributes) Merge(newAs Attributes) Attributes {
	ret := make(Attributes, len(as))
	copy(ret, as)
	for _, v := range newAs {
		found := false
		for i, v2 := range ret {
			if v2.Key == v.Key {
				ret[i].Val = v.Val
				found = true
				break
			}
		}
		if !found {
			ret = append(ret, v)
		}
	}
	return ret
}
