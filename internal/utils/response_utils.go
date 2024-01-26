// utils package contains general utility functions.
package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"

	apierrors "github.com/latitudesh/lsh/internal/api/errors"
	"github.com/latitudesh/lsh/internal/formatter"
	"github.com/latitudesh/lsh/internal/output"
	"github.com/olekukonko/tablewriter"
)

type OutputFormat string

const (
	JSONFormat  OutputFormat = "json"
	TableFormat OutputFormat = "table"
)

// PrettifyJson formats a JSON response, enhancing its readability on the terminal.
func PrettifyJson(str string) error {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return err
	}

	fmt.Println(prettyJSON.String())
	return nil
}

func PrintError(respErr error) error {
	switch e := respErr.(type) {
	case *apierrors.BadRequest:
		output.PrintBadRequestError(e)
	case *apierrors.Unauthorized:
		output.PrintUnauthorizedError(e)
	case *apierrors.Forbidden:
		output.PrintForbiddenError(e)
	case *apierrors.NotFound:
		output.PrintNotFoundError(e)
	case *apierrors.NotAcceptable:
		output.PrintNotAcceptableError(e)
	case *apierrors.UnprocessableEntity:
		output.PrintUnprocessableEntityError(e)
	default:
		return errors.New("An unknown error occurred.")
	}

	return nil
}

func PrintResult(jsonData string, format OutputFormat) error {
	// TODO: create a better feedback for empty responses, to let the users
	// know which action was executed and whether it failed or not.
	if jsonData == "" {
		fmt.Println("\nAction has been executed successfully!")
		return nil
	}

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return err
	}

	switch format {
	case JSONFormat:
		PrettifyJson(jsonData)
		return nil
	case TableFormat:
		printTableResult(data)
		return nil
	default:
		return errors.New("unsupported format")
	}
}

func printTableResult(data map[string]interface{}) {
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

	table.Render()
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
		if reflect.ValueOf(value).Kind() != reflect.Map && !ignoredAttributes[key] {
			headersMap[key] = true
		}
	}

	return headersMap
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
