package operation

import (
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/utils"
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

		switch v.Type {
		case "string":
			value, err := flags.FlagSet.GetString(v.Name)

			if err != nil {
				return err
			}

			utils.AssignValue(attributes, v.Name, value)
		case "stringSlice":
			value, err := flags.FlagSet.GetStringSlice(v.Name)

			if err != nil {
				return err
			}

			utils.AssignValue(attributes, v.Name, value)
		case "int64":
			value, err := flags.FlagSet.GetInt64(v.Name)

			if err != nil {
				return err
			}

			utils.AssignValue(attributes, v.Name, value)
		}
	}

	o.PromptAttributes(attributes)

	return nil
}
