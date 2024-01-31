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
		ID:             table.RenderString(project.ID),
		Name:           table.RenderString(attr.Name),
		Description:    table.RenderString(attr.Description),
		Slug:           table.RenderString(attr.Slug),
		IPs:            table.RenderFloat(attr.Stats.IPAddresses),
		Servers:        table.RenderFloat(attr.Stats.Servers),
		Vlans:          table.RenderFloat(attr.Stats.Vlans),
		Team:           table.RenderString(attr.Team.Name),
		Environment:    table.RenderString(attr.Environment),
		ProvisiongType: table.RenderString(attr.ProvisiongType),
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
