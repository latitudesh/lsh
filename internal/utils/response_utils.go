// utils package contains general utility functions.
package utils

import (
	"errors"
	"fmt"

	apierrors "github.com/latitudesh/lsh/internal/api/errors"
	"github.com/latitudesh/lsh/internal/output"
	"github.com/spf13/viper"
)

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

func PrintResult(str string) {
	// TODO: create a better feedback for empty responses, to let the users
	// know which action was executed and whether it failed or not.
	if str == "" {
		fmt.Println("\nAction has been executed successfully!")
	}

	formatAsJSON := viper.GetBool("json")

	if formatAsJSON {
		output.RenderJSON(str)
		return
	}

	formatOutputFlag := viper.GetString("output")

	switch formatOutputFlag {
	case "json":
		output.RenderJSON(str)
	case "table":
		output.RenderTable(str)
	default:
		fmt.Println("Unsupported output format")
	}
}
