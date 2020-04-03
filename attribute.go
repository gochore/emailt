package emailt

import (
	"fmt"
	"strings"
)

type Attribute struct {
	Name  string
	Value string
}

type Attributes []Attribute

func (as Attributes) String() string {
	var strs []string
	for _, a := range as {
		strs = append(strs, fmt.Sprintf(`%s="%s"`, a.Name, a.Value))
	}
	return strings.Join(strs, " ")
}

func (as Attributes) Merge(newAs Attributes) Attributes {
	ret := make(Attributes, len(as))
	copy(ret, as)
	for _, v := range newAs {
		found := false
		for i, v2 := range ret {
			if v2.Name == v.Name {
				ret[i].Value = v.Value
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
