package table

import "strings"

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
	var output []string

	for _, v := range t.Value {
		output = append(output, v)
	}

	return strings.Join(output, ", ")
}
