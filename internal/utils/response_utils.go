// utils package contains general utility functions.
package utils

import (
	"bytes"
	"encoding/json"
)

// PrettifyJson formats a JSON response, enhancing its readability on the terminal.
func PrettifyJson(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return str, err
	}

	return prettyJSON.String(), nil
}
