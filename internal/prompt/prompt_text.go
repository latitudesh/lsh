package prompt

import (
	"fmt"
	"os"

	"github.com/latitudesh/lsh/internal/utils"
	"github.com/manifoldco/promptui"
)

type InputText struct {
	Name  string
	Label string
}

func NewInputText(name, label string) *InputText {
	return &InputText{
		Name:  name,
		Label: label,
	}
}

func (p *InputText) AssignValue(attributes interface{}) {
	currentValue := utils.GetFieldValue(attributes, p.Name).String()

	fmt.Println(p.Name, currentValue)

	if currentValue == "" {
		prompt := promptui.Prompt{
			Label: p.Label,
		}

		value, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			os.Exit(1)
		}

		utils.AssignValue(attributes, p.Name, value)
	}
}
