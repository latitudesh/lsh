package tablerows

import (
	"github.com/latitudesh/lsh/internal/output/table"
	"github.com/latitudesh/lsh/models"
)

func NewBandwidthPlanRow(plan *models.BandwidthPlan) table.Row {
	attr := plan.Attributes

	return table.Row{
		"id":              table.NewCell("ID", table.String(plan.ID)),
		"region":          table.NewCell("Region", table.String(attr.Region)),
		"locations":       table.NewCell("Locations", table.StringList(attr.Locations)),
		"usd_price_hour":  table.NewCell("USD/Hour", table.Int(attr.Pricing.Usd.Hourly)),
		"usd_price_month": table.NewCell("USD/Month", table.Int(attr.Pricing.Usd.Monthly)),
		"brl_price_hour":  table.NewCell("BRL/Hour", table.Int(attr.Pricing.Brl.Hourly)),
		"brl_price_month": table.NewCell("BRL/Month", table.Int(attr.Pricing.Brl.Monthly)),
	}
}
