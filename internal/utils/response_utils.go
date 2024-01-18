// utils package contains general utility functions.
package utils

import (
	"bytes"
	"encoding/json"

	"github.com/latitudesh/cli/internal/formatters"
)

// PrettifyJson formats a JSON response, enhancing its readability on the terminal.
func PrettifyJson(str string) (string, error) {
	// TODO: validate if the request was a DELETE before returning this message
	if str == "" {
		return "Resource deleted successfully!", nil
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return str, err
	}

	return prettyJSON.String(), nil
}

func FormatErrorOutput(inputString string) string {
	return formatters.FormatError(inputString)
}
