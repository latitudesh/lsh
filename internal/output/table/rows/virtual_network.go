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
		ID:          table.String(virtualNetwork.ID),
		VID:         table.Int(attr.Vid),
		Description: table.String(attr.Description),
		Assignments: table.Int(attr.AssignmentsCount),
		Location:    table.String(attr.Region.Site.Slug),
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
