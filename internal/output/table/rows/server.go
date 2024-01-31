package tablerows

import (
	"github.com/latitudesh/lsh/internal/output/table"
	"github.com/latitudesh/lsh/models"
)

type ServerRow struct {
	ID              string `json:"id,omitempty"`
	Hostname        string `json:"hostname,omitempty"`
	PrimaryIPV4     string `json:"primary_ipv4,omitempty"`
	Status          string `json:"status,omitempty"`
	IPMIStatus      string `json:"ipmi_status,omitempty"`
	Location        string `json:"location,omitempty"`
	Project         string `json:"project,omitempty"`
	Team            string `json:"team,omitempty"`
	Plan            string `json:"plan,omitempty"`
	OperatingSystem string `json:"operating_system,omitempty"`
}

func NewServerRow(server *models.ServerData) *ServerRow {
	attr := server.Attributes

	return &ServerRow{
		ID:              table.String(server.ID),
		Hostname:        table.String(attr.Hostname),
		PrimaryIPV4:     table.String(*attr.PrimaryIPV4),
		Location:        table.String(attr.Region.Site.Slug),
		Status:          table.String(attr.Status),
		IPMIStatus:      table.String(attr.IpmiStatus),
		Project:         table.String(attr.Project.Slug),
		Team:            table.String(attr.Team.Name),
		Plan:            table.String(attr.Plan.Name),
		OperatingSystem: table.String(attr.OperatingSystem.Slug),
	}
}

func CreateServerRows(servers []*models.ServerData) []interface{} {
	var rows []ServerRow
	var rowsInterface []interface{}

	for _, server := range servers {
		rows = append(rows, *NewServerRow(server))
	}

	for _, row := range rows {
		rowsInterface = append(rowsInterface, row)
	}

	return rowsInterface
}
