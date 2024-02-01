package tablerows

import (
	"github.com/latitudesh/lsh/internal/output/table"
	"github.com/latitudesh/lsh/models"
)

func NewServerRow(server *models.ServerData) table.Row {
	attr := server.Attributes

	return table.Row{
		"id": table.Cell{
			Label: "ID",
			Value: table.String(server.ID),
		},
		"hostname": table.Cell{
			Label:     "Hostname",
			Value:     table.String(attr.Hostname),
			MaxLength: 15,
		},
		"primary_ipv4": table.Cell{
			Label: "Primary IPV4",
			Value: table.String(*attr.PrimaryIPV4),
		},
		"location": table.Cell{
			Label: "Location",
			Value: table.String(attr.Region.Site.Slug),
		},
		"status": table.Cell{
			Label: "Status",
			Value: table.String(attr.Status),
		},
		"ipmi_status": table.Cell{
			Label: "IPMI Status",
			Value: table.String(attr.IpmiStatus),
		},
		"project": table.Cell{
			Label: "Project",
			Value: table.String(attr.Project.Slug),
		},
		"team": table.Cell{
			Label: "Team",
			Value: table.String(attr.Team.Name),
		},
		"plan": table.Cell{
			Label: "Plan",
			Value: table.String(attr.Plan.Name),
		},
		"operating_system": table.Cell{
			Label:     "OS",
			Value:     table.String(attr.OperatingSystem.Slug),
			MaxLength: 15,
		},
	}
}
