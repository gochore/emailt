package htmlt

import (
	"testing"
)

func TestH(t *testing.T) {
	type args struct {
		level  int
		format string
		a      []interface{}
	}
	tests := []struct {
		name string
		args args
		want Html
	}{
		{
			name: "regular",
			args: args{
				level:  5,
				format: "%v %v",
				a:      []interface{}{1, "A"},
			},
			want: "<h5>1 A</h5>",
		},
		{
			name: "h0",
			args: args{
				level:  0,
				format: "%v %v",
				a:      []interface{}{1, "A"},
			},
			want: "<h1>1 A</h1>",
		},
		{
			name: "h7",
			args: args{
				level:  7,
				format: "%v %v",
				a:      []interface{}{1, "A"},
			},
			want: "<h6>1 A</h6>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := H(tt.args.level, tt.args.format, tt.args.a...); got != tt.want {
				t.Errorf("NewHeadline() = %v, want %v", got, tt.want)
			}
		})
	}
}
