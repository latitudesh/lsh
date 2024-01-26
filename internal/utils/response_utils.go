// utils package contains general utility functions.
package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"time"

	apierrors "github.com/latitudesh/lsh/internal/api/errors"
	"github.com/latitudesh/lsh/internal/output"
	"github.com/olekukonko/tablewriter"
)

// PrettifyJson formats a JSON response, enhancing its readability on the terminal.
func PrettifyJson(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return str, err
	}

	return prettyJSON.String(), nil
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

func PrintResult(jsonData string) {
	// TODO: create a better feedback for empty responses, to let the users
	// know which action was executed and whether it failed or not.
	if jsonData == "" {
		fmt.Println("\nAction has been executed successfully!")
		return
	}

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	printTableResult(data)
}

func printTableResult(data map[string]interface{}) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)

	var headers []string

	entries, ok := data["data"].([]interface{})
	if ok {
		headers = extractHeader(entries[0])
		table.SetHeader(headers)

		for _, entry := range entries {
			row := extractRow(headers, entry.(map[string]interface{}))
			table.Append(row)
		}
	}

	entry, ok := data["data"].(interface{})
	if ok {
		headers = extractHeader(entry)
		table.SetHeader(headers)

		row := extractRow(headers, entry.(map[string]interface{}))
		table.Append(row)
	}

	table.Render()
}

func extractHeader(entry interface{}) []string {
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
			// Ignore headers without attributes
			continue
		}

		value, exists := attr[header]
		if !exists {
			// Ignore missing attribute values
			continue
		}

		// Check if the value is a date and format it
		if isDate(value) {
			dateStr, ok := value.(string)
			if ok {
				// Attempt to parse the date
				date, err := time.Parse(time.RFC3339, dateStr)
				if err == nil {
					// Use the parsed date if successful
					row[i+1] = date.Format("2006-01-02 15:04:05")
					continue
				}
			}
		}

		row[i+1] = truncate(fmt.Sprint(value), 50)
	}

	return row
}

// isDate checks if a value is likely a date (assumes string for simplicity).
func isDate(value interface{}) bool {
	_, ok := value.(string)
	return ok
}

func truncate(input string, maxLength int) string {
	if len(input) > maxLength {
		return input[:maxLength] + "..."
	}
	return input
}
