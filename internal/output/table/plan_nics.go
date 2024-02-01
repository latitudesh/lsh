package table

import (
	"fmt"
	"strings"

	"github.com/latitudesh/lsh/models"
)

type PlanNICs struct {
	Value []*models.PlanDataAttributesSpecsNicsItems0
}

func RenderPlanNICs(value []*models.PlanDataAttributesSpecsNicsItems0) string {
	output := &PlanNICs{
		Value: value,
	}

	return output.Render()
}

func (t *PlanNICs) Render() string {
	var formattedNics strings.Builder

	for _, nic := range t.Value {
		formattedNics.WriteString(fmt.Sprintf("%vx %v\n", nic.Count, nic.Type))
	}

	return formattedNics.String()
}
