package tablerows

import (
	"github.com/latitudesh/lsh/internal/output/table"
	"github.com/latitudesh/lsh/models"
)

type ProjectRow struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	Slug           string `json:"slug,omitempty"`
	IPs            string `json:"ips,omitempty"`
	Servers        string `json:"servers,omitempty"`
	Vlans          string `json:"vlans,omitempty"`
	Team           string `json:"team,omitempty"`
	Environment    string `json:"environment,omitempty"`
	ProvisiongType string `json:"provisioning_type,omitempty"`
}

func NewProjectRow(project *models.Project) *ProjectRow {
	attr := project.Attributes

	return &ProjectRow{
		ID:             table.String(project.ID),
		Name:           table.String(attr.Name),
		Description:    table.String(attr.Description),
		Slug:           table.String(attr.Slug),
		IPs:            table.Float(attr.Stats.IPAddresses),
		Servers:        table.Float(attr.Stats.Servers),
		Vlans:          table.Float(attr.Stats.Vlans),
		Team:           table.String(attr.Team.Name),
		Environment:    table.String(attr.Environment),
		ProvisiongType: table.String(attr.ProvisiongType),
	}
}

func CreateProjectRows(projects []*models.Project) []interface{} {
	var rows []ProjectRow
	var rowsInterface []interface{}

	for _, Project := range projects {
		rows = append(rows, *NewProjectRow(Project))
	}

	for _, row := range rows {
		rowsInterface = append(rowsInterface, row)
	}

	return rowsInterface
}
