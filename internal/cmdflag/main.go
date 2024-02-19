package cmdflag

import (
	"fmt"

	"github.com/spf13/pflag"
)

type RequestParamType int64

const (
	PathParam RequestParamType = iota
	QueryParam
	BodyParam
)

type FlagSchema struct {
	Name             string
	Label            string
	Description      string
	DefaultValue     interface{}
	Type             string
	RequestParamType RequestParamType
	CustomFlagName   string
	Required         bool
}

func (s *FlagSchema) FlagName() string {
	if s.CustomFlagName != "" {
		return s.CustomFlagName
	}

	return s.Name
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
		flagName := v.FlagName()

		switch v.Type {
		case "string":
			if defaultValue, ok := v.DefaultValue.(string); ok {
				f.FlagSet.String(flagName, defaultValue, v.formattedDescription())
			}
		case "stringSlice":
			if defaultValue, ok := v.DefaultValue.([]string); ok {
				f.FlagSet.StringSlice(flagName, defaultValue, v.formattedDescription())
			}
		case "bool":
			if defaultValue, ok := v.DefaultValue.(bool); ok {
				f.FlagSet.Bool(flagName, defaultValue, v.formattedDescription())
			}
		case "int64":
			if defaultValue, ok := v.DefaultValue.(int64); ok {
				f.FlagSet.Int64(flagName, defaultValue, v.formattedDescription())
			}
		}
	}
}

func (f *Flags) PathParamsFlags() FlagsSchema {
	var pathParamsFlags FlagsSchema

	for _, v := range *f.Schema {
		if v.RequestParamType == PathParam {
			pathParamsFlags = append(pathParamsFlags, v)
		}
	}

	return pathParamsFlags
}
