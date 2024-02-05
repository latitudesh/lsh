package input

import (
	"fmt"
	"os"

	"github.com/latitudesh/lsh/internal/utils"
	"github.com/manifoldco/promptui"
)

type PromptField struct {
	Name  string
	Label string
}

func NewPromptField(name, label string) *PromptField {
	return &PromptField{
		Name:  name,
		Label: label,
	}
}

type Prompt []*PromptField

func RunPrompt(attributes interface{}, prompt Prompt) {
	for _, v := range prompt {
		fieldValue := utils.GetFieldValue(attributes, v.Name).String()

		if fieldValue == "" {
			p := promptui.Prompt{
				Label: v.Label,
			}

			value, err := p.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				os.Exit(1)
			}

			utils.AssignValue(attributes, v.Name, value)
		}
	}
}
