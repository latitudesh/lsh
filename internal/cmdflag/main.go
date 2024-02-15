package cmdflag

import (
	"github.com/go-openapi/swag"
	"github.com/latitudesh/lsh/internal/prompt"
	"github.com/latitudesh/lsh/internal/utils"
	"github.com/spf13/pflag"
)

type FlagSchema interface {
	Register(f *pflag.FlagSet)
	GetValue() interface{}
	GetName() string
	Prompt() prompt.PromptInput
}

type FlagsSchema []FlagSchema

type Flags struct {
	schema            *FlagsSchema
	FlagSet           *pflag.FlagSet
	InteractiveMode   bool
	PromptDescription string
}

func (f *Flags) Register(s *FlagsSchema) {
	f.schema = s

	f.InteractiveMode = true // TODO: allow users to enable/disable interactive mode

	for _, v := range *s {
		v.Register(f.FlagSet)
	}
}

func (f *Flags) GetInputs() []prompt.PromptInput {
	var inputs []prompt.PromptInput

	for _, v := range *f.schema {
		inputs = append(inputs, v.Prompt())
	}

	return inputs
}

func (f *Flags) AssignValues(params interface{}) error {
	for _, v := range *f.schema {
		value := v.GetValue()

		if !swag.IsZero(value) {
			utils.AssignValue(params, v.GetName(), value)
		}
	}

	if f.InteractiveMode {
		p := prompt.Prompt{
			Description: f.PromptDescription,
			Inputs:      f.GetInputs(),
		}
		p.Run(params)
	}

	return nil
}
