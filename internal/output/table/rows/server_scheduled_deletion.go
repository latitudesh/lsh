package tablerows

import (
	"github.com/latitudesh/lsh/internal/output/table"
	"github.com/latitudesh/lsh/models"
)

func NewServerScheduledDeletionRow(scheduledDeletion *models.ServerScheduleDeletionData) table.Row {
	attr := scheduledDeletion.Attributes

	return table.Row{
		"id": table.Cell{
			Label: "ID",
			Value: table.String(scheduledDeletion.ID),
		},
		"server_id": table.Cell{
			Label: "Server ID",
			Value: table.String(attr.ServerID),
		},
		"scheduled_at": table.Cell{
			Label: "Scheduled At",
			Value: table.String(attr.ScheduledDeletionAt),
		},
	}
}
