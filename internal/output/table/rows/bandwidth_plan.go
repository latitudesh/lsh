package tablerows

import (
	"github.com/latitudesh/lsh/internal/output/table"
	"github.com/latitudesh/lsh/models"
)

func NewBandwidthPlanRow(plan *models.BandwidthPlan) table.Row {
	attr := plan.Attributes

	return table.Row{
		"id": table.Cell{
			Label: "ID",
			Value: table.String(plan.ID),
		},
		"region": table.Cell{
			Label: "Region",
			Value: table.String(attr.Region),
		},
		"locations": table.Cell{
			Label: "Locations",
			Value: table.StringList(attr.Locations),
		},
		"usd_price_hour": table.Cell{
			Label: "USD/Hour",
			Value: table.Int(attr.Pricing.Usd.Hourly),
		},
		"usd_price_month": table.Cell{
			Label: "USD/Month",
			Value: table.Int(attr.Pricing.Usd.Monthly),
		},
		"brl_price_hour": table.Cell{
			Label: "BRL/Hour",
			Value: table.Int(attr.Pricing.Brl.Hourly),
		},
		"brl_price_month": table.Cell{
			Label: "BRL/Month",
			Value: table.Int(attr.Pricing.Brl.Monthly),
		},
	}
}
