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
	PromptID(params interface{})
	PromptAttributes(attributes interface{})
}

func AssignResourceID(o Operation, params interface{}) error {
	flags := o.GetFlags()
	idFlagName := flags.ResourceIDFlagName()
	id, err := flags.FlagSet.GetString(idFlagName)
	if err != nil {
		return err
	}

	utils.AssignValue(params, idFlagName, id)

	o.PromptID(params)

	return nil
}

func AssignBodyAttributes(o Operation, attributes interface{}) error {
	flags := o.GetFlags()

	for _, v := range *flags.Schema {
		if v.Name == "id" {
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
