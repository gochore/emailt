package emailt

import (
	"bytes"
	"testing"
)

func TestStringElement_Render(t *testing.T) {
	tests := []struct {
		name       string
		e          StringElement
		wantWriter string
		wantErr    bool
	}{
		{
			name:       "regular",
			e:          "<p>test</p>",
			wantWriter: "<p>test</p>",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			err := tt.e.Render(writer)
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
