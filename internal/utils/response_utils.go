// utils package contains general utility functions.
package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// PrettifyJson formats a JSON response, enhancing its readability on the terminal.
func PrettifyJson(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return str, err
	}

	return prettyJSON.String(), nil
}

func PrintOutput(jsonData string) {
	// TODO: create a better feedback for empty responses, to let the users
	// know which action was executed and whether it failed or not.
	if jsonData == "" {
		fmt.Println("\nAction has been executed successfully!")
		return
	}

	fmt.Println(PrettifyJson(jsonData))
}
