package table

import "fmt"

type Int struct {
	Value int64
}

func RenderInt(value int64) string {
	output := &Int{
		Value: value,
	}

	return output.Render()
}

func (t *Int) Render() string {
	return fmt.Sprint(t.Value)
}
