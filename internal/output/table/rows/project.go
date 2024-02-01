package tablerows

import (
	"github.com/latitudesh/lsh/internal/output/table"
	"github.com/latitudesh/lsh/models"
)

func NewProjectRow(project *models.Project) table.Row {
	attr := project.Attributes

	return table.Row{
		"id": table.Cell{
			Label: "ID",
			Value: table.String(project.ID),
		},
		"name": table.Cell{
			Label: "Name",
			Value: table.String(attr.Name),
		},
		"description": table.Cell{
			Label: "Description",
			Value: table.String(attr.Description),
		},
		"slug": table.Cell{
			Label: "Slug",
			Value: table.String(attr.Slug),
		},
		"ips": table.Cell{
			Label: "IPs",
			Value: table.Float(attr.Stats.IPAddresses),
		},
		"servers": table.Cell{
			Label: "Servers",
			Value: table.Float(attr.Stats.Servers),
		},
		"vlans": table.Cell{
			Label: "VLANs",
			Value: table.Float(attr.Stats.Vlans),
		},
		"team": table.Cell{
			Label: "Team",
			Value: table.String(attr.Team.Name),
		},
		"environment": table.Cell{
			Label: "Plan",
			Value: table.String(attr.Environment),
		},
		"provisioning_type": table.Cell{
			Label: "Provisioning Type",
			Value: table.String(attr.ProvisiongType),
		},
	}
}
