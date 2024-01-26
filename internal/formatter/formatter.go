package formatter

import (
	"fmt"
	"time"
)

func FormatValue(value interface{}) string {
	switch data := value.(type) {
	case string:
		return formatString(data)
	}

	return fmt.Sprint(value)
}

func formatString(value string) string {
	date, err := time.Parse(time.RFC3339, value)

	if err == nil {
		return date.Format("2006-01-02 15:04:05")
	}

	return truncate(value, 50)
}

func truncate(input string, maxLength int) string {
	if len(input) > maxLength {
		return input[:maxLength] + "..."
	}
	return input
}
