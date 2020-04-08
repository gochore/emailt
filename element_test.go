package emailt

import (
	"bytes"
	"testing"
	"time"
)

func TestStringElement_Render(t *testing.T) {
	tests := []struct {
		name       string
		e          StringElement
		style      Theme
		wantWriter string
		wantErr    bool
	}{
		{
			name:       "regular",
			e:          "<p>test</p>",
			wantWriter: "<p>test</p>",
			wantErr:    false,
		},
		{
			name: "with_style",
			e:    "<p>test</p>",
			style: MapTheme{
				"p": Attributes{
					{
						Name:  "style",
						Value: "background-color:#dedede;",
					},
				},
			},
			wantWriter: `<p style="background-color:#dedede;">test</p>`,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			err := tt.e.Render(writer, tt.style)
			if (err != nil) != tt.wantErr {
				t.Errorf("Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("Render() gotWriter = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}

func TestTemplateElement_Render(t *testing.T) {
	type fields struct {
		Data     interface{}
		Template string
	}
	tests := []struct {
		name       string
		fields     fields
		wantWriter string
		wantErr    bool
	}{
		{
			name: "regular",
			fields: fields{
				Data: struct {
					A string
					B int
				}{
					A: "a",
					B: 1,
				},
				Template: "A:{{.A}}, B:{{.B}}",
			},
			wantWriter: "A:a, B:1",
			wantErr:    false,
		},
		{
			name: "invalid_template",
			fields: fields{
				Data: struct {
					A string
					B int
				}{
					A: "a",
					B: 1,
				},
				Template: "A:{{.A}}, B:{{.B}",
			},
			wantWriter: "",
			wantErr:    true,
		},
		{
			name: "invalid_data",
			fields: fields{
				Data:     "test",
				Template: "A:{{.A}}, B:{{.B}}",
			},
			wantWriter: "",
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := TemplateElement{
				Data:     tt.fields.Data,
				Template: tt.fields.Template,
			}
			writer := &bytes.Buffer{}
			err := e.Render(writer)
			if (err != nil) != tt.wantErr {
				t.Errorf("Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("Render() gotWriter = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}

func TestNewStringElement(t *testing.T) {
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name string
		args args
		want StringElement
	}{
		{
			name: "regular",
			args: args{
				format: "%v %v",
				a:      []interface{}{1, time.Minute},
			},
			want: "1 1m0s",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStringElement(tt.args.format, tt.args.a...); got != tt.want {
				t.Errorf("NewStringElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
