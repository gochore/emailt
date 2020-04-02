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

func TestTable_Render(t1 *testing.T) {

	type fields struct {
		Data    []interface{}
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
				Data: []interface{}{
					TestStruct1{
						A: "a1",
						B: 1,
					},
					TestStruct1{
						A: "a2",
						B: 2,
					},
					TestStruct1{
						A: "a3",
						B: 3,
					},
				},
				Columns: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Table{
				Data: tt.fields.Data,
			}
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
			_ = os.Mkdir(dir, os.ModeDir)
			_ = ioutil.WriteFile(filepath.Join(dir, fmt.Sprintf("%s.html", tt.name)), got.Bytes(), 0644)
		})
	}
}
