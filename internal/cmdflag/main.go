package cmdflag

type Flag struct {
	Name         string
	Label        string
	DefaultValue string
	Type         string
}

type Flags []Flag
