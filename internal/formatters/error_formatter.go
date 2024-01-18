package formatters

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/latitudesh/cli/internal"
)

// FormatError formats the error details as per the expected format.
func FormatError(errorString string) string {
	// Extract the JSON part from the input
	jsonPart, err := extractJSONFromError(errorString)
	if err != nil {
		return errorString
	}

	// Unmarshal the JSON into the ErrorResponse struct
	var errorResponse internal.ErrorResponse
	err = json.Unmarshal([]byte(jsonPart), &errorResponse)
	if err != nil {
		return errorString
	}

	if len(errorResponse.Errors) == 0 {
		return ""
	}

	errorDetail := errorResponse.Errors[0]

	if errorDetail.Source.Pointer == "filter" {
		return "'--filter' flag is missing."
	}

	// Construct the formatted error message
	formattedError := fmt.Sprintf("%s: '%s' %s.", errorDetail.Title, errorDetail.Source.Pointer, errorDetail.Detail)

	return formattedError
}

// ExtractJSONFromBody extracts the JSON part from the given input.
func extractJSONFromError(input string) (string, error) {
	// Define a regular expression to match JSON
	jsonRegex := regexp.MustCompile(`\{.*\}`)

	// Find the first match in the input string
	match := jsonRegex.FindString(input)

	if match == "" {
		return "", fmt.Errorf("JSON not found in the input string")
	}

	return match, nil
}
