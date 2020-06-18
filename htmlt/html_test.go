package htmlt

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"golang.org/x/net/html"

	"github.com/gochore/emailt/style"
)

func TestT(t *testing.T) {
	type args struct {
		tag    string
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
				tag:    "h1",
				format: "abc%d",
				a:      []interface{}{1},
			},
			want: "<h1>abc1</h1>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := T(tt.args.tag, tt.args.format, tt.args.a...); got != tt.want {
				t.Errorf("T() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func TestHr(t *testing.T) {
	tests := []struct {
		name string
		want Html
	}{
		{
			name: "regular",
			want: "<hr/>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hr(); got != tt.want {
				t.Errorf("Hr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestP(t *testing.T) {
	type args struct {
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
				format: "abc%d",
				a:      []interface{}{1},
			},
			want: "<p>abc1</p>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P(tt.args.format, tt.args.a...); got != tt.want {
				t.Errorf("P() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPre(t *testing.T) {
	type args struct {
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
				format: "abc%d",
				a:      []interface{}{1},
			},
			want: "<pre>abc1</pre>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pre(tt.args.format, tt.args.a...); got != tt.want {
				t.Errorf("Pre() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestA(t *testing.T) {
	type args struct {
		href   string
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
				href:   "http://example.com",
				format: "abc%d",
				a:      []interface{}{1},
			},
			want: `<a href="http://example.com">abc1</a>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := A(tt.args.href, tt.args.format, tt.args.a...); got != tt.want {
				t.Errorf("A() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestB(t *testing.T) {
	type args struct {
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
				format: "abc%d",
				a:      []interface{}{1},
			},
			want: "<b>abc1</b>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := B(tt.args.format, tt.args.a...); got != tt.want {
				t.Errorf("B() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBr(t *testing.T) {
	tests := []struct {
		name string
		want Html
	}{
		{
			name: "regular",
			want: "<br/>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Br(); got != tt.want {
				t.Errorf("Br() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCode(t *testing.T) {
	type args struct {
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
				format: "abc%d",
				a:      []interface{}{1},
			},
			want: "<code>abc1</code>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code(tt.args.format, tt.args.a...); got != tt.want {
				t.Errorf("Code() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDel(t *testing.T) {
	type args struct {
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
				format: "abc%d",
				a:      []interface{}{1},
			},
			want: "<del>abc1</del>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Del(tt.args.format, tt.args.a...); got != tt.want {
				t.Errorf("Del() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEm(t *testing.T) {
	type args struct {
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
				format: "abc%d",
				a:      []interface{}{1},
			},
			want: "<em>abc1</em>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Em(tt.args.format, tt.args.a...); got != tt.want {
				t.Errorf("Em() = %v, want %v", got, tt.want)
			}
		})
	}
}

type errWriter struct{}

func (errWriter) Write(p []byte) (int, error) {
	return 0, fmt.Errorf("error")
}

func TestHtml_Render(t *testing.T) {
	type args struct {
		writer io.Writer
		themes []style.Theme
	}
	tests := []struct {
		name       string
		e          Html
		args       args
		wantWriter string
		wantErr    bool
	}{
		{
			name: "regular",
			e:    "<p>abc</p>",
			args: args{
				writer: &bytes.Buffer{},
				themes: []style.Theme{style.MapTheme{
					"p": []html.Attribute{
						{
							Key: "a",
							Val: "1",
						},
					},
				}},
			},
			wantWriter: `<p a="1">abc</p>`,
			wantErr:    false,
		},
		{
			name: "writer with error",
			e:    "<p>abc</p>",
			args: args{
				writer: errWriter{},
				themes: []style.Theme{style.MapTheme{
					"p": []html.Attribute{
						{
							Key: "a",
							Val: "1",
						},
					},
				}},
			},
			wantWriter: `<p a="1">abc</p>`,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.e.Render(tt.args.writer, tt.args.themes...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			stringer := tt.args.writer.(fmt.Stringer)
			if gotWriter := stringer.String(); gotWriter != tt.wantWriter {
				t.Errorf("Render() gotWriter = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}

func TestI(t *testing.T) {
	type args struct {
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
				format: "abc%d",
				a:      []interface{}{1},
			},
			want: "<i>abc1</i>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := I(tt.args.format, tt.args.a...); got != tt.want {
				t.Errorf("I() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImg(t *testing.T) {
	type args struct {
		src    string
		alt    string
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
				src:    "http://example.com",
				alt:    "example",
				format: "abc%d",
				a:      []interface{}{1},
			},
			want: `<a src="http://example.com" alt="example">abc1</a>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Img(tt.args.src, tt.args.alt, tt.args.format, tt.args.a...); got != tt.want {
				t.Errorf("Img() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIns(t *testing.T) {
	type args struct {
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
				format: "abc%d",
				a:      []interface{}{1},
			},
			want: "<ins>abc1</ins>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ins(tt.args.format, tt.args.a...); got != tt.want {
				t.Errorf("Ins() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSmall(t *testing.T) {
	type args struct {
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
				format: "abc%d",
				a:      []interface{}{1},
			},
			want: "<small>abc1</small>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Small(tt.args.format, tt.args.a...); got != tt.want {
				t.Errorf("Small() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSprintf(t *testing.T) {
	type args struct {
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
				format: "abc%d",
				a:      []interface{}{1},
			},
			want: "abc1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sprintf(tt.args.format, tt.args.a...); got != tt.want {
				t.Errorf("Sprintf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrong(t *testing.T) {
	type args struct {
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
				format: "abc%d",
				a:      []interface{}{1},
			},
			want: "<strong>abc1</strong>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Strong(tt.args.format, tt.args.a...); got != tt.want {
				t.Errorf("Strong() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNesting(t *testing.T) {
	var want Html = `<a href="/a"><b><code><del><em><h1><i><p>hello</p></i></h1></em></del></code></b></a>`
	got := A("/a", B(Code(Del(Em(H(1, I(P("hello"))))))))
	if got != want {
		t.Errorf("Sprintf() = %v, want %v", got, want)
	}
}
