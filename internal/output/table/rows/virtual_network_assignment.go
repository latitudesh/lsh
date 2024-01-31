package tablerows

import (
	"github.com/latitudesh/lsh/internal/output/table"
	"github.com/latitudesh/lsh/models"
)

type VirtualNetworkAssignment struct {
	ID               string `json:"id,omitempty"`
	Vid              string `json:"vid,omitempty"`
	VirtualNetworkID string `json:"assignment_id,omitempty"`
	Description      string `json:"description,omitempty"`
	Status           string `json:"status,omitempty"`
	Server           string `json:"server,omitempty"`
}

func NewVirtualNetworkAssignment(assignment *models.VirtualNetworkAssignment) *VirtualNetworkAssignment {
	attr := assignment.Attributes

	return &VirtualNetworkAssignment{
		ID:               table.String(assignment.ID),
		Vid:              table.Int(attr.Vid),
		VirtualNetworkID: table.String(attr.VirtualNetworkID),
		Description:      table.String(attr.Description),
		Status:           table.String(attr.Status),
		Server:           table.String(attr.Server.ID),
	}
}

func CreateVirtualNetworkAssignmentsRows(assignments []*models.VirtualNetworkAssignment) []interface{} {
	var rows []VirtualNetworkAssignment
	var rowsInterface []interface{}

	for _, assignment := range assignments {
		rows = append(rows, *NewVirtualNetworkAssignment(assignment))
	}

	for _, row := range rows {
		rowsInterface = append(rowsInterface, row)
	}

	return rowsInterface
}
