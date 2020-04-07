package emailt

import (
	"fmt"
	"io"
	"reflect"
	"sort"
)

type Column struct {
	Name     string
	Template string
}

type Columns []Column

type Table struct {
	Dataset interface{}
	Columns Columns
}

func (t Table) WithDataset(dataset interface{}) Table {
	t.Dataset = dataset
	return t
}

func (t Table) WithColumns(columns Columns) Table {
	t.Columns = columns
	return t
}

func (t Table) Render(writer io.Writer, themes ...Theme) error {
	theme := mergeThemes(themes)

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

	render := newFmtWriter(writer)

	render.Printlnf("<table %s>", theme.Attributes("table"))

	render.Println("<tr>")

	for _, column := range columns {
		render.Printlnf("<th %s>%s</th>", theme.Attributes("th"), column.Name)
	}
	render.Println("</tr>")

	for i := 0; i < dataset.Len(); i++ {
		render.Println("<tr>")
		for _, column := range columns {
			render.Printlnf("<td %s>", theme.Attributes("td"))
			e := TemplateElement{
				Data:     dataset.Index(i),
				Template: column.Template,
			}
			if err := e.Render(writer, theme); err != nil {
				return err
			}
			render.Println("\n</td>")
		}
		render.Println("</tr>")
	}

	render.Println("</table>")

	return render.Err()
}
