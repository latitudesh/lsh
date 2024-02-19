package prompt

import (
	"fmt"
	"os"
	"strconv"

	"github.com/latitudesh/lsh/internal/utils"
	"github.com/manifoldco/promptui"
)

type InputBool struct {
	Name  string
	Label string
}

func NewInputBool(name string, label string) *InputBool {
	return &InputBool{
		Name:  name,
		Label: label,
	}
}

func (p *InputBool) AssignValue(attributes interface{}) {
	currentValue := utils.GetFieldValue(attributes, p.Name).Bool()

	if !currentValue {
		prompt := promptui.Select{
			Label: p.Label,
			Items: []string{"SKIP", "true", "false"},
		}

		_, value, err := prompt.Run()

		if err != nil {
			fmt.Printf("Failed to read input: %v\n", err)
			os.Exit(1)
		}

		if value == "SKIP" {
			return
		}

		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			fmt.Printf("Failed to read input: %v\n", err)
			os.Exit(1)
		}

		utils.AssignValue(attributes, p.Name, boolValue)
	}
}
