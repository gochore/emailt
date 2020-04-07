package emailt

import (
	"fmt"
	"html/template"
	"io"
	"reflect"
	"sort"
	"strings"
)

type Table struct {
	Dataset    interface{}
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

func (t Table) WithDataset(dataset interface{}) Table {
	t.Dataset = dataset
	return t
}

func (t Table) WithColumns(columns []Column) Table {
	t.Columns = columns
	return t
}

func (t Table) Render(writer io.Writer) error {
	dataset := reflect.ValueOf(t.Dataset)
	if dataset.Kind() != reflect.Slice {
		return fmt.Errorf("%v is not a slice", dataset.Type())
	}

	if dataset.Len() == 0 {
		return fmt.Errorf("empty data")
	}

	mapItem := false
	typ := dataset.Index(0).Type()
	switch typ.Kind() {
	case reflect.Map:
		mapItem = true
	case reflect.Struct:
		// do nothing
	default:
		return fmt.Errorf("unsupported slice item type: %v", typ)
	}
	if typ.Kind() != reflect.Struct && typ.Kind() != reflect.Map {
		return fmt.Errorf("%v is not a struct", typ)
	}

	for i := 0; i < dataset.Len(); i++ {
		if t := dataset.Index(i).Type(); t != typ {
			return fmt.Errorf("item %v is %v, not %v", i, t, typ)
		}
	}

	columns := t.Columns
	if len(columns) == 0 {
		if mapItem {
			var keys []string
			for _, v := range dataset.Index(0).MapKeys() {
				keys = append(keys, fmt.Sprint(v.Interface()))
			}
			sort.Strings(keys)
			for _, key := range keys {
				columns = append(columns, Column{
					Name:     key,
					Template: fmt.Sprintf("{{.%s}}", key),
				})
			}
		} else {
			numField := typ.NumField()
			for i := 0; i < numField; i++ {
				field := typ.Field(i)
				columns = append(columns, Column{
					Name:     field.Name,
					Template: fmt.Sprintf("{{.%s}}", field.Name),
				})
			}
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

	for i := 0; i < dataset.Len(); i++ {
		if err := rowTemplate.Execute(render, dataset.Index(i)); err != nil {
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
		{Name: "style", Value: "border:1px black solid; padding:3px 3px 3px 3px; border-collapse:collapse;"},
	}
	DefaultTableHeaderAttr = Attributes{
		{Name: "style", Value: "border:1px black solid; padding:3px 3px 3px 3px; border-collapse:collapse; background-color:#dedede;"},
	}
	DefaultTableDataAttr = Attributes{
		{Name: "style", Value: "border:1px black solid; padding:3px 3px 3px 3px; border-collapse:collapse;"},
	}
)
