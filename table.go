package emailt

import (
	"fmt"
	"html/template"
	"io"
	"reflect"
	"strings"
)

type Table struct {
	Dataset    []interface{}
	Columns    []Column
	Attr       Attributes
	HeaderAttr Attributes
	DataAttr   Attributes
}

func NewTable() Table {
	return Table{
		Attr:       DefaultTableAttr,
		HeaderAttr: DefaultTableHeaderAttr,
		DataAttr:   DefaultTableDataAttr,
	}
}

func (t Table) WithDataset(dataset []interface{}) Table {
	t.Dataset = dataset
	return t
}

func (t Table) WithColumns(columns []Column) Table {
	t.Columns = columns
	return t
}

func (t Table) Render(writer io.Writer) error {
	if len(t.Dataset) == 0 {
		return fmt.Errorf("empty data")
	}

	typ := reflect.TypeOf(t.Dataset[0])
	if typ.Kind() != reflect.Struct {
		return fmt.Errorf("%v is not a struct", typ)
	}

	for i, v := range t.Dataset {
		t := reflect.TypeOf(v)
		if t != typ {
			return fmt.Errorf("item %v is %v, not %v", i, t, typ)
		}
	}

	columns := t.Columns
	if len(columns) == 0 {
		numField := typ.NumField()
		for i := 0; i < numField; i++ {
			field := typ.Field(i)
			columns = append(columns, Column{
				Name:     field.Name,
				Template: fmt.Sprintf("{{.%s}}", field.Name),
			})
		}
	}
	var rowTemplate *template.Template
	{
		rowBuilder := strings.Builder{}
		rowBuilder.WriteString("<tr>\n")
		for _, column := range columns {
			rowBuilder.WriteString(fmt.Sprintf("<td %s>%s</td>\n", t.DataAttr, column.Template))
		}
		rowBuilder.WriteString("</tr>")
		var err error
		rowTemplate, err = template.New("").Parse(rowBuilder.String())
		if err != nil {
			return fmt.Errorf("Parse: %w", err)
		}
	}

	render := newFmtWriter(writer)

	render.Printlnf("<table %s>", t.Attr.String())

	render.Println("<tr>")

	for _, column := range columns {
		render.Printlnf("<th %s>%s</th>", t.HeaderAttr, column.Name)
	}
	render.Println("</tr>")

	for _, row := range t.Dataset {
		if err := rowTemplate.Execute(render, row); err != nil {
			return fmt.Errorf("Execute: %w", err)
		}
		render.Println()
	}

	render.Println("</table>")

	return render.Err()
}

type Column struct {
	Name     string
	Template string
}

var (
	DefaultTableAttr = Attributes{
		{Name: "style", Value: "font-family: verdana,arial,sans-serif;font-size:14px;color:#333333;border-width: 1px;border-color: #666666;border-collapse: collapse;"},
	}
	DefaultTableHeaderAttr = Attributes{
		{Name: "style", Value: "border-width: 1px;padding: 8px;border-style: solid;border-color: #666666;background-color: #dedede;"},
	}
	DefaultTableDataAttr = Attributes{
		{Name: "style", Value: "border-width: 1px;padding: 8px;border-style: solid;border-color: #666666;background-color: #ffffff"},
	}
)
