package tablerows

import (
	"github.com/latitudesh/lsh/internal/output/table"
	"github.com/latitudesh/lsh/models"
)

type BandwidthPlanRow struct {
	ID            string `json:"id,omitempty"`
	Region        string `json:"region,omitempty"`
	Locations     string `json:"locations,omitempty"`
	USDPriceHour  string `json:"usd_price_hour,omitempty"`
	USDPriceMonth string `json:"usd_price_month,omitempty"`
	BRLPriceHour  string `json:"brl_price_hour,omitempty"`
	BRLPriceMonth string `json:"brl_price_month,omitempty"`
}

func NewBandwidthPlanRow(plan *models.BandwidthPlan) *BandwidthPlanRow {
	attr := plan.Attributes

	return &BandwidthPlanRow{
		ID:            table.RenderString(plan.ID),
		Region:        table.RenderString(attr.Region),
		Locations:     table.RenderStringList(attr.Locations, ", "),
		USDPriceHour:  table.RenderInt(attr.Pricing.Usd.Hourly),
		USDPriceMonth: table.RenderInt(attr.Pricing.Usd.Monthly),
		BRLPriceHour:  table.RenderInt(attr.Pricing.Brl.Hourly),
		BRLPriceMonth: table.RenderInt(attr.Pricing.Brl.Monthly),
	}
}

func CreateBandwidthPlanRows(plans []*models.BandwidthPlan) []interface{} {
	var rows []BandwidthPlanRow
	var rowsInterface []interface{}

	for _, plan := range plans {
		rows = append(rows, *NewBandwidthPlanRow(plan))
	}

	for _, row := range rows {
		rowsInterface = append(rowsInterface, row)
	}

	return rowsInterface
}
