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

func TestEmail_Render(t *testing.T) {
	type fields struct {
		elements []Element
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "regular",
			fields: fields{
				elements: []Element{
					Table{
						Dataset: []TestStruct1{
							{
								A: "a1",
								B: 1,
							},
							{
								A: "a2",
								B: 2,
							},
							{
								A: "a3",
								B: 3,
							},
						},
						Columns: []Column{
							{
								Name:     "列1",
								Template: "{{.A}}",
							},
							{
								Name:     "列2",
								Template: "{{.B}}",
							},
							{
								Name:     "列3",
								Template: "{{.A}}({{.B}})",
							},
						},
						Attr:       DefaultTableAttr,
						HeaderAttr: DefaultTableHeaderAttr,
						DataAttr:   DefaultTableDataAttr,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Email{
				elements: tt.fields.elements,
			}
			got := bytes.NewBuffer(nil)
			err := e.Render(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if _, err := html.Parse(bytes.NewReader(got.Bytes())); err != nil {
				t.Error(err)
			}
			t.Log(got.String())

			dir := "output"
			_ = os.Mkdir(dir, 0755)
			_ = ioutil.WriteFile(filepath.Join(dir, fmt.Sprintf("TestEmail_Render.%s.html", tt.name)), got.Bytes(), 0644)
		})
	}
}
