package emailt

import (
	"fmt"
	"html/template"
	"io"
	"reflect"
	"strings"
)

type Table struct {
	Data    []interface{}
	Columns []Column
}

func (t Table) Render(writer io.Writer) error {
	if len(t.Data) == 0 {
		return fmt.Errorf("empty data")
	}

	typ := reflect.TypeOf(t.Data[0])
	if typ.Kind() != reflect.Struct {
		return fmt.Errorf("%v is not a struct", typ)
	}

	for i, v := range t.Data {
		t := reflect.TypeOf(v)
		if t != typ {
			return fmt.Errorf("item %v is %v, not %v", i, t, typ)
		}
	}

	columns := make([]Column, len(t.Columns))
	copy(columns, t.Columns)
	if len(columns) == 0 {
		numField := typ.NumField()
		for i := 0; i < numField; i++ {
			field := typ.Field(i)
			columns = append(columns, Column{
				Name: field.Name,
			})
		}
	}

	render := newFmtWriter(writer)

	render.Println(`<table>`)

	render.Println(`<tr>`)

	var rowTplt *template.Template
	{
		rowBuilder := strings.Builder{}
		rowBuilder.WriteString(`<tr>`)
		for _, column := range columns {
			t := column.Template
			if t == "" {
				t = fmt.Sprintf("{{.%s}}", column.Name)
			}
			rowBuilder.WriteString(fmt.Sprintf(`<td>%s</td>`, t))
			render.Printlnf(`<th>%s</th>`, column.Name)
		}
		rowBuilder.WriteString(`</tr>`)
		var err error
		rowTplt, err = template.New("").Parse(rowBuilder.String())
		if err != nil {
			return fmt.Errorf("Parse: %w", err)
		}
	}

	render.Println(`</tr>`)

	for _, row := range t.Data {
		if err := rowTplt.Execute(render, row); err != nil {
			return fmt.Errorf("Execute: %w", err)
		}
		render.Println()
	}

	render.Println(`</table>`)

	return render.Err()
}

type Column struct {
	Name             string
	Template         string
	HeaderAttributes Attributes
	DataAttributes   Attributes
}
