package operation

import (
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Operation interface {
	GetFlagSet() *pflag.FlagSet
	GetFlagsDefinition() cmdflag.FlagsSchema
	RegisterFlags(cmd *cobra.Command)
	ResourceIDFlag() cmdflag.FlagSchema
	PromptID(params interface{})
	PromptAttributes(attributes interface{})
}

func AssignResourceID(o Operation, params interface{}) error {
	flagSet := o.GetFlagSet()
	idFlag := o.ResourceIDFlag()
	id, err := flagSet.GetString(idFlag.Name)
	if err != nil {
		return err
	}

	utils.AssignValue(params, idFlag.Name, id)

	o.PromptID(params)

	return nil
}

func AssignBodyAttributes(o Operation, attributes interface{}) error {
	flagSet := o.GetFlagSet()
	flags := o.GetFlagsDefinition()

	for _, v := range flags {
		if v.Name == "id" {
			continue
		}

		switch v.Type {
		case "string":
			value, err := flagSet.GetString(v.Name)

			if err != nil {
				return err
			}

			utils.AssignValue(attributes, v.Name, value)
		case "stringSlice":
			value, err := flagSet.GetStringSlice(v.Name)

			if err != nil {
				return err
			}

			utils.AssignValue(attributes, v.Name, value)
		case "int64":
			value, err := flagSet.GetInt64(v.Name)

			if err != nil {
				return err
			}

			utils.AssignValue(attributes, v.Name, value)
		}
	}

	o.PromptAttributes(attributes)

	return nil
}
