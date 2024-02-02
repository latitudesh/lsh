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
			Value: table.String(location(attr)),
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
			Value: table.String(project(attr)),
		},
		"plan": table.Cell{
			Label: "Plan",
			Value: table.String(plan(attr)),
		},
		"operating_system": table.Cell{
			Label:     "OS",
			Value:     table.String(attr.OperatingSystem.Slug),
			MaxLength: 10,
		},
	}
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
