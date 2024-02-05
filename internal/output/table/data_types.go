package table

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
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
