package operation

import (
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/utils"
	"github.com/spf13/pflag"
)

type Operation interface {
	GetFlagSet() *pflag.FlagSet
	GetFlagsDefinition() cmdflag.Flags
	RegisterFlags()
	PromptAttributes(attributes interface{})
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
