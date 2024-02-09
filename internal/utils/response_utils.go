// utils package contains general utility functions.
package utils

import (
	"errors"
	"fmt"

	apierrors "github.com/latitudesh/lsh/internal/api/errors"
	"github.com/latitudesh/lsh/internal/output"
	"github.com/latitudesh/lsh/internal/renderer"
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

func Render(data []renderer.ResponseData) {
	var r renderer.Renderer

	formatAsJSON := viper.GetBool("json")

	if formatAsJSON {
		r = renderer.JSONRenderer{}
	}

	formatOutputFlag := viper.GetString("output")

	switch formatOutputFlag {
	case "json":
		r = renderer.JSONRenderer{}
	case "table":
		r = renderer.TableRenderer{}
	default:
		fmt.Println("Unsupported output format")
	}

	renderContext := renderer.NewRenderContext(r)

	renderContext.Render(data)
}
