package table

import (
	"fmt"
	"os"
	"reflect"

	"github.com/olekukonko/tablewriter"
)

type Table struct {
	Headers []string
	Rows    [][]string
}

func Render(table Table) {
	tableWriter := tablewriter.NewWriter(os.Stdout)
	tableWriter.SetRowLine(true)
	tableWriter.SetHeader(table.Headers)

	for _, row := range table.Rows {
		tableWriter.Append(row)
	}

	fmt.Printf("\n")
	tableWriter.Render()
	fmt.Printf("\n")
}

func ExtractHeaders(obj interface{}) []string {
	var headers []string

	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		headers = append(headers, typ.Field(i).Name)
	}

	return headers
}

func GetFieldValue(obj interface{}, fieldName string) (interface{}, error) {
	val := reflect.ValueOf(obj)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	field := val.FieldByName(fieldName)
	if !field.IsValid() {
		return nil, fmt.Errorf("Field %s not found", fieldName)
	}

	return field.Interface(), nil
}
