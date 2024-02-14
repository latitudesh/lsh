package cmdflag

import (
	"fmt"

	"github.com/spf13/pflag"
)

type FlagSchema struct {
	Name         string
	Label        string
	Description  string
	DefaultValue interface{}
	Type         string
	Required     bool
}

func (s *FlagSchema) formattedDescription() string {
	if !s.Required {
		return s.Description
	}

	return fmt.Sprintf("[Required] %v", s.Description)
}

type FlagsSchema []FlagSchema

type Flags struct {
	Schema  *FlagsSchema
	FlagSet *pflag.FlagSet
}

func (f *Flags) Register(s *FlagsSchema) {
	f.Schema = s

	for _, v := range *f.Schema {
		switch v.Type {
		case "string":
			if defaultValue, ok := v.DefaultValue.(string); ok {
				f.FlagSet.String(v.Name, defaultValue, v.formattedDescription())
			}
		case "stringSlice":
			if defaultValue, ok := v.DefaultValue.([]string); ok {
				f.FlagSet.StringSlice(v.Name, defaultValue, v.formattedDescription())
			}
		case "int64":
			if defaultValue, ok := v.DefaultValue.(int64); ok {
				f.FlagSet.Int64(v.Name, defaultValue, v.formattedDescription())
			}
		case "int64Slice":
			if defaultValue, ok := v.DefaultValue.([]int64); ok {
				f.FlagSet.Int64Slice(v.Name, defaultValue, v.formattedDescription())
			}
		}
	}
}

func (f *Flags) ResourceIDFlagName() string {
	schema := *f.Schema
	return schema[0].Name
}
