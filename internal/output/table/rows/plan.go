package tablerows

import (
	"github.com/latitudesh/lsh/internal/output/table"
	"github.com/latitudesh/lsh/models"
)

type PlanRow struct {
	ID          string `json:"id,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Features    string `json:"features,omitempty"`
	InStock     string `json:"in_stock,omitempty"`
	AvailableIn string `json:"available_in,omitempty"`
	CPU         string `json:"cpu,omitempty"`
	Memory      string `json:"memory,omitempty"`
	Drives      string `json:"drives,omitempty"`
	NIC         string `json:"nic,omitempty"`
}

func NewPlanRow(plan *models.PlanData) *PlanRow {
	attr := plan.Attributes

	var inStock []string
	var availableIn []string

	for _, region := range attr.Regions {
		inStock = append(inStock, region.Locations.InStock...)
		availableIn = append(availableIn, region.Locations.Available...)
	}

	return &PlanRow{
		ID:          table.RenderString(plan.ID),
		Slug:        table.RenderString(attr.Slug),
		AvailableIn: table.RenderStringList(availableIn, ", "),
		InStock:     table.RenderStringList(inStock, "\n"),
		Features:    table.RenderStringList(attr.Features, "\n\n"),
		CPU:         table.RenderPlanCPU(*attr.Specs.CPU),
		Memory:      table.RenderPlanMemory(*attr.Specs.Memory),
		Drives:      table.RenderPlanDrives(attr.Specs.Drives),
		NIC:         table.RenderPlanNICs(attr.Specs.Nics),
	}
}

func CreatePlanRows(plans []*models.PlanData) []interface{} {
	var rows []PlanRow
	var rowsInterface []interface{}

	for _, plan := range plans {
		rows = append(rows, *NewPlanRow(plan))
	}

	for _, row := range rows {
		rowsInterface = append(rowsInterface, row)
	}

	return rowsInterface
}
