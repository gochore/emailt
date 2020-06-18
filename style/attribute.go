package style

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

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
