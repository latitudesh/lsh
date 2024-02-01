package table

import (
	"fmt"

	"github.com/latitudesh/lsh/models"
)

type PlanCPU struct {
	Value models.PlanDataAttributesSpecsCPU
}

func RenderPlanCPU(value models.PlanDataAttributesSpecsCPU) string {
	output := &PlanCPU{
		Value: value,
	}

	return output.Render()
}

func (t *PlanCPU) Render() string {
	value := t.Value
	return fmt.Sprintf("%vx %v %vGHz (%v cores)", value.Count, value.Type, value.Clock, value.Cores)
}
