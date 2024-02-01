package tablerows

import (
	"github.com/latitudesh/lsh/internal/output/table"
	"github.com/latitudesh/lsh/models"
)

func NewVirtualNetworkAssignmentRow(assignment *models.VirtualNetworkAssignment) table.Row {
	attr := assignment.Attributes

	return table.Row{
		"id": table.Cell{
			Label: "ID",
			Value: table.String(assignment.ID),
		},
		"vid": table.Cell{
			Label: "VID",
			Value: table.Int(attr.Vid),
		},
		"virtual_network_id": table.Cell{
			Label: "Virtual Network ID",
			Value: table.String(attr.VirtualNetworkID),
		},
		"description": table.Cell{
			Label: "Description",
			Value: table.String(attr.Description),
		},
		"status": table.Cell{
			Label: "Status",
			Value: table.String(attr.Status),
		},
		"server_id": table.Cell{
			Label: "Server ID",
			Value: table.String(attr.Server.ID),
		},
	}
}
