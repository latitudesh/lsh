package prompt

import (
	"fmt"
	"os"
	"strconv"

	"github.com/latitudesh/lsh/internal/utils"
	"github.com/manifoldco/promptui"
)

type InputNumber struct {
	Name  string
	Label string
}

func NewInputNumber(name, label string) *InputNumber {
	return &InputNumber{
		Name:  name,
		Label: label,
	}
}

func (p *InputNumber) AssignValue(attributes interface{}) {
	currentValue := utils.GetFieldValue(attributes, p.Name)

	if currentValue.IsZero() {
		prompt := promptui.Prompt{
			Label: p.Label,
		}

		value, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			os.Exit(1)
		}

		if value == "" {
			return
		}

		numberValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			fmt.Println("Invalid input")
		}

		utils.AssignValue(attributes, p.Name, numberValue)
	}
}
