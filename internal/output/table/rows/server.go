package tablerows

import (
	"github.com/latitudesh/lsh/internal/output/table"
	"github.com/latitudesh/lsh/models"
)

type ServerRow struct {
	ID              string `json:"id,omitempty"`
	Hostname        string `json:"hostname,omitempty"`
	PrimaryIPV4     string `json:"ipv4,omitempty"`
	Status          string `json:"status,omitempty"`
	IPMIStatus      string `json:"ipmi_status,omitempty"`
	Location        string `json:"location,omitempty"`
	Project         string `json:"project,omitempty"`
	Plan            string `json:"plan,omitempty"`
	OperatingSystem string `json:"os,omitempty"`
}

func NewServerRow(server *models.ServerData) *ServerRow {
	attr := server.Attributes

	return &ServerRow{
		ID:              table.RenderString(server.ID),
		Hostname:        table.RenderString(attr.Hostname),
		PrimaryIPV4:     table.RenderString(*attr.PrimaryIPV4),
		Location:        table.RenderString(location(attr)),
		Status:          table.RenderString(attr.Status),
		IPMIStatus:      table.RenderString(attr.IpmiStatus),
		Project:         table.RenderString(project(attr)),
		Plan:            table.RenderString(plan(attr)),
		OperatingSystem: table.RenderString(table.Truncate(operatingSystem(attr), 10)),
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

func location(attributes *models.ServerDataAttributes) string {
	region := attributes.Region

	if region != nil && region.Site != nil {
		return region.Site.Slug
	}

	return ""
}

func project(attributes *models.ServerDataAttributes) string {
	project := attributes.Project

	if project != nil {
		return project.Slug
	}

	return ""
}

func operatingSystem(attributes *models.ServerDataAttributes) string {
	operatingSystem := attributes.OperatingSystem

	if operatingSystem != nil {
		return operatingSystem.Slug
	}

	return ""
}

func plan(attributes *models.ServerDataAttributes) string {
	plan := attributes.Plan

	if plan != nil {
		return plan.Name
	}

	return ""
}
