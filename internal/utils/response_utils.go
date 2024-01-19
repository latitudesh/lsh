// utils package contains general utility functions.
package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ErrorResponse struct {
	Errors []ErrorDetail `json:"errors"`
}

type ErrorDetail struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
	Meta   struct {
		Attribute string `json:"attribute"`
		Message   string `json:"message"`
	} `json:"meta"`
	Status string `json:"status"`
	Title  string `json:"title"`
}

// PrettifyJson formats a JSON response, enhancing its readability on the terminal.
func PrettifyJson(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return str, err
	}

	return prettyJSON.String(), nil
}

func PrintOutput(jsonData string) {
	if jsonData == "" {
		fmt.Println("\nResource deleted successfully!")
		return
	}

	isError, err := isErrorResponse(jsonData)
	if err != nil {
		fmt.Println("Error when parsing the response data", err)
		return
	}

	if isError {
		// Unmarshal the JSON into the ErrorResponse structure
		var errorResponse ErrorResponse
		if err := json.Unmarshal([]byte(jsonData), &errorResponse); err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}

		// Format and display the errors
		printError(errorResponse.Errors)
	} else {
		fmt.Println(PrettifyJson(jsonData))
	}
}

func printError(errors []ErrorDetail) {
	refError := errors[0]

	fmt.Printf("\n")

	if refError.Code == "not_found" {
		fmt.Printf("%s\n", refError.Detail)
	} else if refError.Code == "VALIDATION_ERROR" {
		fmt.Println("The following errors have been found:")

		for _, err := range errors {
			fmt.Printf("     • '%s' %s\n", err.Meta.Attribute, err.Meta.Message)
		}
	} else {
		fmt.Println(errors)
	}

	fmt.Printf("\n")
}

func isErrorResponse(jsonString string) (bool, error) {
	// Check if the JSON string contains the "errors" key at the root level
	var check map[string]interface{}
	if err := json.Unmarshal([]byte(jsonString), &check); err != nil {
		return false, err
	}

	_, isError := check["errors"]
	return isError, nil
}
