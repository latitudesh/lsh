package tablerows

import (
	"github.com/latitudesh/lsh/internal/output/table"
	"github.com/latitudesh/lsh/models"
)

func NewVirtualNetworkRow(virtualNetwork *models.VirtualNetwork) table.Row {
	attr := virtualNetwork.Attributes

	return table.Row{
		"id": table.Cell{
			Label: "ID",
			Value: table.String(virtualNetwork.ID),
		},
		"vid": table.Cell{
			Label: "VID",
			Value: table.Int(attr.Vid),
		},
		"description": table.Cell{
			Label: "Description",
			Value: table.String(attr.Description),
		},
		"assignments": table.Cell{
			Label: "Assignments",
			Value: table.Int(attr.AssignmentsCount),
		},
		"location": table.Cell{
			Label: "Location",
			Value: table.String(virtualNetworkLocation(attr)),
		},
	}
}

func virtualNetworkLocation(attributes *models.VirtualNetworkAttributes) string {
	region := attributes.Region

	if region != nil && region.Site != nil {
		return region.Site.Slug
	}

	return ""
}
