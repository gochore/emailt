package emailt

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/net/html"
)

func TestHeadline_Render(t *testing.T) {
	type fields struct {
		level  int
		format string
		a      []interface{}
	}
	type args struct {
		themes []Theme
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "regular",
			fields: fields{
				level:  1,
				format: "%v %v",
				a:      []interface{}{1, "A"},
			},
			args:    args{},
			wantErr: false,
		},
		{
			name: "h-1",
			fields: fields{
				level:  -1,
				format: "%v %v",
				a:      []interface{}{1, "A"},
			},
			args:    args{},
			wantErr: false,
		},
		{
			name: "h7",
			fields: fields{
				level:  7,
				format: "%v %v",
				a:      []interface{}{1, "A"},
			},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHeadline(tt.fields.level, tt.fields.format, tt.fields.a...)

			writer := &bytes.Buffer{}
			err := h.Render(writer, tt.args.themes...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if _, err := html.Parse(bytes.NewReader(writer.Bytes())); err != nil {
				t.Error(err)
			}
			t.Log(writer.String())

			dir := "output"
			_ = os.Mkdir(dir, 0744)
			_ = ioutil.WriteFile(filepath.Join(dir, fmt.Sprintf("TestHeadline_Render.%s.html", tt.name)), writer.Bytes(), 0644)
		})
	}
}
