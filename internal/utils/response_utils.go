// utils package contains general utility functions.
package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	apierrors "github.com/latitudesh/lsh/internal/api/errors"
	"github.com/latitudesh/lsh/internal/output"
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

	fmt.Println(PrettifyJson(jsonData))
}
