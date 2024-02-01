package table

type String struct {
	Value string
}

func RenderString(value string) string {
	output := &String{
		Value: value,
	}

	return output.Render()
}

func (t *String) Render() string {
	return t.Value
}
