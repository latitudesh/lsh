package table

import "fmt"

type Float struct {
	Value float64
}

func RenderFloat(value float64) string {
	output := &Float{
		Value: value,
	}

	return output.Render()
}

func (t *Float) Render() string {
	return fmt.Sprint(t.Value)
}
