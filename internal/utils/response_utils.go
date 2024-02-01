// utils package contains general utility functions.
package utils

import (
	"errors"
	"fmt"
	"os"

	apierrors "github.com/latitudesh/lsh/internal/api/errors"
	"github.com/latitudesh/lsh/internal/output"
	"github.com/latitudesh/lsh/internal/output/table"
	"github.com/olekukonko/tablewriter"
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

func RenderTableU(rows []table.Row) {
	headers := ExtractHeadersU(rows[0])

	tableWriter := tablewriter.NewWriter(os.Stdout)

	tableWriter.SetRowLine(true)
	tableWriter.SetHeader(headers.Labels)
	tableWriter.SetColWidth(20)

	for _, row := range rows {
		var tr []string

		for _, attr := range headers.IDs {
			tr = append(tr, renderCell(row[attr]))
		}

		tableWriter.Append(tr)
	}

	fmt.Printf("\n")
	tableWriter.Render()
	fmt.Printf("\n")
}

type Header struct {
	IDs    []string
	Labels []string
}

func ExtractHeadersU(row table.Row) Header {
	var headers Header

	for k, v := range row {
		headers.IDs = append(headers.IDs, k)
		headers.Labels = append(headers.Labels, v.Label)
	}

	return headers
}

func renderCell(cell table.Cell) string {
	maxLength := 50

	if cell.MaxLength > 0 {
		maxLength = cell.MaxLength
	}

	return fmt.Sprintf("%v", truncate(cell.Value, maxLength))
}

func truncate(input string, maxLength int) string {
	if len(input) > maxLength {
		return input[:maxLength] + "..."
	}
	return input
}
