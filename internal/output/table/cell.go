package table

func NewCell(label, value string) Cell {
	return Cell{
		Label: label,
		Value: value,
	}
}

type Cell struct {
	Value     string
	Label     string
	MaxLength int
}
