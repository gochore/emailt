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

func TestList_Render(t *testing.T) {
	type fields struct {
		items   []Element
		ordered bool
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
			name: "unordered",
			fields: fields{
				items: []Element{
					NewStringElement("A"),
					NewStringElement("B"),
				},
				ordered: false,
			},
			args:    args{},
			wantErr: false,
		},
		{
			name: "ordered",
			fields: fields{
				items: []Element{
					NewStringElement("A"),
					NewStringElement("B"),
				},
				ordered: true,
			},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewUnorderedList()
			if tt.fields.ordered {
				l = NewOrderedList()
			}
			l.Add(tt.fields.items...)
			writer := &bytes.Buffer{}
			err := l.Render(writer, tt.args.themes...)
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
			_ = ioutil.WriteFile(filepath.Join(dir, fmt.Sprintf("TestList_Render.%s.html", tt.name)), writer.Bytes(), 0644)
		})
	}
}
