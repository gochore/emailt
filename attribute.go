package emailt

import (
	"fmt"
	"strings"
)

type Attribute struct {
	Name  string
	Value string
}

func (a Attribute) Render() (string, error) {
	if a.Name == "" {
		return "", fmt.Errorf("empty name")
	}
	return fmt.Sprintf(`%s="%s"`, a.Name, a.Value), nil
}

type Attributes []Attribute

func (as Attributes) Render() (string, error) {
	var strs []string
	for _, a := range as {
		t, err := a.Render()
		if err != nil {
			return "", err
		}
		strs = append(strs, t)
	}
	return strings.Join(strs, " "), nil
}
