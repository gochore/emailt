package emailt

import "testing"

func TestAttributes_String(t *testing.T) {
	tests := []struct {
		name string
		as   Attributes
		want string
	}{
		{
			name: "regular",
			as: Attributes{
				{
					Name:  "a",
					Value: "aa",
				},
				{
					Name:  "b",
					Value: "bb",
				},
			},
			want: `a="aa" b="bb"`,
		},
		{
			name: "nil",
			as:   nil,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.as.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
