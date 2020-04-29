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

type TestStruct1 struct {
	A string
	B int
}

func TestTable_Render(t1 *testing.T) {
	type fields struct {
		Dataset interface{}
		Columns []Column
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "regular",
			fields: fields{
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
				Columns: nil,
			},
			wantErr: false,
		},
		{
			name: "with_columns",
			fields: fields{
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
			},
			wantErr: false,
		},
		{
			name: "map_dataset",
			fields: fields{
				Dataset: []map[string]interface{}{
					{
						"A": "a1",
						"B": 1,
					},
					{
						"A": "a2",
						"B": 2,
					},
					{
						"A": "a3",
						"B": 3,
					},
				},
				Columns: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := NewTable()
			t.SetColumns(tt.fields.Columns)
			t.SetDataset(tt.fields.Dataset)
			got := bytes.NewBuffer(nil)
			err := t.Render(got)
			if (err != nil) != tt.wantErr {
				t1.Errorf("Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if _, err := html.Parse(bytes.NewReader(got.Bytes())); err != nil {
				t1.Error(err)
			}
			t1.Log(got.String())

			dir := "output"
			_ = os.Mkdir(dir, 0744)
			_ = ioutil.WriteFile(filepath.Join(dir, fmt.Sprintf("TestTable_Render.%s.html", tt.name)), got.Bytes(), 0644)
		})
	}
}
