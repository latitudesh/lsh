package table

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/latitudesh/lsh/models"
)

func String(value string) string {
	return value
}

func StringList(value []string, separator string) string {
	var output strings.Builder

	for i, v := range value {
		output.WriteString(v)

		if i < len(value)-1 {
			output.WriteString(separator)
		}
	}

	return output.String()
}

func Int(value int64) string {
	return fmt.Sprint(value)
}

func Float(value float64) string {
	return fmt.Sprint(value)
}

func DateTime(value *strfmt.DateTime) string {
	if value == nil {
		return ""
	}

	dateString := value.String()
	date, err := time.Parse(time.RFC3339, dateString)

	if err == nil {
		return date.Format("2006-01-02 15:04:05")
	}

	return dateString
}

func PlanNICs(value []*models.PlanDataAttributesSpecsNicsItems0) string {
	var formattedNics strings.Builder

	for _, nic := range value {
		formattedNics.WriteString(fmt.Sprintf("%vx %v\n", nic.Count, nic.Type))
	}

	return formattedNics.String()
}

func PlanMemory(value models.PlanDataAttributesSpecsMemory) string {
	return fmt.Sprintf("%vGB", value.Total)
}

func PlanDrives(value []*models.PlanDataAttributesSpecsDrivesItems0) string {
	var formattedDrives strings.Builder

	for _, drive := range value {
		formattedDrives.WriteString(fmt.Sprintf("%vx %v %v\n", drive.Count, drive.Type, drive.Size))
	}

	return formattedDrives.String()
}

func PlanCPU(value models.PlanDataAttributesSpecsCPU) string {
	return fmt.Sprintf("%vx %v %vGHz (%v cores)", value.Count, value.Type, value.Clock, value.Cores)
}
