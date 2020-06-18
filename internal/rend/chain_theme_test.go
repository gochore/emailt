package rend

import (
	"testing"

	"github.com/gochore/emailt/style"
)

func TestChainTheme_Exists(t1 *testing.T) {
	type fields struct {
		upstream style.Theme
		inner    style.Theme
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
				upstream: style.MapTheme{
					"b": style.Attributes{
						{
							Key: "bn",
							Val: "bv",
						},
					},
				},
				inner: style.MapTheme{
					"a": style.Attributes{
						{
							Key: "an",
							Val: "av",
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
				upstream: style.MapTheme{
					"b": style.Attributes{
						{
							Key: "bn",
							Val: "bv",
						},
					},
				},
				inner: style.MapTheme{
					"a": style.Attributes{
						{
							Key: "an",
							Val: "av",
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
				Upstream: tt.fields.upstream,
				Inner:    tt.fields.inner,
			}
			if got := t.Exists(tt.args.tag); got != tt.want {
				t1.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}
