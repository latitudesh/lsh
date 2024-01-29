package formatter

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

const MaxLength = 50
const ListSeparator = " • "

func FormatValue(header string, value interface{}) string {
	switch data := value.(type) {
	case string:
		return formatString(data)
	case map[string]interface{}:
		return formatNestedValue(header, data)
	case []interface{}:
		return formatList(header, data)
	}

	return fmt.Sprint(value)
}

func formatList(header string, list []interface{}) string {
	if header == "regions" {
		return formatPlansRegions(list)
	}

	var stringList []string
	for _, c := range list {
		if str, ok := c.(string); ok {
			stringList = append(stringList, str)
		}
	}

	return strings.Join(stringList, "\n\n")
}

func convertToStringList(input interface{}) string {
	// Check if the input is a slice
	inputValue := reflect.ValueOf(input)
	if inputValue.Kind() != reflect.Slice {
		return ""
	}

	// Convert each element to a string
	stringList := make([]string, inputValue.Len())
	for i := 0; i < inputValue.Len(); i++ {
		stringList[i] = fmt.Sprintf("%v", inputValue.Index(i).Interface())
	}

	// Join the strings with commas
	result := strings.Join(stringList, ", ")
	return result
}

func formatPlansRegions(regions []interface{}) string {
	var formattedRegions strings.Builder

	for _, region := range regions {
		r, ok := region.(map[string]interface{})

		if ok {
			var formattedRegion strings.Builder
			var formattedPrice strings.Builder

			regionName := fmt.Sprintf("%v\n\n", r["name"])
			formattedRegion.WriteString(regionName)

			locations := formatPlanLocations(r["locations"])
			formattedRegion.WriteString(locations)

			prices := formatPlanPrices(r["pricing"])
			formattedPrice.WriteString(prices)

			formattedRegions.WriteString(formattedRegion.String())
			formattedRegions.WriteString(formattedPrice.String())

			if len(regions) > 1 {
				formattedRegions.WriteString(ListSeparator)
			}
		}
	}

	return formattedRegions.String()
}

func formatPlansSpecs(specs map[string]interface{}) string {
	var formattedSpecs strings.Builder
	for spec, attributes := range specs {
		str, isString := attributes.(string)

		if isString {
			formattedSpecs.WriteString(fmt.Sprintf("%v %v", ListSeparator, str))
			continue
		}

		specsObject, isObject := attributes.(map[string]interface{})

		if isObject {
			switch spec {
			case "cpu":
				CPUSpecs := formatPlanSpecsCPU(specsObject)
				formattedSpecs.WriteString(CPUSpecs)
			case "memory":
				memorySpecs := formatPlanSpecsRAM(specsObject)
				formattedSpecs.WriteString(memorySpecs)
			}
		}

		specsList, isList := attributes.([]interface{})

		if isList {
			switch spec {
			case "drives":
				drivesSpecs := formatPlanSpecsDrives(specsList)
				formattedSpecs.WriteString(drivesSpecs)
			case "nics":
				nicsSpecs := formatPlanSpecsNICs(specsList)
				formattedSpecs.WriteString(nicsSpecs)
			}
		}
	}

	return formattedSpecs.String()
}

func formatPlanSpecsRAM(specs map[string]interface{}) string {
	return fmt.Sprintf("\n%v GB RAM\n", specs["total"])
}

func formatPlanSpecsCPU(specs map[string]interface{}) string {
	return fmt.Sprintf("\n%vx %v %vGHz (%v cores)\n", specs["count"], specs["type"], specs["clock"], specs["cores"])
}

func formatPlanSpecsDrives(specs []interface{}) string {
	var formattedDrives strings.Builder

	for _, drive := range specs {
		d, ok := drive.(map[string]interface{})

		if ok {
			formattedDrive := fmt.Sprintf("\n%vx %v %v\n", d["count"], d["type"], d["size"])
			formattedDrives.WriteString(formattedDrive)
		}
	}

	return formattedDrives.String()
}

func formatPlanSpecsNICs(specs []interface{}) string {
	var formattedNics strings.Builder

	for _, drive := range specs {
		d, ok := drive.(map[string]interface{})

		if ok {
			formattedNIC := fmt.Sprintf("\n%vx %v\n", d["count"], d["type"])
			formattedNics.WriteString(formattedNIC)
		}
	}

	return formattedNics.String()
}

func formatPlanLocations(locations interface{}) string {
	var formattedLocation strings.Builder
	l, ok := locations.(map[string]interface{})

	if ok {
		locationsAvailable := l["available"]
		locationsInStock := l["in_stock"]

		formattedLocation.WriteString(fmt.Sprintf("\n\nAvailable: %v\n", convertToStringList(locationsAvailable)))
		formattedLocation.WriteString(fmt.Sprintf("\n\nIn stock: %v\n", convertToStringList(locationsInStock)))
	}

	return formattedLocation.String()
}

func formatPlanPrices(prices interface{}) string {
	var formattedPrices strings.Builder
	p, ok := prices.(map[string]interface{})

	if ok {
		USDPrice := p["USD"]
		BRLPrice := p["BRL"]

		x, ok := USDPrice.(map[string]interface{})

		if ok {
			formattedPrices.WriteString(fmt.Sprintf("\nUSD Pricing\nHour: %v\nMonth: %v\nYear: %v\n", x["hour"], x["month"], x["year"]))
		}

		z, ok := BRLPrice.(map[string]interface{})

		if ok {
			formattedPrices.WriteString(fmt.Sprintf("\nBRL Pricing\nHour: %v\nMonth: %v\nYear: %v\n", z["hour"], z["month"], z["year"]))
		}
	}

	return formattedPrices.String()
}

func formatPlanPrice() string {
	return ""
}

func formatString(value string) string {
	date, err := time.Parse(time.RFC3339, value)

	if err == nil {
		return date.Format("2006-01-02 15:04:05")
	}

	return truncate(value)
}

func formatNestedValue(header string, value map[string]interface{}) string {
	// if header == "specs" {
	// 	var formattedSpecs strings.Builder

	// 	for _, v := range value {
	// 		formattedSpecs.WriteString(fmt.Sprintf("%v\n", v))
	// 	}

	// 	return formattedSpecs.String()
	// }

	if header == "specs" {
		return formatPlansSpecs(value)
	}

	if header == "region" {
		return fmt.Sprintf("%v", value["city"])
	}

	if header == "features" {
		return fmt.Sprintf("%v", value["city"])
	}

	if header == "user" {
		return fmt.Sprintf("%v", value["email"])
	}

	keys := [3]string{"name", "slug", "id"}

	for _, v := range keys {
		f, ok := value[v].(string)
		if !ok {
			continue
		}

		return f
	}

	return ""
}

func truncate(input string) string {
	if len(input) > MaxLength {
		return input[:MaxLength] + "..."
	}
	return input
}
