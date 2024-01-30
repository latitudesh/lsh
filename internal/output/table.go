package output

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/latitudesh/lsh/internal/formatter"
	"github.com/olekukonko/tablewriter"
)

func RenderTable(str string) error {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)

	entries := extractEntries(data)

	var headers []string
	headers = extractHeaders(entries[0])
	table.SetHeader(headers)

	for _, entry := range entries {
		row := extractRow(headers, entry.(map[string]interface{}))
		table.Append(row)
	}

	fmt.Printf("\n")
	table.Render()
	fmt.Printf("\n")

	return nil
}

func extractEntries(data map[string]interface{}) []interface{} {
	var entries []interface{}

	// Determine whether the data is a list of resources or a single resource
	switch data := data["data"].(type) {
	case []interface{}:
		entries = data
	case interface{}:
		entries = append(entries, data)
	}

	return entries
}

func extractHeaders(entry interface{}) []string {
	headersMap := extractHeadersFromAttributes(entry)

	headers := make([]string, 0, len(headersMap))
	for header := range headersMap {
		headers = append(headers, header)
	}

	return append([]string{"ID"}, headers...)
}

func extractHeadersFromAttributes(entry interface{}) map[string]bool {
	headersMap := make(map[string]bool)

	ignoredAttributes := map[string]bool{"created_at": true, "updated_at": true}

	for key, value := range entry.(map[string]interface{})["attributes"].(map[string]interface{}) {
		if !containsNestedValues(value) && !ignoredAttributes[key] {
			headersMap[key] = true
		}
	}

	return headersMap
}

func containsNestedValues(value interface{}) bool {
	kind := reflect.ValueOf(value).Kind()
	return kind == reflect.Map || kind == reflect.Slice
}

func extractRow(headers []string, entry map[string]interface{}) []string {
	row := make([]string, len(headers))

	row[0] = fmt.Sprint(entry["id"])

	for i, header := range headers[1:] {
		attr, ok := entry["attributes"].(map[string]interface{})
		if !ok {
			continue
		}

		value, exists := attr[header]
		if !exists {
			continue
		}

		row[i+1] = formatter.FormatValue(value)
	}

	return row
}
