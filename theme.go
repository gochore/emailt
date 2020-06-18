package emailt

import "github.com/gochore/emailt/style"

var (
	DefaultTheme style.Theme = style.MapTheme{
		"table": style.Attributes{
			{Key: "style", Val: "border:1px black solid; padding:3px 3px 3px 3px; border-collapse:collapse;"},
		},
		"th": style.Attributes{
			{Key: "style", Val: "border:1px black solid; padding:3px 3px 3px 3px; border-collapse:collapse; background-color:#dedede;"},
		},
		"td": style.Attributes{
			{Key: "style", Val: "border:1px black solid; padding:3px 3px 3px 3px; border-collapse:collapse;"},
		},
	}
)
