package table

import (
	"strings"
)

const Separator = ", "

type StringList struct {
	Value []string
}

func RenderStringList(value []string) string {
	output := &StringList{
		Value: value,
	}

	return output.Render()
}

func (t *StringList) Render() string {
	var output strings.Builder

	for i, v := range t.Value {
		output.WriteString(v)

		if i < len(t.Value)-1 {
			output.WriteString(Separator)
		}
	}

	return output.String()
}
