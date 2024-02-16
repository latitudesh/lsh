package cmdflag

import (
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
}

func (s *FlagSchema) FlagName() string {
	if s.CustomFlagName != "" {
		return s.CustomFlagName
	}

	return s.Name
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
				f.FlagSet.String(v.FlagName(), defaultValue, v.Description)
			}
		case "stringSlice":
			if defaultValue, ok := v.DefaultValue.([]string); ok {
				f.FlagSet.StringSlice(v.FlagName(), defaultValue, v.Description)
			}
		case "int64":
			if defaultValue, ok := v.DefaultValue.(int64); ok {
				f.FlagSet.Int64(v.FlagName(), defaultValue, v.Description)
			}
		case "bool":
			if defaultValue, ok := v.DefaultValue.(bool); ok {
				f.FlagSet.Bool(v.FlagName(), defaultValue, v.Description)
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
