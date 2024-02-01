package table

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

const MaxColWidth = 20
const MaxLength = 50

type Table struct {
	Headers []string
	Rows    [][]string
}

type Header struct {
	IDs    []string
	Labels []string
}

func Render(rows []Row) {
	headers := extractHeaders(rows[0])

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

func extractHeaders(row Row) Header {
	var headers Header

	for k, v := range row {
		headers.IDs = append(headers.IDs, k)
		headers.Labels = append(headers.Labels, v.Label)
	}

	return headers
}

func renderCell(cell Cell) string {
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
