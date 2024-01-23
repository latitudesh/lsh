package api

import (
	"errors"
	"fmt"

	"github.com/go-openapi/swag"
	apierrors "github.com/latitudesh/cli/internal/api/errors"
)

func RenderErrorOutput(respErr error) error {
	switch e := respErr.(type) {
	case *BadRequest:
		parseBadRequestError(e)
	case *Forbidden:
		parseForbiddenError(e)
	case *NotFound:
		parseNotFoundError(e)
	case *NotAcceptable:
		parseNotAcceptableError(e)
	case *UnprocessableEntity:
		parseUnprocessableEntityError(e)
	default:
		return errors.New("An unknown error occurred.")
	}

	return nil
}

func parseUnprocessableEntityError(respErr *UnprocessableEntity) error {
	if swag.IsZero(respErr) || swag.IsZero(respErr.Payload) {
		return nil
	}

	fmt.Printf("\n The following errors have been found: \n")

	for _, err := range respErr.GetPayload().Errors {
		if err.Meta.Attribute == "" {
			renderGenericError(&err)
		} else {
			fmt.Printf("     • '%s' %s\n", err.Meta.Attribute, err.Meta.Message)
		}
	}

	fmt.Printf("\n")

	return nil
}

func parseNotFoundError(respErr *NotFound) error {
	if swag.IsZero(respErr) || swag.IsZero(respErr.Payload) {
		return nil
	}

	refError := respErr.Payload.Errors[0]

	fmt.Printf("\n")
	fmt.Printf("     • Error: %s\n", refError.Detail)
	fmt.Printf("\n")

	return nil
}

func parseForbiddenError(respErr *Forbidden) error {
	if swag.IsZero(respErr) || swag.IsZero(respErr.Payload) {
		return nil
	}

	refError := respErr.Payload.Errors[0]

	fmt.Printf("\n")
	fmt.Printf("     • Error: %s\n", refError.Detail)
	fmt.Printf("\n")

	return nil
}

func parseBadRequestError(respErr *BadRequest) error {
	if swag.IsZero(respErr) || swag.IsZero(respErr.Payload) {
		return nil
	}

	renderGenericError(&respErr.Payload.Errors[0])

	return nil
}

func parseNotAcceptableError(respErr *NotAcceptable) error {
	if swag.IsZero(respErr) || swag.IsZero(respErr.Payload) {
		return nil
	}

	renderGenericError(&respErr.Payload.Errors[0])

	return nil
}

func renderGenericError(respErr *apierrors.ErrorDetail) {
	fmt.Printf("\n %s: \n", respErr.Title)
	fmt.Printf("     • %s\n", respErr.Detail)
	fmt.Printf("\n")
}
