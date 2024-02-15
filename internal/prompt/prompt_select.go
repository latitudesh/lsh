package prompt

import (
	"fmt"
	"os"

	"github.com/latitudesh/lsh/internal/utils"
	"github.com/manifoldco/promptui"
)

type InputSelect struct {
	Name  string
	Label string
	Items []string
}

func NewInputSelect(name string, label string, items []string) *InputSelect {
	return &InputSelect{
		Name:  name,
		Label: label,
		Items: items,
	}
}

func (p *InputSelect) AssignValue(attributes interface{}) {
	currentValue := utils.GetFieldValue(attributes, p.Name).String()

	if currentValue == "" {
		prompt := promptui.Select{
			Label: p.Label,
			Items: p.Items,
		}

		_, value, err := prompt.Run()

		if err != nil {
			fmt.Printf("Failed to read input: %v\n", err)
			os.Exit(1)
		}

		if value == "SKIP" {
			return
		}

		utils.AssignValue(attributes, p.Name, value)
	}
}
