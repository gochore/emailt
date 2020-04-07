package emailt

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
						Name:  "an",
						Value: "av",
					},
				},
			},
			args: args{
				tag: "a",
			},
			want: Attributes{
				{
					Name:  "an",
					Value: "av",
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
						Name:  "an",
						Value: "av",
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

func TestChainTheme_Exists(t1 *testing.T) {
	type fields struct {
		upstream Theme
		inner    Theme
	}
	type args struct {
		tag string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "regular",
			fields: fields{
				upstream: MapTheme{
					"b": Attributes{
						{
							Name:  "bn",
							Value: "bv",
						},
					},
				},
				inner: MapTheme{
					"a": Attributes{
						{
							Name:  "an",
							Value: "av",
						},
					},
				},
			},
			args: args{
				tag: "a",
			},
			want: true,
		},
		{
			name: "not exist",
			fields: fields{
				upstream: MapTheme{
					"b": Attributes{
						{
							Name:  "bn",
							Value: "bv",
						},
					},
				},
				inner: MapTheme{
					"a": Attributes{
						{
							Name:  "an",
							Value: "av",
						},
					},
				},
			},
			args: args{
				tag: "c",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := ChainTheme{
				upstream: tt.fields.upstream,
				inner:    tt.fields.inner,
			}
			if got := t.Exists(tt.args.tag); got != tt.want {
				t1.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}
