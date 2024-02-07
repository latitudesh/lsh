package prompt

import (
	"fmt"
	"os"

	"github.com/latitudesh/lsh/internal/utils"
	"github.com/manifoldco/promptui"
)

type PromptInputSelect struct {
	Name  string
	Label string
	Items []string
}

func NewPromptInputSelect(name string, label string, items []string) *PromptInputSelect {
	return &PromptInputSelect{
		Name:  name,
		Label: label,
		Items: items,
	}
}

func (p *PromptInputSelect) AssignValue(attributes interface{}) {
	currentValue := utils.GetFieldValue(attributes, p.Name).String()

	if currentValue == "" {
		prompt := promptui.Select{
			Label: p.Label,
			Items: p.Items,
		}

		_, value, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			os.Exit(1)
		}

		utils.AssignValue(attributes, p.Name, value)
	}
}
