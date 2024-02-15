package prompt

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/latitudesh/lsh/internal/utils"
	"github.com/manifoldco/promptui"
)

type InputStringList struct {
	Name  string
	Label string
}

func NewInputStringList(name, label string) *InputStringList {
	return &InputStringList{
		Name:  name,
		Label: label,
	}
}

func (p *InputStringList) AssignValue(attributes interface{}) {
	currentValue := utils.GetFieldValue(attributes, p.Name)

	if currentValue.Kind() == reflect.Slice && currentValue.Len() == 0 {
		prompt := promptui.Prompt{
			Label: fmt.Sprintf("%v (separate by comma, for example: option1,option2)", p.Label),
		}

		value, err := prompt.Run()
		if err != nil {
			utils.Exit("Failed to read input", err)
		}

		if len(strings.TrimSpace(value)) <= 0 {
			return
		}

		values := strings.Split(value, ",")
		utils.AssignValue(attributes, p.Name, values)
	}
}
