package tablerows

import (
	"github.com/latitudesh/lsh/internal/output/table"
	"github.com/latitudesh/lsh/models"
)

func NewAPIKeyRow(APIKey *models.APIKey) table.Row {
	attr := APIKey.Attributes

	return table.Row{
		"id": table.Cell{
			Label: "ID",
			Value: table.String(APIKey.ID),
		},
		"name": table.Cell{
			Label: "Name",
			Value: table.String(attr.Name),
		},
		"token_last_slice": table.Cell{
			Label: "Token Last Slice",
			Value: table.String(attr.TokenLastSlice),
		},
		"user": table.Cell{
			Label: "User",
			Value: table.String(attr.User.Email),
		},
		"api_version": table.Cell{
			Label: "API Version",
			Value: table.String(attr.APIVersion),
		},
		"last_used_at": table.Cell{
			Label: "Last Used At",
			Value: table.DateTime(attr.LastUsedAt),
		},
	}
}
