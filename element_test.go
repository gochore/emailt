package emailt

import (
	"bytes"
	"testing"
)

func TestStringElement_Render(t *testing.T) {
	tests := []struct {
		name       string
		e          HtmlElement
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
