package cmdflag

import (
	"fmt"
	"os"

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
	PromptDescription string
}

func (f *Flags) Register(s *FlagsSchema) {
	f.schema = s

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
	interactiveModeDisabled, err := f.FlagSet.GetBool("no-input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, v := range *f.schema {
		value := v.GetValue()

		if !swag.IsZero(value) {
			utils.AssignValue(params, v.GetName(), value)
		}
	}

	if !interactiveModeDisabled {
		p := prompt.Prompt{
			Description: f.PromptDescription,
			Inputs:      f.GetInputs(),
		}
		p.Run(params)
	}

	return nil
}
