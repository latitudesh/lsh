package output

import (
	"fmt"

	"github.com/go-openapi/swag"
	apierrors "github.com/latitudesh/lsh/internal/api/errors"
)

func PrintUnprocessableEntityError(respErr *apierrors.UnprocessableEntity) error {
	if swag.IsZero(respErr) || swag.IsZero(respErr.Payload) {
		return nil
	}

	fmt.Printf("\n The following errors have been found: \n")

	for _, err := range respErr.GetPayload().Errors {
		if err.Meta.Attribute == "" {
			printGenericError(&err)
		} else {
			fmt.Printf("     • '%s' %s\n", err.Meta.Attribute, err.Meta.Message)
		}
	}

	fmt.Printf("\n")

	return nil
}

func PrintNotFoundError(respErr *apierrors.NotFound) error {
	if swag.IsZero(respErr) || swag.IsZero(respErr.Payload) {
		return nil
	}

	refError := respErr.Payload.Errors[0]

	fmt.Printf("\n")
	fmt.Printf("     • Error: %s\n", refError.Detail)
	fmt.Printf("\n")

	return nil
}

func PrintForbiddenError(respErr *apierrors.Forbidden) error {
	if swag.IsZero(respErr) || swag.IsZero(respErr.Payload) {
		return nil
	}

	refError := respErr.Payload.Errors[0]

	fmt.Printf("\n")
	fmt.Printf("     • Error: %s\n", refError.Detail)
	fmt.Printf("\n")

	return nil
}

func PrintBadRequestError(respErr *apierrors.BadRequest) error {
	if swag.IsZero(respErr) || swag.IsZero(respErr.Payload) {
		return nil
	}

	printGenericError(&respErr.Payload.Errors[0])

	return nil
}

func PrintNotAcceptableError(respErr *apierrors.NotAcceptable) error {
	if swag.IsZero(respErr) || swag.IsZero(respErr.Payload) {
		return nil
	}

	printGenericError(&respErr.Payload.Errors[0])

	return nil
}

func PrintUnauthorizedError(respErr *apierrors.Unauthorized) error {
	if swag.IsZero(respErr) || swag.IsZero(respErr.Payload) {
		return nil
	}

	fmt.Printf("\nUnauthorized request.\n")
	fmt.Printf("\nMake sure your API Key is set up by running:\n ")
	fmt.Printf("     • lsh login YOUR_API_KEY\n\n")

	return nil
}

func printGenericError(respErr *apierrors.ErrorDetail) {
	fmt.Printf("\n %s: \n", respErr.Title)
	fmt.Printf("     • %s\n", respErr.Detail)
	fmt.Printf("\n")
}
