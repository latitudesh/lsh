package table

import (
	"strings"
)

type StringList struct {
	Value []string
}

func RenderStringList(value []string, separator string) string {
	output := &StringList{
		Value: value,
	}

	return output.Render(separator)
}

func (t *StringList) Render(separator string) string {
	var output strings.Builder

	for i, v := range t.Value {
		output.WriteString(v)

		if i < len(t.Value)-1 {
			output.WriteString(separator)
		}
	}

	return output.String()
}
