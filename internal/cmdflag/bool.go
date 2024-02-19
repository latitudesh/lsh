package cmdflag

import (
	"github.com/latitudesh/lsh/internal/prompt"
	"github.com/spf13/pflag"
)

type Bool struct {
	Name         string
	Label        string
	Description  string
	Value        *bool
	defaultValue bool
	Required     bool
}

func NewBool(name string, label string, description string) *Bool {
	return &Bool{
		Name:         name,
		Label:        label,
		Description:  description,
		defaultValue: false,
	}
}

func (f *Bool) GetValue() interface{} {
	return *f.Value
}

func (f *Bool) GetName() string {
	return f.Name
}

func (f *Bool) Prompt() prompt.PromptInput {
	return prompt.NewInputBool(f.Name, f.Label)
}

func (f *Bool) Register(s *pflag.FlagSet) {
	f.Value = s.Bool(f.Name, f.defaultValue, f.Description)
}
