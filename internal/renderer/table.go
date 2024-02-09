package renderer

import "github.com/latitudesh/lsh/internal/output/table"

type TableRenderer struct{}

func (tr TableRenderer) Render(data []ResponseData) {
	if len(data) == 0 {
		table.RenderEmptyState("No results found.")
		return
	}

	var rows []table.Row

	for _, resource := range data {
		rows = append(rows, resource.TableRow())
	}

	table.Render(rows)
}
