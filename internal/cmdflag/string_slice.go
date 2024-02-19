package cmdflag

import (
	"github.com/latitudesh/lsh/internal/prompt"
	"github.com/spf13/pflag"
)

type StringSlice struct {
	Name         string
	Label        string
	Description  string
	defaultValue []string
	Value        *[]string
	Required     bool
}

func (f *StringSlice) GetValue() interface{} {
	return *f.Value
}

func (f *StringSlice) GetName() string {
	return f.Name
}

func (f *StringSlice) Prompt() prompt.PromptInput {
	return prompt.NewInputStringList(f.Name, f.Label)
}

func (f *StringSlice) Register(s *pflag.FlagSet) {
	f.Value = s.StringSlice(f.Name, f.defaultValue, f.Description)
}
