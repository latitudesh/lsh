package table

import (
	"fmt"
	"strings"

	"github.com/latitudesh/lsh/models"
)

type PlanDrives struct {
	Value []*models.PlanDataAttributesSpecsDrivesItems0
}

func RenderPlanDrives(value []*models.PlanDataAttributesSpecsDrivesItems0) string {
	output := &PlanDrives{
		Value: value,
	}

	return output.Render()
}

func (t *PlanDrives) Render() string {
	var formattedDrives strings.Builder

	for _, drive := range t.Value {
		formattedDrives.WriteString(fmt.Sprintf("%vx %v %v\n", drive.Count, drive.Type, drive.Size))
	}

	return formattedDrives.String()
}
