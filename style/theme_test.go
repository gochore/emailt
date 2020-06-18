package style

import (
	"reflect"
	"testing"
)

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
					Key: "a",
					Val: "aa",
				},
				{
					Key: "b",
					Val: "bb",
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

func TestMapTheme_Attributes(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name string
		t    MapTheme
		args args
		want Attributes
	}{
		{
			name: "regular",
			t: MapTheme{
				"a": Attributes{
					{
						Key: "an",
						Val: "av",
					},
				},
			},
			args: args{
				tag: "a",
			},
			want: Attributes{
				{
					Key: "an",
					Val: "av",
				},
			},
		},
		{
			name: "nil",
			t:    nil,
			args: args{
				tag: "a",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Attributes(tt.args.tag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Attributes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapTheme_Exists(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name string
		t    MapTheme
		args args
		want bool
	}{
		{
			name: "regular",
			t: MapTheme{
				"a": Attributes{
					{
						Key: "an",
						Val: "av",
					},
				},
			},
			args: args{
				tag: "a",
			},
			want: true,
		},
		{
			name: "nil",
			t:    nil,
			args: args{
				tag: "a",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Exists(tt.args.tag); got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}
