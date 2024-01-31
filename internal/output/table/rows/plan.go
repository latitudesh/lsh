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
		"id":           table.NewCell("ID", table.String(plan.ID)),
		"slug":         table.NewCell("Slug", table.String(attr.Slug)),
		"available_in": table.NewCell("Available In", table.StringList(regions.AvailableIn)),
		"in_stock":     table.NewCell("In Stock", table.StringList(regions.InStock)),
		"features":     table.NewCell("Features", table.StringList(attr.Features)),
		"cpu":          table.NewCell("CPU", table.PlanCPU(*attr.Specs.CPU)),
		"memory":       table.NewCell("Memory", table.PlanMemory(*attr.Specs.Memory)),
		"drives":       table.NewCell("Drives", table.PlanDrives(attr.Specs.Drives)),
		"nic":          table.NewCell("NIC", table.PlanNICs(attr.Specs.Nics)),
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
