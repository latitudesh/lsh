package tablerows

import (
	"github.com/latitudesh/lsh/internal/output/table"
	"github.com/latitudesh/lsh/models"
)

type PlanRegions struct {
	InStock     []string
	AvailableIn []string
}

func NewPlanRow(plan *models.PlanData) table.Row {
	attr := plan.Attributes

	regions := parseRegions(attr.Regions)

	return table.Row{
		"id": table.Cell{
			Label: "ID",
			Value: table.String(plan.ID),
		},
		"slug": table.Cell{
			Label: "Slug",
			Value: table.String(attr.Slug),
		},
		"available_in": table.Cell{
			Label: "Available In",
			Value: table.StringList(regions.AvailableIn),
		},
		"in_stock": table.Cell{
			Label: "In Stock",
			Value: table.StringList(regions.InStock),
		},
		"features": table.Cell{
			Label: "Features",
			Value: table.StringList(attr.Features),
		},
		"cpu": table.Cell{
			Label: "CPU",
			Value: table.PlanCPU(*attr.Specs.CPU),
		},
		"memory": table.Cell{
			Label: "Memory",
			Value: table.PlanMemory(*attr.Specs.Memory),
		},
		"drives": table.Cell{
			Label: "Drives",
			Value: table.PlanDrives(attr.Specs.Drives),
		},
		"nic": table.Cell{
			Label: "NIC",
			Value: table.PlanNICs(attr.Specs.Nics),
		},
	}
}

func parseRegions(regions []*models.PlanDataAttributesRegionsItems0) PlanRegions {
	var planRegions PlanRegions

	for _, region := range regions {
		planRegions.InStock = append(planRegions.InStock, region.Locations.InStock...)
		planRegions.AvailableIn = append(planRegions.AvailableIn, region.Locations.Available...)
	}

	return planRegions
}
