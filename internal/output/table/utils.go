package table

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/olekukonko/tablewriter"
)

const MaxColWidth = 20
const MaxLength = 50

type Table struct {
	Headers []string
	Rows    [][]string
}

func Render(rows []interface{}) {
	headers := extractHeaders(rows[0])
	fieldNames := extractFieldNames(rows[0])

	var values [][]string

	for _, row := range rows {
		var tr []string
		val := reflect.ValueOf(row)

		for _, fieldKey := range fieldNames {
			field := val.FieldByName(fieldKey)
			if !field.IsValid() {
				fmt.Printf("Field not found: %s\n", fieldKey)
				continue
			}
			tr = append(tr, truncate(fmt.Sprintf("%s", field.Interface())))
		}

		values = append(values, tr)
	}

	render(Table{Headers: headers, Rows: values})
}

func render(table Table) {
	tableWriter := tablewriter.NewWriter(os.Stdout)

	// tableWriter.SetAutoWrapText(false)
	tableWriter.SetRowLine(true)
	tableWriter.SetHeader(table.Headers)
	tableWriter.SetColWidth(MaxColWidth)
	for _, row := range table.Rows {
		tableWriter.Append(row)
	}

	fmt.Printf("\n")
	tableWriter.Render()
	fmt.Printf("\n")
}

func extractHeaders(obj interface{}) []string {
	var headers []string

	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		jsonTag := field.Tag.Get("json")
		jsonTagWithoutOmitempty := removeOmitempty(jsonTag)
		headers = append(headers, jsonTagWithoutOmitempty)
	}

	return headers
}

func extractFieldNames(obj interface{}) []string {
	var fieldNames []string

	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldName := field.Name
		fieldNames = append(fieldNames, fieldName)
	}

	return fieldNames
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

func removeOmitempty(tag string) string {
	return strings.Split(tag, ",")[0]
}

func truncate(input string) string {
	if len(input) > MaxLength {
		return input[:MaxLength] + "..."
	}
	return input
}
