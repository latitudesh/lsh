package cmdflag

import "github.com/spf13/pflag"

type Flag struct {
	Name         string
	Label        string
	Description  string
	DefaultValue interface{}
	Type         string
}

type Flags []Flag

func (f *Flags) Register(flagset *pflag.FlagSet) {
	for _, v := range *f {
		switch v.Type {
		case "string":
			if defaultValue, ok := v.DefaultValue.(string); ok {
				flagset.String(v.Name, defaultValue, v.Description)
			}
		case "stringSlice":
			if defaultValue, ok := v.DefaultValue.([]string); ok {
				flagset.StringSlice(v.Name, defaultValue, v.Description)
			}
		case "int64":
			if defaultValue, ok := v.DefaultValue.(int64); ok {
				flagset.Int64(v.Name, defaultValue, v.Description)
			}
		}
	}
}
