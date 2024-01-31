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

func NewVirtualNetwork(virtualNetwork *models.VirtualNetwork) *VirtualNetwork {
	attr := virtualNetwork.Attributes

	return &VirtualNetwork{
		ID:          table.RenderString(virtualNetwork.ID),
		VID:         table.RenderInt(attr.Vid),
		Description: table.RenderString(attr.Description),
		Assignments: table.RenderInt(attr.AssignmentsCount),
		Location:    table.RenderString(attr.Region.Site.Slug),
	}
}

func CreateVirtualNetworksRows(virtualNetworks []*models.VirtualNetwork) []interface{} {
	var rows []VirtualNetwork
	var rowsInterface []interface{}

	for _, virtualNetwork := range virtualNetworks {
		rows = append(rows, *NewVirtualNetwork(virtualNetwork))
	}

	for _, row := range rows {
		rowsInterface = append(rowsInterface, row)
	}

	return rowsInterface
}
