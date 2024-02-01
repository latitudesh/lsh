package table

import (
	"time"

	"github.com/go-openapi/strfmt"
)

type DateTime struct {
	Value *strfmt.DateTime
}

func RenderDateTime(value *strfmt.DateTime) string {
	output := &DateTime{
		Value: value,
	}

	return output.Render()
}

func (t *DateTime) Render() string {
	if t.Value == nil {
		return ""
	}

	dateString := t.Value.String()
	date, err := time.Parse(time.RFC3339, dateString)

	if err == nil {
		return date.Format("2006-01-02 15:04:05")
	}

	return dateString
}
