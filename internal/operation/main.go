package operation

import (
	"errors"
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/utils"
	"github.com/spf13/pflag"
)

type Operation interface {
	GetFlags() cmdflag.Flags
	PromptPathParams(params interface{})
	PromptQueryParams(params interface{})
	PromptAttributes(attributes interface{})
}

func AssignPathParams(o Operation, params interface{}) error {
	flags := o.GetFlags()

	for _, v := range flags.PathParamsFlags() {
		value, err := flags.FlagSet.GetString(v.Name)
		if err != nil {
			return err
		}

		if !swag.IsZero(value) {
			utils.AssignValue(params, v.Name, value)
		}
	}

	o.PromptPathParams(params)

	return nil
}

func AssignQueryParams(o Operation, params interface{}) error {
	flags := o.GetFlags()

	for _, v := range flags.QueryParamsFlags() {
		value, err := flags.FlagSet.GetString(v.Name)
		if err != nil {
			return err
		}

		if !swag.IsZero(value) {
			utils.AssignValue(params, v.Name, value)
		}
	}

	o.PromptQueryParams(params)

	return nil
}

func AssignBodyAttributes(o Operation, attributes interface{}) error {
	flags := o.GetFlags()

	for _, v := range *flags.Schema {
		if v.RequestParamType != cmdflag.BodyParam {
			continue
		}

		value, err := getFlagValue(v, flags.FlagSet)
		if err != nil {
			fmt.Println(err)
			return err
		}

		if !swag.IsZero(value) {
			utils.AssignValue(attributes, v.Name, value)
		}
	}

	o.PromptAttributes(attributes)

	return nil
}

func getFlagValue(flag cmdflag.FlagSchema, flagSet *pflag.FlagSet) (interface{}, error) {
	switch flag.Type {
	case "string":
		return flagSet.GetString(flag.Name)
	case "stringSlice":
		value, err := flagSet.GetStringSlice(flag.Name)
		if err != nil {
			return nil, err
		}

		if len(value) > 0 {
			return value, nil
		}

		return nil, nil
	case "int64":
		return flagSet.GetInt64(flag.Name)
	case "bool":
		return flagSet.GetBool(flag.Name)
	default:
		return nil, errors.New(fmt.Sprintf("Unsupported data type for flag: %v", flag.Name))
	}
}
