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
		ID:              table.RenderString(server.ID),
		Hostname:        table.RenderString(attr.Hostname),
		PrimaryIPV4:     table.RenderString(*attr.PrimaryIPV4),
		Location:        table.RenderString(attr.Region.Site.Slug),
		Status:          table.RenderString(attr.Status),
		IPMIStatus:      table.RenderString(attr.IpmiStatus),
		Project:         table.RenderString(attr.Project.Slug),
		Team:            table.RenderString(attr.Team.Name),
		Plan:            table.RenderString(attr.Plan.Name),
		OperatingSystem: table.RenderString(attr.OperatingSystem.Slug),
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
