package cmdflag

import (
	"github.com/latitudesh/lsh/internal/prompt"
	"github.com/spf13/pflag"
)

type Int64 struct {
	Name         string
	Label        string
	Description  string
	Value        *int64
	defaultValue int64
}

func NewInt64(name string, label string, description string) *Int64 {
	return &Int64{
		Name:         name,
		Label:        label,
		Description:  description,
		defaultValue: int64(0),
	}
}

func (f *Int64) GetValue() interface{} {
	return *f.Value
}

func (f *Int64) GetName() string {
	return f.Name
}

func (f *Int64) Prompt() prompt.PromptInput {
	return prompt.NewInputNumber(f.Name, f.Label)
}

func (f *Int64) Register(s *pflag.FlagSet) {
	f.Value = s.Int64(f.Name, f.defaultValue, f.Description)
}
