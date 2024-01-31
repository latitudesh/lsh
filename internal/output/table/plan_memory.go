package table

import (
	"fmt"

	"github.com/latitudesh/lsh/models"
)

type PlanMemory struct {
	Value models.PlanDataAttributesSpecsMemory
}

func RenderPlanMemory(value models.PlanDataAttributesSpecsMemory) string {
	output := &PlanMemory{
		Value: value,
	}

	return output.Render()
}

func (t *PlanMemory) Render() string {
	return fmt.Sprintf("%vGB", t.Value.Total)
}
