package formatter

import (
	"fmt"
	"time"
)

const MaxLength = 50

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

	return truncate(value)
}

func truncate(input string) string {
	if len(input) > MaxLength {
		return input[:MaxLength] + "..."
	}
	return input
}
