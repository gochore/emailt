package emailt

import (
	"bytes"
	"strings"
	"testing"
	"text/template"
)

func TestTemplateElement_Render(t *testing.T) {
	type fields struct {
		Data     interface{}
		Template string
		Funcs    template.FuncMap
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
		{
			name: "with_funcs",
			fields: fields{
				Data: struct {
					A string
					B int
				}{
					A: "hello",
					B: 1,
				},
				Template: "A:{{title .A}}, B:{{.B}}",
				Funcs: template.FuncMap{
					"title": strings.Title,
				},
			},
			wantWriter: "A:Hello, B:1",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Template{
				Data:     tt.fields.Data,
				Template: tt.fields.Template,
				Funcs:    tt.fields.Funcs,
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
