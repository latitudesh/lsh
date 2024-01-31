package tablerows

import (
	"github.com/latitudesh/lsh/internal/output/table"
	"github.com/latitudesh/lsh/models"
)

type VirtualNetwork struct {
	ID          string `json:"id,omitempty"`
	VID         string `json:"vid,omitempty"`
	Description string `json:"description,omitempty"`
	Assignments string `json:"assignments,omitempty"`
	Location    string `json:"location,omitempty"`
}

func NewVirtualNetwork(virtual_network *models.VirtualNetwork) *VirtualNetwork {
	attr := virtual_network.Attributes

	return &VirtualNetwork{
		ID:          table.RenderString(virtual_network.ID),
		VID:         table.RenderInt(attr.Vid),
		Description: table.RenderString(attr.Description),
		Assignments: table.RenderInt(attr.AssignmentsCount),
		Location:    table.RenderString(attr.Region.Site.Slug),
	}
}

func CreateVirtualNetworksRows(virtual_networks []*models.VirtualNetwork) []interface{} {
	var rows []VirtualNetwork
	var rowsInterface []interface{}

	for _, virtual_network := range virtual_networks {
		rows = append(rows, *NewVirtualNetwork(virtual_network))
	}

	for _, row := range rows {
		rowsInterface = append(rowsInterface, row)
	}

	return rowsInterface
}
